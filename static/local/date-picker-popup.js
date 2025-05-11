// used by date range picker test page
document.getElementById("check-date-button")?.addEventListener("click", function () {
  let html = `
    <form id="check-availability-form" action="" method="post" novalidate class="needs-calidation">
      <div class="row">
        <div class="col">
          <div class="row" id="reservation-date-modal">
            <div class="col"></div>
            <div class="col">
              <input disabled required class="form-control" type="text" name="selected_date" id="selected_date" placeholder="Select Date">
            </div>
            <div class="col"></div>
          </div>
        </div>
      </div>
    </form>
  `;
  attention.customHtml({
    title: "Choose Your Date",
    msg: html,
    willOpen: () => {
      const elem = document.querySelector('input[name="selected_date"]');
      const datepicker = new Datepicker(elem, {
        format: "yyyy-mm-dd",
        minDate: new Date(), // remove this if you need access to past dates
      });
    },
    didOpen: () => {
      document.getElementById('selected_date').removeAttribute('disabled');
    },
    callback: function(formValues) {
      let form = document.getElementById("check-availability-form");
      let formData = new FormData(form);
      formData.append("csrf_token", "{{ .CSRFToken }}");
      fetch('/availability-json', {
        method: "post",
        body: formData,
      })
      .then(response => response.json())
      .then(data => {
        console.log(data);
        console.log(data.ok);
        console.log(data.message);
      })
    }
  });
})