package main

import (
    "log"
    "net/http"
	"bakery/routes" // Update this line of code
	"bakery/adapters" // Update this line of code
	"bakery/controller" // Update this line of code
)

var server *http.Server


func RunServer() {

	mux := http.NewServeMux()
	SetupHandlers(mux)


	server = &http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func SetupHandlers(mux *http.ServeMux) {

	mux.HandleFunc("/home", routes.HomeHandler)
	controller.SetupRecipeHandler(mux)
}

func main() {
	adapters.InitMongo()
	RunServer()
}