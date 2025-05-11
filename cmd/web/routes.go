/* cmd/web/routes.go */

package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"godev/internal/config"
	"godev/internal/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	// create application router
	mux := chi.NewRouter()

	// add middleware
	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// create GET routes
	mux.Get("/", handlers.Handlers.Index)
	mux.Get("/about", handlers.Handlers.About)
	mux.Get("/date-picker1", handlers.Handlers.DatePicker)
	mux.Get("/date-picker2", handlers.Handlers.DateRangePicker)
	mux.Get("/date-picker3", handlers.Handlers.DatePickerPopUp)
	mux.Get("/date-picker4", handlers.Handlers.DateRangePickerPopUp)
	mux.Get("/notie", handlers.Handlers.Notie)
	mux.Get("/sweet-alert2", handlers.Handlers.SweetAlert)
	mux.Get("/login", handlers.Handlers.LogIn)
	mux.Get("/logout", handlers.Handlers.LogOut)

	// create POST routes
	mux.Post("/login", handlers.Handlers.PostLogIn)

	// create "list" routes
	/*
	   mux.Route("/list", func mux chi.Router) {
	     if app.InProduction {
	       mux.Use(Auth)
	     }
	*/
	// create Handlers GET routes
	mux.Get("/open", handlers.Handlers.Open)
	mux.Get("/closed", handlers.Handlers.Closed)
	/*
	   })
	*/

	// provide access to static files
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// provide the router to the application
	return mux
}
