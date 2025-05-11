/* internal/config/config.go */

package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// application config
type AppConfig struct {
	ErrorLog      *log.Logger
	InProduction  bool
	Port_Number   string
	InfoLog       *log.Logger
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
	UseCache      bool
	LDAP_URL      string
	BASE_DN       string
	ID_LVL        string
	LDAP_ID       string
	LDAP_PWD      string
	USERNAME      string
	USER_OU       string
	GROUP_CN      string
}
