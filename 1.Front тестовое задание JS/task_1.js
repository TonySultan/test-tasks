const tableBody = document.querySelector("table tbody");

async function fetchPage(page) {
  try {
    const response = await fetch(`https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=${page}`);
    const data = await response.json();
    return data;
  } catch (error) {
    throw error;
  }
}
async function addDataToTable() {
  let page = 1;
  let index = 0;

  while (true) {
    const data = await fetchPage(page);

    if (data.length === 0) {
      break;
    }

    data.forEach(currency => {
      const row = document.createElement("tr");
      row.innerHTML = `
        <td class="${(index <= 5 && currency.symbol != "usdt" ) ? "blue-bg"  : (currency.symbol === "usdt" ? "green-bg" : "")}">${currency.id}</td>
        <td class="${(index <= 5 && currency.symbol != "usdt" ) ? "blue-bg"  : (currency.symbol === "usdt" ? "green-bg" : "")}">${currency.symbol}</td>
        <td class="${(index <= 5 && currency.symbol != "usdt" ) ? "blue-bg"  : (currency.symbol === "usdt" ? "green-bg" : "")}">${currency.name}</td>
      `;
      tableBody.appendChild(row);
      index++;
    });

    page++;
  }
}

addDataToTable().catch(error => {
  console.error("Произошла ошибка при получении данных:", error);
});
