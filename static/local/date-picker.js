// used by date picker test page
const elem = document.querySelector('input[name="date"]');
const datepicker = new Datepicker(elem, {
  format: "yyyy-mm-dd",
  minDate: new Date(), // remove this if you need access to past dates
});
