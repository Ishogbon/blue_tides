package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection,
		manager,
	}
}

func (client *Client) readMessages() {
	defer func() {

	}()

	for {
		messageType, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected Close Error %v", err)
			}
			break
		}
		log.Print(messageType)
		log.Print(payload)
	}
}
