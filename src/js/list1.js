// untuk modal
const exampleModal = document.getElementById("exampleModal");
if (exampleModal) {
  exampleModal.addEventListener("show.bs.modal", (event) => {
    // Button that triggered the modal
    const button = event.relatedTarget;
    // Extract info from data-bs-* attributes
    const recipient = button.getAttribute("data-bs-whatever");
    // If necessary, you could initiate an Ajax request here
    // and then do the updating in a callback.

    // Update the modal's content.
    const modalTitle = exampleModal.querySelector(".modal-title");
    const modalBodyInput = exampleModal.querySelector(".modal-body input");

    modalTitle.textContent = "Beli Sekarang!";
    modalBodyInput.value = recipient;
  });
}

// untuk alert
document
  .getElementById("purchaseForm")
  .addEventListener("submit", function (event) {
    event.preventDefault(); // Prevent the form from submitting normally
    var myModalEl = document.getElementById("exampleModal");
    var modal = bootstrap.Modal.getInstance(myModalEl); // Get the Bootstrap modal instance
    modal.hide(); // Hide the modal

    // Display alert in the center of the screen
    var alertDiv = document.createElement("div");
    alertDiv.className = "alert alert-success text-center";
    alertDiv.style.position = "fixed";
    alertDiv.style.top = "50%";
    alertDiv.style.left = "50%";
    alertDiv.style.transform = "translate(-50%, -50%)";
    alertDiv.style.zIndex = "1050"; // Higher than modal backdrop
    alertDiv.innerHTML = "Pembelian berhasil!";

    document.body.appendChild(alertDiv);
    setTimeout(function () {
      alertDiv.classList.add("fade-out");
    }, 2000); // Start fade out after 2 seconds

    setTimeout(function () {
      document.body.removeChild(alertDiv);
      // Optionally, submit the form data via AJAX here
      // Example:
      // var formData = new FormData(event.target);
      // fetch('/penjualan', { method: 'POST', body: formData });
    }, 3000); // Remove alert after 3 seconds
  });
