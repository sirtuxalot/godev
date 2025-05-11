/* internal/handlers/access.go */

package handlers

import (
	"fmt"
	"godev/internal/auth"
	"godev/internal/forms"
	"godev/internal/models"
	"godev/internal/render"
	"log"
	"net/http"
)

// send user to login form
func (handlers *Repository) LogIn(writer http.ResponseWriter, request *http.Request) {
	// open login template without form data
	render.Template(writer, request, "login.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// log out user
func (handlers *Repository) LogOut(writer http.ResponseWriter, request *http.Request) {
	// destroy and renew session tokens
	_ = handlers.App.Session.Destroy(request.Context())
	_ = handlers.App.Session.RenewToken(request.Context())

	// redirect user back to login template
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}

func (handlers *Repository) PostLogIn(writer http.ResponseWriter, request *http.Request) {
	// renew session tokens
	_ = handlers.App.Session.RenewToken(request.Context())

	// parse provided form data
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}

	// retrieve and validate form data and redirect back to form if not valid
	form := forms.New(request.PostForm)
	form.Required("user-id", "password")
	if !form.Valid() {
		render.Template(writer, request, "login.html", &models.TemplateData{
			Form: form,
		})
		return
	}

	// user data struct
	user := auth.UserLogin{
		Login_ID:  request.Form.Get("user_id"),
		Login_PWD: request.Form.Get("password"),
	}

	// bind to ldap/ad server
	ldapConn, err := auth.Connect()
	if err != nil {
		log.Fatal("LDAP Bind failed!")
	}
	defer ldapConn.Close()

	// authenticate user
	authVerified, authErr := auth.Auth(ldapConn, user)
	if authErr != nil {
		log.Fatal("Authentication failed to verify user!")
	}

	// verify authenticated user
	if authVerified {
		fmt.Fprintf(writer, "Welcome, %s!", user.Login_ID)
	} else {
		http.Error(writer, "Invalid credentials", http.StatusUnauthorized)
	}

	// flash success and redirect to landing page
	//handlers.App.Session.Put(request.Context(), "user_id", id)
	handlers.App.Session.Put(request.Context(), "flash", "Logged In Successfully!")
	http.Redirect(writer, request, "/login", http.StatusSeeOther)
}
