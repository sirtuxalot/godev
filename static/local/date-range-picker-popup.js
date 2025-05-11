// used by date range picker test page
document.getElementById("check-dates-button")?.addEventListener("click", function () {
  let html = `
    <form id="check-availability-form" action="" method="post" novalidate class="needs-calidation">
      <div class="row">
        <div class="col">
          <div class="row" id="reservation-dates-modal">
            <div class="col">
              <input disabled required class="form-control" type="text" name="arrival_date" id="arrival_date" placeholder="Arrival Date">
            </div>
            <div class="col">
              <input disabled required class="form-control" type="text" name="departure_date" id="departure_date" placeholder="Departure Date">
            </div>
          </div>
        </div>
      </div>
    </form>
  `;
  attention.customHtml({
    title: "Choose Your Dates",
    msg: html,
    willOpen: () => {
      const elem = document.getElementById('reservation-dates-modal');
      const rangePicker = new DateRangePicker(elem, {
        format: "yyyy-mm-dd",
        minDate: new Date(), // remove this if you need access to past dates
      });
    },
    didOpen: () => {
      document.getElementById('arrival_date').removeAttribute('disabled');
      document.getElementById('departure_date').removeAttribute('disabled');
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
