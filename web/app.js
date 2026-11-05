const state = { skis: [], people: [], loans: [], reservations: [] };
const field = (object, name) => object[name] ??
  object[name[0].toUpperCase() + name.slice(1)] ?? object[name.toUpperCase()];
const message = (text, error = false) => {
  const element = document.querySelector("#message");
  element.textContent = text;
  element.className = error ? "error" : "";
};

async function request(path, options = {}) {
  const response = await fetch(path, { headers: { "Content-Type": "application/json" }, ...options });
  const body = response.status === 204 ? null : await response.json();
  if (!response.ok) throw new Error(body?.error || `HTTP ${response.status}`);
  return body;
}

async function refresh() {
  try {
    [state.skis, state.people, state.loans, state.reservations] = await Promise.all([
      request("/api/skis"), request("/api/people"), request("/api/loans"), request("/api/reservations")
    ]);
    render();
    message("");
  } catch (error) {
    message(error.message, true);
  }
}

function render() {
  renderDashboard();
  renderTable("#ski-table", ["Ski-ID", "Marke", "Modell", "Größe", "Status", "Aktion"], sortSkis(state.skis).map(ski => {
    const id = field(ski, "id"), status = field(ski, "status");
    const action = status === "verliehen" ? `<button data-return="${id}">Leihe beenden</button>` : "";
    return [id, field(ski, "brand"), field(ski, "model"), field(ski, "size") || "-", `<span class="status ${statusClass(status)}">${status}</span>`, action];
  }));
  renderTable("#people-table", ["Person-ID", "Name", "Ausweis-Nummer", "Ausgeliehene Ski", "Aktion"], [...state.people].sort(comparePeople).map(person => {
    const id = field(person, "id");
    const activeLoans = state.loans.filter(loan => field(loan, "personID") === id && !field(loan, "endedAt"));
    const rented = sortLoanSkis(activeLoans).map(loan => field(loan, "skiID")).join(", ");
    const returnActions = activeLoans.length
      ? `<button data-person-return="${id}">Ausleihe beenden</button>`
      : "";
    const activeReservations = state.reservations.filter(item =>
      field(item, "personID") === id && field(item, "status") === "aktiv");
    const reservationActions = activeReservations.length
      ? activeReservations.map(reservation => `<button data-convert="${field(reservation, "id")}">Miete reservierter Ski beginnen</button>
         <button data-edit-reservation="${field(reservation, "id")}">Reservierung bearbeiten</button>
         <button data-discard="${field(reservation, "id")}">Reservierung löschen</button>`).join(" ")
      : `<button data-person-loan="${id}">Neue Ausleihe</button>`;
    const action = `${returnActions} ${reservationActions}`;
    return [id, `${field(person, "firstName")} ${field(person, "lastName")}`, field(person, "identification"), rented || "-", action];
  }));
  renderTable("#reservation-table", ["Reservierungs-ID", "Person-ID", "Ski-ID", "Aktion"], state.reservations.filter(r => field(r, "status") === "aktiv").map(reservation => {
    const id = field(reservation, "id");
    return [id, field(reservation, "personID"), field(reservation, "skiID"),
      `<button data-convert="${id}">Miete beginnen</button> <button data-discard="${id}">Verwerfen</button>`];
  }));
}

function compareText(left, right) {
  return String(left ?? "").localeCompare(String(right ?? ""), "de", {
    numeric: true,
    sensitivity: "base"
  });
}

function compareSkis(left, right) {
  return compareText(field(left, "brand"), field(right, "brand")) ||
    compareText(field(left, "model"), field(right, "model")) ||
    (Number(field(left, "size")) || 0) - (Number(field(right, "size")) || 0) ||
    compareText(field(left, "id"), field(right, "id"));
}

function sortSkis(skis) {
  return [...skis].sort(compareSkis);
}

function sortLoanSkis(loans) {
  return [...loans].sort((left, right) => {
    const leftSki = state.skis.find(ski => field(ski, "id") === field(left, "skiID")) || {};
    const rightSki = state.skis.find(ski => field(ski, "id") === field(right, "skiID")) || {};
    return compareSkis(leftSki, rightSki);
  });
}

function comparePeople(left, right) {
  return compareText(field(left, "lastName"), field(right, "lastName")) ||
    compareText(field(left, "firstName"), field(right, "firstName")) ||
    compareText(field(left, "id"), field(right, "id"));
}

function renderDashboard() {
  request("/api/dashboard").then(data => {
    document.querySelector("#metrics").innerHTML = [
      ["Verfügbar", data.available_skis], ["Verliehen", data.rented_skis],
      ["Gesamt", data.total_skis], ["Personen", data.people]
    ].map(([label, value]) => `<div class="card"><strong>${value}</strong>${label}</div>`).join("");
  }).catch(error => message(error.message, true));
}

function statusClass(status) {
  return status === "verfügbar" ? "available" : status === "reserviert" ? "reserved" : "rented";
}

function renderTable(selector, headers, rows) {
  const target = document.querySelector(selector);
  if (!rows.length) { target.innerHTML = '<p class="empty">Keine Einträge vorhanden.</p>'; return; }
  target.innerHTML = `<table><thead><tr>${headers.map(header => `<th>${header}</th>`).join("")}</tr></thead><tbody>${rows
    .map(row => `<tr>${row.map(cell => `<td>${cell}</td>`).join("")}</tr>`).join("")}</tbody></table>`;
}

async function addSki() {
  const brand = prompt("Marke:");
  const model = prompt("Modell:");
  if (!brand || !model) return;
  await request("/api/skis", { method: "POST", body: JSON.stringify({ brand, model }) });
  await refresh();
}

async function registerPerson() {
  const first_name = prompt("Vorname:");
  const last_name = prompt("Nachname:");
  const birth_date = promptBirthDate();
  const identification = prompt("Ausweis-Nummer:");
  const content_type = prompt("Foto-MIME-Typ (image/png):", "image/png");
  if (!first_name || !last_name || birth_date === null || !identification || !content_type) return;
  await request("/api/people", { method: "POST", body: JSON.stringify({ first_name, last_name, birth_date, identification, photo: { content_type } }) });
  await refresh();
}

function promptBirthDate() {
  while (true) {
    const value = prompt("Geburtsdatum (DD.MM.YYYY):");
    if (value === null) return null;
    if (/^\d{2}\.\d{2}\.\d{4}$/.test(value)) {
      const [day, month, year] = value.split(".").map(Number);
      const date = new Date(Date.UTC(year, month - 1, day));
      if (date.getUTCFullYear() === year && date.getUTCMonth() === month - 1 && date.getUTCDate() === day) {
        const now = new Date();
        const today = new Date(Date.UTC(now.getFullYear(), now.getMonth(), now.getDate()));
        if (date > today) {
          alert("Das Geburtsdatum darf nicht in der Zukunft liegen.");
          continue;
        }
        const eighteenthBirthday = new Date(date);
        eighteenthBirthday.setUTCFullYear(eighteenthBirthday.getUTCFullYear() + 18);
        if (eighteenthBirthday > today) {
          alert("Die Person muss mindestens 18 Jahre alt sein.");
          continue;
        }
        return value;
      }
    }
    alert("Ungültiges Datum. Bitte verwenden Sie das Format DD.MM.YYYY, zum Beispiel 05.11.1990.");
  }
}

async function startLoan() {
  const person_id = prompt("Person-ID:");
  if (!person_id) return;
  const ski_ids = await selectAvailableSkis();
  if (!ski_ids.length) return;
  await startLoans(ski_ids, person_id);
  await refresh();
}

async function startLoanForPerson(personID) {
  const ski_ids = await selectAvailableSkis();
  if (!ski_ids.length || !confirm("Ausleihe für die ausgewählten Ski starten?")) return;
  await startLoans(ski_ids, personID);
  await refresh();
}

async function endLoansForPerson(personID) {
  const ski_ids = await selectPersonLoans(personID);
  if (!ski_ids.length) return;
  for (const ski_id of ski_ids) {
    await request(`/api/loans/${ski_id}`, { method: "DELETE" });
  }
  await refresh();
}

function selectAvailableSkis(title = "Verfügbare Ski auswählen", inputName = "ski", actionLabel = "Ausleihe starten") {
  const available = sortSkis(state.skis.filter(ski => field(ski, "status") === "verfügbar"));
  if (!available.length) {
    alert("Es sind keine verfügbaren Ski vorhanden.");
    return Promise.resolve([]);
  }

  return new Promise(resolve => {
    const dialog = document.createElement("dialog");
    const options = available.map(ski => {
      const id = field(ski, "id");
      const label = `${id} – ${field(ski, "brand")} ${field(ski, "model")} (${field(ski, "size") || "-"} cm)`;
      return `<label><input type="checkbox" name="${inputName}" value="${id}"> ${label}</label>`;
    }).join("");
    dialog.innerHTML = `
      <form method="dialog">
        <h3>${title}</h3>
        <div class="ski-options">${options}</div>
        <div class="dialog-actions">
          <button value="cancel">Abbrechen</button>
          <button value="confirm">${actionLabel}</button>
        </div>
      </form>`;
    dialog.addEventListener("close", () => {
      const selected = dialog.returnValue === "confirm"
        ? [...dialog.querySelectorAll(`input[name="${inputName}"]:checked`)].map(input => input.value)
        : [];
      dialog.remove();
      resolve(selected);
    }, { once: true });
    document.body.append(dialog);
    dialog.showModal();
  });
}

function selectPersonLoans(personID) {
  const loans = sortLoanSkis(state.loans.filter(loan =>
    field(loan, "personID") === personID && !field(loan, "endedAt")));
  return new Promise(resolve => {
    const dialog = document.createElement("dialog");
    const options = loans.map(loan => {
      const skiID = field(loan, "skiID");
      const ski = state.skis.find(item => field(item, "id") === skiID);
      const label = ski
        ? `${skiID} – ${field(ski, "brand")} ${field(ski, "model")} (${field(ski, "size") || "-"} cm)`
        : skiID;
      return `<label><input type="checkbox" name="loan-ski" value="${skiID}"> ${label}</label>`;
    }).join("");
    dialog.innerHTML = `
      <form method="dialog">
        <h3>Ausleihen beenden</h3>
        <div class="ski-options">${options}</div>
        <div class="dialog-actions">
          <button value="cancel">Abbrechen</button>
          <button value="confirm">Ausleihe beenden</button>
        </div>
      </form>`;
    dialog.addEventListener("close", () => {
      const selected = dialog.returnValue === "confirm"
        ? [...dialog.querySelectorAll('input[name="loan-ski"]:checked')].map(input => input.value)
        : [];
      dialog.remove();
      resolve(selected);
    }, { once: true });
    document.body.append(dialog);
    dialog.showModal();
  });
}

async function startLoans(ski_ids, person_id) {
  for (const ski_id of ski_ids) {
    await request("/api/loans", { method: "POST", body: JSON.stringify({ ski_id, person_id }) });
  }
}

async function editReservation(id) {
  const ski_id = prompt("Neue Ski-ID:");
  if (!ski_id || !confirm("Reservierung bearbeiten?")) return;
  await request(`/api/reservations/${id}`, { method: "PUT", body: JSON.stringify({ ski_id }) });
  await refresh();
}

async function addReservation() {
  const person_id = prompt("Person-ID:");
  if (!person_id) return;
  const ski_ids = await selectAvailableSkis("Verfügbare Ski für Reservierung auswählen", "reservation-ski", "Reservierung anlegen");
  if (!ski_ids.length) return;
  for (const ski_id of ski_ids) {
    await request("/api/reservations", { method: "POST", body: JSON.stringify({ ski_id, person_id }) });
  }
  await refresh();
}

document.addEventListener("click", async event => {
  const view = event.target.dataset.view;
  if (view) {
    document.querySelectorAll(".tab, .view").forEach(element => element.classList.remove("active"));
    event.target.classList.add("active");
    document.querySelector(`#${view}`).classList.add("active");
  }
  try {
    if (event.target.dataset.action === "add-ski") await addSki();
    if (event.target.dataset.action === "register-person") await registerPerson();
    if (event.target.dataset.action === "start-loan") await startLoan();
    if (event.target.dataset.personLoan) await startLoanForPerson(event.target.dataset.personLoan);
    if (event.target.dataset.personReturn) await endLoansForPerson(event.target.dataset.personReturn);
    if (event.target.dataset.action === "add-reservation") await addReservation();
    if (event.target.dataset.return && confirm("Ausleihe beenden?")) await request(`/api/loans/${event.target.dataset.return}`, { method: "DELETE" });
    if (event.target.dataset.convert && confirm("Reservierung in Ausleihe umwandeln?")) await request(`/api/reservations/${event.target.dataset.convert}/convert`, { method: "POST" });
    if (event.target.dataset.editReservation) await editReservation(event.target.dataset.editReservation);
    if (event.target.dataset.discard && confirm("Reservierung verwerfen?")) await request(`/api/reservations/${event.target.dataset.discard}`, { method: "DELETE" });
    if (event.target.dataset.return || event.target.dataset.convert || event.target.dataset.editReservation || event.target.dataset.discard) await refresh();
  } catch (error) { message(error.message, true); }
});

refresh();
