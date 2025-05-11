/* internal/handlers/handlers.go */

package handlers

import (
  "godev/internal/config"
)

// provide access to application settings
var app *config.AppConfig

// repository variable used by the handlers
var Handlers *Repository

// repository type struct
type Repository struct {
  App *config.AppConfig
}

// creates a new examples repository
func NewRepo(appConfig *config.AppConfig) *Repository {
  return &Repository{
    App: appConfig,
  }
}

// sets the examplessitory for the handlers
func NewHandlers(handlers *Repository) {
  Handlers = handlers
}
