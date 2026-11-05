const state = { skis: [], people: [], loans: [], reservations: [] };
const field = (object, name) => object[name] ?? object[name[0].toUpperCase() + name.slice(1)];
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
  renderTable("#ski-table", ["Ski-ID", "Marke", "Modell", "Größe", "Status", "Aktion"], state.skis.map(ski => {
    const id = field(ski, "id"), status = field(ski, "status");
    const action = status === "verliehen" ? `<button data-return="${id}">Leihe beenden</button>` : "";
    return [id, field(ski, "brand"), field(ski, "model"), field(ski, "size") || "-", `<span class="status ${statusClass(status)}">${status}</span>`, action];
  }));
  renderTable("#people-table", ["Person-ID", "Name", "Ausweis-Nummer", "Ausgeliehene Ski"], state.people.map(person => {
    const id = field(person, "id");
    const rented = state.loans.filter(loan => field(loan, "personID") === id).map(loan => field(loan, "skiID")).join(", ");
    return [id, `${field(person, "firstName")} ${field(person, "lastName")}`, field(person, "identification"), rented || "-"];
  }));
  renderTable("#reservation-table", ["Reservierungs-ID", "Person-ID", "Ski-ID", "Aktion"], state.reservations.filter(r => field(r, "status") === "aktiv").map(reservation => {
    const id = field(reservation, "id");
    return [id, field(reservation, "personID"), field(reservation, "skiID"),
      `<button data-convert="${id}">Miete beginnen</button> <button data-discard="${id}">Verwerfen</button>`];
  }));
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
  const birth_date = prompt("Geburtsdatum (ISO):");
  const identification = prompt("Ausweis-Nummer:");
  const content_type = prompt("Foto-MIME-Typ (image/png):", "image/png");
  if (!first_name || !last_name || !birth_date || !identification || !content_type) return;
  await request("/api/people", { method: "POST", body: JSON.stringify({ first_name, last_name, birth_date, identification, photo: { content_type } }) });
  await refresh();
}

async function startLoan() {
  const ski_id = prompt("Ski-ID:");
  const person_id = prompt("Person-ID:");
  if (!ski_id || !person_id) return;
  await request("/api/loans", { method: "POST", body: JSON.stringify({ ski_id, person_id }) });
  await refresh();
}

async function addReservation() {
  const ski_id = prompt("Ski-ID:");
  const person_id = prompt("Person-ID:");
  if (!ski_id || !person_id) return;
  await request("/api/reservations", { method: "POST", body: JSON.stringify({ ski_id, person_id }) });
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
    if (event.target.dataset.action === "add-reservation") await addReservation();
    if (event.target.dataset.return && confirm("Ausleihe beenden?")) await request(`/api/loans/${event.target.dataset.return}`, { method: "DELETE" });
    if (event.target.dataset.convert && confirm("Reservierung in Ausleihe umwandeln?")) await request(`/api/reservations/${event.target.dataset.convert}/convert`, { method: "POST" });
    if (event.target.dataset.discard && confirm("Reservierung verwerfen?")) await request(`/api/reservations/${event.target.dataset.discard}`, { method: "DELETE" });
    if (event.target.dataset.return || event.target.dataset.convert || event.target.dataset.discard) await refresh();
  } catch (error) { message(error.message, true); }
});

refresh();
