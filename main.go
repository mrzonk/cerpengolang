package main

import (
	"cerpengolang/config"
	"cerpengolang/routes"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running")
	//http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r)
	http.ListenAndServe(":3434", r)
}
