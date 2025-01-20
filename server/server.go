package server

import (
	"encoding/json"
	"log"
	"net/http"
)

type Data struct{}

type Response struct {
	Status  string
	Message string
	Data    Data
}

func (rd *Response) writeJsonResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(rd); err != nil {
		http.Error(w, "Failed To Encode Error", http.StatusInternalServerError)
	}
}

func imAlive(w http.ResponseWriter, r *http.Request) {
	var returnData Response = Response{
		"success", "Ping was successfull", Data{},
	}
	returnData.writeJsonResponse(w, 200)
}

func StartServer() {
	http.HandleFunc("/health", imAlive)
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
