document.getElementById("success-modal-sw2-button").addEventListener("click", function() {
  attention.notifyModal("My Title", "Hello World", "success", "Cool Button Text");
})

document.getElementById("warning-modal-sw2-button").addEventListener("click", function() {
  attention.notifyModal("Look Out", "<em>Careful!</em>", "warning", "Oh No!");
})

document.getElementById("danger-modal-sw2-button").addEventListener("click", function() {
  attention.notifyModal(title="BORKED!", text="<strong>Oppps!</strong>", icon="error", confirmButtonText="You broke something!!!");
})


document.getElementById("success-toast-sw2-button").addEventListener("click", function() {
  attention.toast("success", "Success message")
})

document.getElementById("warning-toast-sw2-button").addEventListener("click", function() {
  attention.toast("warning", "Warning message")
})

document.getElementById("danger-toast-sw2-button").addEventListener("click", function() {
  attention.toast("error", "Danger message")
})
