package websocket

import "github.com/gorilla/websocket"

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
