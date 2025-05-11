document.getElementById("success-notie-button").addEventListener("click", function() {
  attention.notify("success", "Success message")
})

document.getElementById("warning-notie-button").addEventListener("click", function() {
  attention.notify("warning", "Warning message")
})

document.getElementById("danger-notie-button").addEventListener("click", function() {
  attention.notify("error", "Danger message")
})
