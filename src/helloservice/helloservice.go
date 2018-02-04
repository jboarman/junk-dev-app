package helloservice

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	log.Print("Initializing helloservice")
	http.HandleFunc("/hi", handle)
	http.HandleFunc("/_health", healthCheckHandler)
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Print("`Hello` sent to client")
	fmt.Fprint(w, "Hello world!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("`OK` Health Check sent to client")
	fmt.Fprint(w, "ok, we gotta clean bill of health")
}
