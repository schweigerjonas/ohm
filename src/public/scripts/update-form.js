window.onload = () => {
  const form = document.getElementById("transaction-form");
  const typeSelect = document.getElementById("type-select");

  typeSelect.addEventListener("change", (event) => {
    const selectedType = event.target.value;
    if (selectedType === "income") {
      form.setAttribute("hx-post", "/api/income");
    } else {
      form.setAttribute("hx-post", "/api/expense");
    }
  });
};
