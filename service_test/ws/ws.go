package ws

import "github.com/gorilla/websocket"

var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Receive struct {
	MessageType string `json:"type"`
	Data        any    `json:"data"`
}
type Output struct {
	MessageType string `json:"type"`
	Data        any    `json:"data"`
}
