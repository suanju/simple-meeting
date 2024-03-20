package main

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	server := socketio.NewServer(nil)
	server.OnEvent("/", "test", func(conn socketio.Conn, msg string) error {
		fmt.Println(msg)
		conn.Emit("test", "test 123")
		return nil
	})
	go func() {
		_ = server.Serve()
	}()
	defer func(server *socketio.Server) {
		_ = server.Close()
	}(server)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})
	r := mux.NewRouter()
	r.Handle("/socket.io/", server)
	r.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(123123)
	})
	r.Use(c.Handler)

	http.Handle("/", r)
	log.Println("Serving at localhost:3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
