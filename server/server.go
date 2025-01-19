package server

import (
	"log"
	"net/http"
)

type ReturnData struct {
	message string
}

func imAlive(w http.ResponseWriter, r *http.Request) {

}

func StartServer() {
	http.HandleFunc("/health", imAlive)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
