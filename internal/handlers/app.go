/* internal/handlers/app.go */

package handlers

import (
	"godev/internal/models"
	"godev/internal/render"
	"net/http"
)

// handler function for index page
func (handlers *Repository) Open(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "list.html", &models.TemplateData{})
}

// handler function for index page
func (handlers *Repository) Closed(writer http.ResponseWriter, request *http.Request) {
	render.Template(writer, request, "list.html", &models.TemplateData{})
}
