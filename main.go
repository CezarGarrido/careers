package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CezarGarrido/careers/driver"
	httpHandler "github.com/CezarGarrido/careers/handlers/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// const (
// 	host     = "localhost"
// 	portDB   = "5432"
// 	user     = "postgres"
// 	password = "C102030g"
// 	dbname   = "app_advogados"
// )

func main() {

	port := map[bool]string{true: os.Getenv("PORT"), false: "8084"}[os.Getenv("PORT") != ""]

	// postgresSettings := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portDB, user, password, dbname)

	databaseURL, exists := os.LookupEnv("DATABASE_URL")
	if !exists {
		log.Panic("Variable DATABASE_URL not found")
	}

	pgConnection, err := driver.OpenPostgres(databaseURL)
	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()

	superHeroHandlerHttp := httpHandler.NewSuperHero(pgConnection)

	r.HandleFunc("/api/v1/superhero", superHeroHandlerHttp.Create).Methods("POST")

	headersOk := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Cache-Control", "X-File-Name", "Origin", "X-Session-ID"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE"})

	log.Println("Start app in port:", port)

	err = http.ListenAndServe(":"+port, handlers.CORS(headersOk, methodsOk, originsOk)(r))
	if err != nil {
		log.Panic(err)
	}
}
