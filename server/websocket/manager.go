package websocket

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

type Manager struct {
	clients ClientList
	sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{clients: make(ClientList)}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (manager *Manager) ServeWs(responseWriter http.ResponseWriter, request *http.Request) {

	conn, err := upgrader.Upgrade(responseWriter, request, nil)

	if err != nil {
		log.Println("Error occurred during connection", err)
		return
	}

	client := NewClient(conn, manager)
	manager.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (manager *Manager) addClient(client *Client) {
	manager.Lock()
	defer manager.Unlock()

	manager.clients[client] = true
	log.Println("Added Client nth: " + strconv.Itoa(len(manager.clients)))
}

func (manager *Manager) removeClient(client *Client) {
	manager.Lock()
	defer manager.Lock()

	if _, ok := manager.clients[client]; ok {
		client.connection.Close()
		delete(manager.clients, client)
	}
}
