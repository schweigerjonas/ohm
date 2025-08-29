window.onload = () => {
  const form = document.getElementById("transaction-form");
  const typeSelect = document.getElementById("type-select");

  form.addEventListener("htmx:configRequest", (event) => {
    if (event.detail.elt === form) {
      const selectedType = typeSelect.value;

      if (selectedType === "income") {
        event.detail.path = "/api/income";
      } else {
        event.detail.path = "/api/expense";
      }
    }
  });
};
