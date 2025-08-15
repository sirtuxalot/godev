/* cmd/web/main.go */

package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"godev/internal/config"
	"godev/internal/database"
	"godev/internal/handlers"
	"godev/internal/render"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// access application config
var app config.AppConfig
var session *scs.SessionManager

func main() {
	appDB, err := run()
	if err != nil {
		log.Fatal(err)
	}

	// close access to database
	appDB.SQL.Close()

	// start the application server and set the port
	log.Println(fmt.Sprintf("Starting application on port %s", app.Port_Number))

	// create server (appServer), setting portnumber and routes
	appServer := &http.Server{
		Addr:    app.Port_Number,
		Handler: routes(&app),
	}

	// start http server
	err = appServer.ListenAndServe()
	log.Fatal(err)
}

func run() (*database.DB, error) {
	// read .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// read environment variables
	InProduction, err := strconv.ParseBool(os.Getenv("InProduction"))
	if err != nil {
		log.Fatal(err)
	}
	Port_Number := os.Getenv("Port_Number")
	DB_Name := os.Getenv("DB_Name")
	DB_User := os.Getenv("DB_User")
	DB_Password := os.Getenv("DB_Password")
	DB_Host := os.Getenv("DB_Host")
	DB_Port := os.Getenv("DB_Port")
	DB_SSL := os.Getenv("DB_SSL")
	LDAP_URL := os.Getenv("LDAP_URL")
	BASE_DN := os.Getenv("BASE_DN")
	ID_LVL := os.Getenv("ID_LVL")
	LDAP_ID := os.Getenv("LDAP_ID")
	LDAP_PWD := os.Getenv("LDAP_PWD")
	USERNAME := os.Getenv("USERNAME")
	USER_OU := os.Getenv("USER_OU") + "," + os.Getenv("BASE_DN")
	GROUP_CN := os.Getenv("APP_GROUP_CN") + "," + os.Getenv("GROUP_OU") + "," + os.Getenv("BASE_DN")

	if !InProduction {
		log.Println("+-------------------------------------------------+")
		log.Printf("| godotenv: In Production = %t", InProduction)
		log.Printf("| godotenv: UseCache = %t", InProduction)
		log.Println("+-------------------------------------------------+")
		log.Printf("| godotenv: Application Port = %s", Port_Number)
		log.Println("+-------------------------------------------------+")
		log.Printf("| godotenv: Database Name = %s", DB_Name)
		log.Printf("| godotenv: Database User = %s", DB_User)
		log.Printf("| godotenv: Database Password = %s", DB_Password)
		log.Printf("| godotenv: Database Host = %s", DB_Host)
		log.Printf("| godotenv: Database Port = %s", DB_Port)
		log.Printf("| godotenv: Database SSL Mode = %s", DB_SSL)
		log.Println("+--------------------------------------------------+")
		log.Printf("| godotenv: LDAP URL = %s", LDAP_URL)
		log.Printf("| godotenv: Base DN = %s", BASE_DN)
		log.Printf("| godotenv: User OU = %s", USER_OU)
		log.Printf("| godotenv: Appliation Group CN = %s", GROUP_CN)
		log.Println("+--------------------------------------------------+")
	}

	// application production status
	app.InProduction = InProduction
	app.UseCache = InProduction
	app.Port_Number = Port_Number
	app.LDAP_URL = LDAP_URL
	app.BASE_DN = BASE_DN
	app.ID_LVL = ID_LVL
	app.LDAP_ID = LDAP_ID
	app.LDAP_PWD = LDAP_PWD
	app.USERNAME = USERNAME
	app.USER_OU = USER_OU
	app.GROUP_CN = GROUP_CN

	// initialize session management
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	connectString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password='%s' sslmode=%s", DB_Host, DB_Port, DB_Name, DB_User, DB_Password, DB_SSL)
	appDB, err := database.ConnectSQL(connectString)
	if err != nil {
		log.Fatal("Unable to connect to database! Closing application!")
	}
	log.Println("Connected to database...")

	// initiate creation of template cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Unable to create template cache!")
	}

	// set values with application config
	app.TemplateCache = templateCache

	// apply application config to variable (repo) and pass it back to handlers package
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// add template cache data to application config and send to render package
	render.NewTemplates(&app)

	return appDB, nil
}
