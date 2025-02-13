package websocket

import (
	"blue_tides/handler"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

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

	// I removed the for loop to prevent multiple movie plays attempt, considering the fact for this test, only needs be done once.
	// for {
	_, payload, err := client.connection.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("Unexpected Close Error %v", err)
		}
		// break
	}
	movieName := string(payload)

	// Norms this should have been abstracted away into it's own function
	movie := handler.Movie{}

	frameLoadTimeChannel := make(chan time.Duration)

	filePath := filepath.Join(os.Getenv("MOVIE_DIRECTORY"), movieName+".timestamp.log")
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Printf("failed to create directory: %v\n", err)
		return
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	go func() {
		frameLoadTime := <-frameLoadTimeChannel
		if err != nil {
			fmt.Printf("failed to open file: %v\n", err)
			return
		}
		defer file.Close()

		// Write the data to the file
		if _, err := file.WriteString(fmt.Sprintf("%v\n", frameLoadTime)); err != nil {
			fmt.Printf("failed to write to file: %v\n", err)
			return

		}

	}()
	movie.PlayMovie(os.Getenv("MOVIE_DIRECTORY"), movieName, func(movieByteFrame []byte, frameLoadDuration time.Duration) {
		client.egress <- movieByteFrame
		frameLoadTimeChannel <- frameLoadDuration
	})
	// }
}

func (client *Client) writeMessages() {
	defer func() {
		close(client.egress)
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
