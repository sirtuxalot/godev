// used by date range picker test page
const elem = document.getElementById('range-picker-test');
const rangepicker = new DateRangePicker(elem, {
  format: "yyyy-mm-dd",
  minDate: new Date(), // remove this if you need access to past dates
});
