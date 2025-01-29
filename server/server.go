package server

import (
	"blue_tides/server/websocket"
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

func (rd *Response) WriteJsonResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if json.NewEncoder(w).Encode(rd) != nil {
		http.Error(w, "Failed To Encode Error", http.StatusInternalServerError)
	}
}

func imAlive(w http.ResponseWriter, r *http.Request) {
	var returnData Response = Response{
		"success", "Ping was successfull", Data{},
	}
	returnData.WriteJsonResponse(w, 200)
}

func StartServer() {
	http.HandleFunc("/health", imAlive)
	manager := websocket.NewManager()
	http.HandleFunc("/ws", manager.ServeWs)
	log.Println("Server running on port 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
