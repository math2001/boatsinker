package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

const PORT = 9999

func main() {
	addr := fmt.Sprintf("localhost:%d", PORT)
	fmt.Println("Listening on", addr)
	http.Handle("/", http.FileServer(http.Dir("../public")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w, nil)
		if err != nil {
			log.Printf("Couldn't upgrade HTTP to WebSocket: %s", err)
			return
		}
		// read from socket and emit an event as soon as there is a message
		conn, ok := e.(net.Conn)
		if !ok {
			log.Fatal("Invalid argument for 'new socket'.")
		}
		EM.Emit("connection.new", conn)
		// read new messages and emits them as events
		go func() {
			defer conn.Close()
			var (
				r       = wsutil.NewReader(conn, ws.StateServerSide)
				decoder = json.NewDecoder(r)
			)
			// read from the connection forever and close when there is an error
			for {
				header, err := conn.NextFrame()
				if err != nil {
					EM.Emit("connection.error", conn)
					return
				}
				if header.OpCode == ws.OpClose {
					EM.Emit("connection.closed", conn)
					return
				}

				var msg Message
				msg.Data = make(map[string]interface{})

				if err := decoder.Decode(&msg.Data); err != nil {
					log.Printf("Error occured while parsing WebSocket: %s", err)
					return
				}

				EM.Emit("connection.msg", msg)

			}
		}()
	})
	game.Start()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
