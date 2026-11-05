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

function promptPostalCode(defaultValue = "") {
  while (true) {
    const value = prompt("PLZ (4 oder 5 Ziffern):", defaultValue);
    if (value === null) return null;
    if (/^\d{4,5}$/.test(value.trim())) return value.trim();
    alert("Die PLZ muss aus vier oder fünf Ziffern bestehen.");
  }
}

function selectIDPhoto() {
  return new Promise(resolve => {
    const dialog = document.createElement("dialog");
    dialog.innerHTML = `
      <form method="dialog">
        <h3>Ausweisdatei auswählen</h3>
        <label for="id-photo-file">Datei auswählen</label>
        <input id="id-photo-file" type="file" accept="image/png,image/jpeg,.png,.jpg,.jpeg">
        <div class="dialog-actions">
          <button value="cancel">Abbrechen</button>
          <button value="confirm">Datei verwenden</button>
        </div>
      </form>`;
    dialog.addEventListener("close", () => {
     const file = dialog.querySelector("#id-photo-file").files?.[0];
     const confirmed = dialog.returnValue === "confirm";
     dialog.remove();
     if (!confirmed) {
       resolve(null);
       return;
     }
     const extension = file?.name.toLowerCase().split(".").pop();
     const contentType = file?.type.toLowerCase() ||
       (extension === "jpg" || extension === "jpeg" ? "image/jpeg" : extension === "png" ? "image/png" : "");
     if (!file || !["image/png", "image/jpeg"].includes(contentType)) {
       alert("Bitte wählen Sie eine PNG- oder JPEG-Datei als Ausweisfoto aus.");
       resolve(selectIDPhoto());
       return;
     }
     const reader = new FileReader();
     reader.addEventListener("load", () => resolve({
       content_type: file.type,
       data_url: reader.result
     }), { once: true });
     reader.addEventListener("error", () => {
       alert("Das Ausweisfoto konnte nicht gelesen werden.");
       resolve(selectIDPhoto());
     }, { once: true });
     reader.readAsDataURL(file);
    }, { once: true });
    document.body.append(dialog);
    dialog.showModal();
  });
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
  renderTable("#ski-table", ["Ski-ID", "Marke", "Modell", "Größe", "Tagespreis", "Status", "Aktion"], sortSkis(state.skis).map(ski => {
    const id = field(ski, "id"), status = field(ski, "status");
    const action = status === "verliehen" ? `<button data-return="${id}">Leihe beenden</button>` : "";
    return [id, field(ski, "brand"), field(ski, "model"), field(ski, "size") || "-", formatPrice(field(ski, "dailyPriceCents")), `<span class="status ${statusClass(status)}">${status}</span>`, action];
  }));
  renderTable("#people-table", ["Person-ID", "Name", "Adresse", "Handynummer", "Ausweis-Nummer", "Ausweisdokument", "Ausgeliehene Ski", "Aktion"], [...state.people].sort(comparePeople).map(person => {
    const id = field(person, "id");
    const activeLoans = state.loans.filter(loan => field(loan, "personID") === id && !field(loan, "endedAt"));
    const rented = sortLoanSkis(activeLoans).map(loan => field(loan, "skiID")).join(", ");
    const address = `${field(person, "street")} ${field(person, "houseNumber")}, ${field(person, "postalCode")} ${field(person, "city")}, ${field(person, "country")}`;
    const photoAction = field(person, "idPhoto")?.dataUrl || field(person, "IDPhoto")?.DataURL
      ? `<button data-person-photo="${id}">Ausweis ansehen</button>` : "-";
    const editAction = `<button data-edit-person="${id}">Stammdaten bearbeiten</button>`;
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
    const reservationForPersonAction = `<button data-person-reservation="${id}">Registrierung anlegen</button>`;
    const action = `${editAction} ${returnActions} ${reservationActions} ${reservationForPersonAction}`;
    return [id, `${field(person, "firstName")} ${field(person, "lastName")}`, address, field(person, "mobileNumber"), field(person, "identification"), photoAction, rented || "-", action];
  }));
  renderTable("#reservation-table", ["Reservierungs-ID", "Person-ID", "Person", "Ski-ID", "Beginn", "Ende", "Preis", "Aktion"], state.reservations.filter(r => field(r, "status") === "aktiv").map(reservation => {
    const id = field(reservation, "id");
    const person = state.people.find(item => field(item, "id") === field(reservation, "personID"));
    const personName = person
      ? `${field(person, "firstName")} ${field(person, "lastName")}`
      : "Unbekannte Person";
    const days = rentalDays(field(reservation, "startedAt"), field(reservation, "endDate"));
    const total = days * (Number(field(reservation, "dailyPriceCents")) || 0);
    return [id, field(reservation, "personID"), personName, field(reservation, "skiID"),
      formatDate(field(reservation, "startedAt")), formatDate(field(reservation, "endDate")), formatPrice(total),
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

function viewPersonPhoto(personID) {
  const person = state.people.find(item => field(item, "id") === personID);
  const photo = field(person, "idPhoto") || field(person, "IDPhoto");
  const dataURL = field(photo, "dataUrl") || field(photo, "DataURL");
  if (!dataURL) {
    alert("Für diese Person ist kein Ausweisdokument hinterlegt.");
    return;
  }
  const dialog = document.createElement("dialog");
  dialog.innerHTML = `
    <h3>Ausweisdokument von ${field(person, "firstName")} ${field(person, "lastName")}</h3>
    <img src="${dataURL}" alt="Ausweisdokument" style="max-width: min(90vw, 700px); max-height: 70vh;">
    <form method="dialog"><button>Schließen</button></form>`;
  document.body.append(dialog);
  dialog.showModal();
  dialog.addEventListener("close", () => dialog.remove(), { once: true });
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
  const brand = promptRequired("Marke:");
  if (brand === null) return;
  const model = promptRequired("Modell:");
  if (model === null) return;
  const size = promptPositiveInteger("Länge (cm):");
  if (size === null) return;
  const daily_price = promptPrice();
  if (daily_price === null) return;
  await request("/api/skis", { method: "POST", body: JSON.stringify({ brand, model, size, daily_price }) });
  await refresh();
}

async function editPerson(personID) {
  const person = state.people.find(item => field(item, "id") === personID);
  if (!person) return;
  const first_name = promptRequired("Vorname:", field(person, "firstName"));
  if (first_name === null) return;
  const last_name = promptRequired("Nachname:", field(person, "lastName"));
  if (last_name === null) return;
  const currentBirthDate = new Date(field(person, "birthDate")).toLocaleDateString("de-DE");
  const birth_date = promptBirthDate(currentBirthDate);
  if (birth_date === null) return;
  const identification = promptRequired("Ausweis-Nummer:", field(person, "identification"));
  if (identification === null) return;
  const country = promptRequired("Land:", field(person, "country"));
  if (country === null) return;
  const city = promptRequired("Stadt:", field(person, "city"));
  if (city === null) return;
  const postal_code = promptPostalCode(field(person, "postalCode"));
  if (postal_code === null) return;
  const street = promptRequired("Straße:", field(person, "street"));
  if (street === null) return;
  const house_number = promptRequired("Hausnummer:", field(person, "houseNumber"));
  if (house_number === null) return;
  const mobile_number = promptRequired("Handynummer:", field(person, "mobileNumber"));
  if (mobile_number === null) return;
  while (true) {
    const photo = await selectIDPhoto();
    if (photo === null) return;
    try {
      await request(`/api/people/${personID}`, { method: "PUT", body: JSON.stringify({
        first_name, last_name, birth_date, identification, country, city, postal_code,
        street, house_number, mobile_number, photo
      }) });
      break;
    } catch (error) {
      if (!error.message.toLowerCase().includes("id photo") &&
          !error.message.toLowerCase().includes("png") &&
          !error.message.toLowerCase().includes("jpeg")) throw error;
      alert(error.message);
    }
  }
  await refresh();
}

function promptRequired(label, defaultValue = "") {
  while (true) {
    const value = prompt(label, defaultValue);
    if (value === null) return null;
    if (value.trim()) return value.trim();
    alert("Dieses Feld darf nicht leer sein.");
  }
}

function promptPositiveInteger(label) {
  while (true) {
    const value = prompt(label);
    if (value === null) return null;
    if (/^[1-9]\d*$/.test(value.trim())) return Number(value);
    alert("Bitte geben Sie eine positive ganze Zahl ein.");
  }
}

function promptPrice(defaultValue = "10.00") {
  while (true) {
    const value = prompt("Tagespreis in EUR (10,00 bis 20,00):", defaultValue);
    if (value === null) return null;
    const normalized = value.trim().replace(",", ".");
    const price = Number(normalized);
    if (Number.isFinite(price) && price >= 10 && price <= 20) return Number(price.toFixed(2));
    alert("Der Tagespreis muss zwischen 10,00 EUR und 20,00 EUR liegen.");
  }
}

function formatPrice(cents) {
  return `${((Number(cents) || 0) / 100).toFixed(2).replace(".", ",")} EUR`;
}

function formatDate(value) {
  if (!value) return "-";
  return new Date(value).toLocaleDateString("de-DE");
}

function rentalDays(start, end) {
  if (!start || !end) return 1;
  return Math.max(1, Math.round((new Date(end) - new Date(start)) / 86400000) + 1);
}

function confirmRentalPrice(ski_ids, period) {
  const days = rentalDays(period.start_date, period.end_date);
  const total = ski_ids.reduce((sum, id) => {
    const ski = state.skis.find(item => field(item, "id") === id);
    return sum + days * (Number(field(ski, "dailyPriceCents")) || 0);
  }, 0);
  return confirm(`${days} Miettag(e), Gesamtpreis für die Auswahl: ${formatPrice(total)}. Vorgang starten?`);
}

function todayISO() {
  return new Date().toISOString().slice(0, 10);
}

function promptRentalPeriod(allowFuture) {
  const today = todayISO();
  const start_date = prompt(`Mietbeginn (JJJJ-MM-TT, ${allowFuture ? "heute oder später" : "heute"}):`, today);
  if (start_date === null) return null;
  const end_date = prompt("Mietende (JJJJ-MM-TT, inklusive):", start_date);
  if (end_date === null) return null;
  if (!/^\d{4}-\d{2}-\d{2}$/.test(start_date) || !/^\d{4}-\d{2}-\d{2}$/.test(end_date) ||
      start_date < today || end_date < start_date || (!allowFuture && start_date !== today)) {
    alert("Ungültiger Zeitraum. Bitte verwenden Sie gültige Daten; das Ende muss am oder nach dem Beginn liegen.");
    return promptRentalPeriod(allowFuture);
  }
  return { start_date, end_date };
}

async function registerPerson() {
  const first_name = promptRequired("Vorname:");
  if (first_name === null) return;
  const last_name = promptRequired("Nachname:");
  if (last_name === null) return;
  const birth_date = promptBirthDate();
  if (birth_date === null) return;
  const identification = promptRequired("Ausweis-Nummer:");
  if (identification === null) return;
  const country = promptRequired("Land:");
  if (country === null) return;
  const city = promptRequired("Stadt:");
  if (city === null) return;
  const postal_code = promptPostalCode();
  if (postal_code === null) return;
  const street = promptRequired("Straße:");
  if (street === null) return;
  const house_number = promptRequired("Hausnummer:");
  if (house_number === null) return;
  const mobile_number = promptRequired("Handynummer:");
  if (mobile_number === null) return;
  while (true) {
    const photo = await selectIDPhoto();
    if (photo === null) return;
    try {
      await request("/api/people", { method: "POST", body: JSON.stringify({
        first_name, last_name, birth_date, identification, country, city, postal_code,
        street, house_number, mobile_number, photo
      }) });
      break;
    } catch (error) {
      if (!error.message.toLowerCase().includes("id photo") &&
          !error.message.toLowerCase().includes("png") &&
          !error.message.toLowerCase().includes("jpeg")) throw error;
      alert(error.message);
    }
  }
  await refresh();
}

function promptBirthDate(defaultValue = "") {
  while (true) {
    const value = prompt("Geburtsdatum (DD.MM.YYYY):", defaultValue);
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
        const hundredthBirthday = new Date(date);
        hundredthBirthday.setUTCFullYear(hundredthBirthday.getUTCFullYear() + 100);
        if (hundredthBirthday < today) {
          alert("Die Person darf nicht älter als 100 Jahre sein.");
          continue;
        }
        return value;
      }
    }
    alert("Ungültiges Datum. Bitte verwenden Sie das Format DD.MM.YYYY, zum Beispiel 05.11.1990.");
  }
}

async function startLoan() {
  const person_id = await selectPersonByName();
  if (!person_id) return;
  const ski_ids = await selectAvailableSkis();
  if (!ski_ids.length) return;
  const period = promptRentalPeriod(false);
  if (!period) return;
  if (!confirmRentalPrice(ski_ids, period)) return;
  await startLoans(ski_ids, person_id, period);
  await refresh();
}

async function startLoanForPerson(personID) {
  const ski_ids = await selectAvailableSkis();
  if (!ski_ids.length || !confirm("Ausleihe für die ausgewählten Ski starten?")) return;
  const period = promptRentalPeriod(false);
  if (!period) return;
  if (!confirmRentalPrice(ski_ids, period)) return;
  try {
    await startLoans(ski_ids, personID, period);
  } catch (error) {
    alert(`Ausleihe konnte nicht gestartet werden: ${error.message}`);
    return;
  }
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
      const label = `${id} – ${field(ski, "brand")} ${field(ski, "model")} (${field(ski, "size") || "-"} cm, ${formatPrice(field(ski, "dailyPriceCents"))}/Tag)`;
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

async function startLoans(ski_ids, person_id, period) {
  for (const ski_id of ski_ids) {
    await request("/api/loans", { method: "POST", body: JSON.stringify({ ski_id, person_id, ...period }) });
  }
}

async function editReservation(id) {
  const ski_id = promptRequired("Neue Ski-ID:");
  if (ski_id === null) return;
  const period = promptRentalPeriod(true);
  if (!period || !confirm("Reservierung bearbeiten?")) return;
  await request(`/api/reservations/${id}`, { method: "PUT", body: JSON.stringify({ ski_id, ...period }) });
  await refresh();
}

async function addReservation() {
  const person_id = await selectPersonByName();
  if (!person_id) return;
  await addReservationForPerson(person_id);
}

async function addReservationForPerson(person_id) {
  const ski_ids = await selectAvailableSkis("Verfügbare Ski für Reservierung auswählen", "reservation-ski", "Reservierung anlegen");
  if (!ski_ids.length) return;
  const period = promptRentalPeriod(true);
  if (!period) return;
  for (const ski_id of ski_ids) {
    await request("/api/reservations", { method: "POST", body: JSON.stringify({ ski_id, person_id, ...period }) });
  }
  await refresh();
}

async function selectPersonByName() {
  const lastName = promptRequired("Nachname:");
  if (lastName === null) return null;
  const firstName = promptRequired("Vorname:");
  if (firstName === null) return null;
  const normalizedLastName = lastName.toLocaleLowerCase("de");
  const normalizedFirstName = firstName.toLocaleLowerCase("de");
  const matches = [...state.people]
    .filter(person => field(person, "lastName").toLocaleLowerCase("de").includes(normalizedLastName) &&
      field(person, "firstName").toLocaleLowerCase("de").includes(normalizedFirstName))
    .sort(comparePeople);
  if (!matches.length) {
    alert("Keine passende Person gefunden.");
    return null;
  }
  return selectOnePerson(matches);
}

function selectOnePerson(people) {
  return new Promise(resolve => {
    const dialog = document.createElement("dialog");
    const options = people.map((person, index) => {
      const id = field(person, "id");
      const address = `${field(person, "street")} ${field(person, "houseNumber")}, ${field(person, "postalCode")} ${field(person, "city")}, ${field(person, "country")}`;
      const birthDate = new Date(field(person, "birthDate")).toLocaleDateString("de-DE");
      const name = `${field(person, "lastName")}, ${field(person, "firstName")} – geboren am ${birthDate} – ${address} – ${id}`;
      return `<label><input type="radio" name="person" value="${id}"${index === 0 ? " checked" : ""}> ${name}</label>`;
    }).join("");
    dialog.innerHTML = `
      <form method="dialog">
        <h3>Person auswählen</h3>
        <div class="ski-options">${options}</div>
        <div class="dialog-actions">
          <button value="cancel">Abbrechen</button>
          <button value="confirm">Person auswählen</button>
        </div>
      </form>`;
    dialog.addEventListener("close", () => {
      const selected = dialog.returnValue === "confirm"
        ? dialog.querySelector('input[name="person"]:checked')?.value || null
        : null;
      dialog.remove();
      resolve(selected);
    }, { once: true });
    document.body.append(dialog);
    dialog.showModal();
  });
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
    if (event.target.dataset.editPerson) await editPerson(event.target.dataset.editPerson);
    if (event.target.dataset.personPhoto) viewPersonPhoto(event.target.dataset.personPhoto);
    if (event.target.dataset.action === "start-loan") await startLoan();
    if (event.target.dataset.personLoan) await startLoanForPerson(event.target.dataset.personLoan);
    if (event.target.dataset.personReturn) await endLoansForPerson(event.target.dataset.personReturn);
    if (event.target.dataset.action === "add-reservation") await addReservation();
    if (event.target.dataset.personReservation) await addReservationForPerson(event.target.dataset.personReservation);
    if (event.target.dataset.return && confirm("Ausleihe beenden?")) await request(`/api/loans/${event.target.dataset.return}`, { method: "DELETE" });
    if (event.target.dataset.convert && confirm("Reservierung in Ausleihe umwandeln?")) await request(`/api/reservations/${event.target.dataset.convert}/convert`, { method: "POST" });
    if (event.target.dataset.editReservation) await editReservation(event.target.dataset.editReservation);
    if (event.target.dataset.discard && confirm("Reservierung verwerfen?")) await request(`/api/reservations/${event.target.dataset.discard}`, { method: "DELETE" });
    if (event.target.dataset.return || event.target.dataset.convert || event.target.dataset.editReservation || event.target.dataset.discard) await refresh();
  } catch (error) { message(error.message, true); }
});

refresh();
