package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CezarGarrido/careers/driver"
	httpHandler "github.com/CezarGarrido/careers/handlers/http"
	"github.com/CezarGarrido/careers/migrate"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	port := map[bool]string{true: os.Getenv("PORT"), false: "8084"}[os.Getenv("PORT") != ""]

	databaseURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))

	pgConnection, err := driver.OpenPostgres(databaseURL)
	if err != nil {
		log.Panic(err)
	}

	err = migrate.New("postgres", "./migrations", pgConnection.SQL)
	if err != nil {
		log.Panic(err)
	}

	r := mux.NewRouter()

	superHeroHandlerHttp := httpHandler.NewSuperHero(
		pgConnection,
		"https://superheroapi.com/api",
		"1426431747558534",
	)

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
