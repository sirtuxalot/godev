/* cmd/web/middleware.go */

package main

import (
  "github.com/justinas/nosurf"
  "godev/internal/helpers"
  "log"
  "net/http"
)

func CSRF_Failure() http.Handler {
  return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    log.Printf("Request Failed. Reason: %v", nosurf.Reason(request))
    http.Error(writer, http.StatusText(nosurf.FailureCode), nosurf.FailureCode)
  })
}

// sets the CSRF token to improve site security
func NoSurf(next http.Handler) http.Handler {
  csrfHandler := nosurf.New(next)
  csrfHandler.SetBaseCookie(http.Cookie{
    HttpOnly: true,
    Path:     "/",
    Secure:   app.InProduction,
    SameSite: http.SameSiteLaxMode,
  })
  csrfHandler.SetFailureHandler(CSRF_Failure())
  return csrfHandler
}

// losds and save session on request
func SessionLoad(next http.Handler) http.Handler {
  return session.LoadAndSave(next)
}

// authentication function
func Auth(next http.Handler) http.Handler {
  return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
    if !helpers.IsAuthenticated(request) {
      session.Put(request.Context(), "error", "Log in first!")
      http.Redirect(writer, request, "/login", http.StatusSeeOther)
      return
    }
    next.ServeHTTP(writer, request)
  })
}
