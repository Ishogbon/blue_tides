package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager
	// egress is a channel used to avoid concurrent writes out of the websocket connection
	egress chan []byte
}

// Using this is non conventional, it is moe appropriate to use the first letter instead e.g c instead of connection
func NewClient(connection *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection,
		manager,
		make(chan []byte),
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

func (client *Client) writeMessages() {
	defer func() {
		client.manager.removeClient(client)
	}()

	// This code will exit if all the channels used are closed, which is set to nil, in this case 1
	for count := 0; count < 1; {
		select {
		case messsage, ok := <-client.egress:
			if !ok {
				if err := client.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					// This case will never be trigerred again.
					client.egress = nil
					count++
					log.Println("connection failed", err)
				}
				return
			}
			if err := client.connection.WriteMessage(websocket.TextMessage, messsage); err != nil {
				log.Println("error occured sending message to client", err)
			}
			log.Println("Message sent to client")
		}
	}
}
