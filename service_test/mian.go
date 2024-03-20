package main

import (
	"encoding/json"
	websocket2 "github.com/gorilla/websocket"
	"log"
	"net/http"
	"simple-meeting/ws"
)

type RoomT map[string][]*websocket2.Conn

var room = make(RoomT, 0)

func main() {
	http.HandleFunc("/ws", websocket)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatalln("service run error")
	}
}

func websocket(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	conn, _ := ws.Upgrade.Upgrade(w, r, nil)
	id := r.URL.Query().Get("room_id")
	if len(room) <= 0 {
		_ = conn.WriteJSON(&ws.Output{
			MessageType: "error",
			Data:        "room_id is not definition",
		})
		_ = conn.Close()
		return
	}
	go listeningConn(conn)
	room[id] = append(room[id], conn)
}

func listeningConn(conn *websocket2.Conn) {
	for {
		_, p, _ := conn.ReadMessage()
		message := new(ws.Receive)
		err := json.Unmarshal(p, message)
		if err != nil {
			_ = conn.WriteJSON(&ws.Output{
				MessageType: "123",
				Data:        "12123123123",
			})
		}
	}
}
