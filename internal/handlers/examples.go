/* internal/handlers/examples.go */

package handlers

import (
  "godev/internal/models"
  "godev/internal/render"
  "net/http"
)

// handler function for index page
func (handlers *Repository) Index(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "index.html", &models.TemplateData{})
}

// handler function for date picker test page
func (handlers *Repository) DatePicker(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "date-picker.html", &models.TemplateData{})
}

// handler function for date picker test page
func (handlers *Repository) DateRangePicker(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "date-range-picker.html", &models.TemplateData{})
}

// handler function for date picker test page
func (handlers *Repository) DatePickerPopUp(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "date-picker-popup.html", &models.TemplateData{})
}

// handler function for date picker test page
func (handlers *Repository) DateRangePickerPopUp(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "date-range-picker-popup.html", &models.TemplateData{})
}

// handler function for notie test page
func (handlers *Repository) Notie(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "notie.html", &models.TemplateData{})
}

// handler function for index page
func (handlers *Repository) SweetAlert(writer http.ResponseWriter, request *http.Request) {
  render.Template(writer, request, "sweet-alert2.html", &models.TemplateData{})
}

// handler function for about page
func (handlers *Repository) About(writer http.ResponseWriter, request *http.Request) {
  stringMap := make(map[string]string)
  stringMap["test"] = "Hello Again!"
  render.Template(writer, request, "about.html", &models.TemplateData{
    StringMap: stringMap,
  })
}
