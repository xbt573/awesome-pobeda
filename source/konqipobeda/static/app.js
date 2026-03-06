const startDate = new Date('1999-04-01T00:00:00Z');

function update() {
    const now = new Date();
    
    const diff = now - startDate;
    const millisecondsInYear = 365.25 * 24 * 60 * 60 * 1000;

    let years = diff / millisecondsInYear;
    years = years.toFixed(9);

    document.getElementById("timer").textContent = `${years}`;
}

setInterval(update, 10)