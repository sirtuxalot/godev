/* internal/helpers/helpers.go */

package helpers

import (
  "fmt"
  "godev/internal/config"
  "net/http"
  "runtime/debug"
)

var app *config.AppConfig

// setup app config for helpers
func NewHelpers(appConfig *config.AppConfig) {
  app = appConfig
}

func ClientError(writer http.ResponseWriter, status int) {
  app.InfoLog.Println("Client Error with status of", status)
  http.Error(writer, http.StatusText(status), status)
}

func ServerError(writer http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.ErrorLog.Println(trace)
  http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func IsAuthenticated(request *http.Request) bool {
  exists := app.Session.Exists(request.Context(), "user_id")
  return exists
}
