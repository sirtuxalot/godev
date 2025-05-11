let attention = Prompt();

// validates fields on forms
(function() {
  'use strict';
  // Fetch all forms we want to apply custom bootstrap validation styles to
  const forms = document.querySelectorAll('.needs-validation');
  // Loop over them and prevent submission
  Array.from(forms).forEach(form => {
    form.addEventListener('submit', event => {
      if (!form.checkValidity()) {
        event.preventDefault();
        event.stopPropagation();
      }
      form.classList.add('was-validated')
    }, false);
  });
})();

function Prompt() {
  let toast = function(icon, title) {
    const Toast = Swal.mixin({
      toast: true,
      position: "top-end",
      title: title,
      icon: icon,
      showConfirmButton: false,
      timer: 3000,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.onmouseenter = Swal.stopTimer;
        toast.onmouseleave = Swal.resumeTimer;
      }
    });
    Toast.fire({});;
  }

  // examples @ https://sweetalert2.github.io/#usage
  let deleteModal = function(title, text, icon, confirmButtonText) {
    Swal.fire({
      title: title,
      text: text,
      icon: icon,
      showCancelButton: true,
      confirmButtonColor: "#3085d6",
      cancelButtonColor: "#d33",
      confirmButtonText: confirmButtonText,
    }).then((result) => {
    if (result.isConfirmed) {
      Swal.fire({
        title: "Deleted!",
        text: "Your file has been deleted.",
        icon: "success"
      });
    }
  });
  }

  // examples @ https://sweetalert2.github.io/#usage
  let notifyModal = function(title, text, icon, confirmButtonText) {
    Swal.fire({
      title: title,
      html: text,
      // icon enum: ["success", "error", "warning", "info", "question"]
      icon: icon,
      confirmButtonText: confirmButtonText
    })
  }

  let notify = function(type, msg) {
    notie.alert({
      // type enum: ["success", "warning", "error", "info", "neutral"]
      type: type,
      text: msg,
    })
  }

  async function customHtml(custom) {
    const {
      icon ="",
      msg = "",
      title = "",
      showConfirmButton = true,
    } = custom;
    const { value: formValues } = await Swal.fire({
      icon: icon,
      title: title,
      html: msg,
      backdrop: false,
      focusConfirm: false,
      showCancelButton: true,
      showConfirmButton: showConfirmButton,
      willOpen: () => {
        if (custom.willOpen !== undefined) {
          custom.willOpen();
        }
      },
      preConfirm: () => {
        return [
          document.getElementById('arrival_date').value,
          document.getElementById('departure_date').value,
        ]
      },
      didOpen: () => {
        if (custom.didOpen !== undefined) {
          custom.didOpen();
        }
      },
    })
    if (formValues) {
      if (formValues.dismiss !== Swal.DismissReason.cancel) {
        if (formValues.value !== "") {
          if (custom.callback !== undefined) {
            custom.callback(formValues);
          }
        } else {
          custom.callback(false);
        }
      } else {
        custom.callback(false);
      }
    }
  }

  return {
    toast: toast,
    deleteModal: deleteModal,
    notifyModal: notifyModal,
    notify: notify,
    customHtml: customHtml,
  }

}
