package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/math2001/boatsinker/server/app"
	"github.com/math2001/boatsinker/server/em"
)

const PORT = 9999

type Message struct {
	Data map[string]interface{}
	From *net.Conn
}

func main() {
	addr := fmt.Sprintf("0.0.0.0:%d", PORT)
	fmt.Printf("Listening on http://%s\n", addr)
	http.Handle("/", http.FileServer(http.Dir("dist")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w, nil)
		if err != nil {
			log.Printf("Couldn't upgrade HTTP to WebSocket: %s", err)
			return
		}
		// read from socket and emit an event as soon as there is a message
		em.Emit("connection.new", conn)
		// read new messages and emits them as events
		go func() {
			defer conn.Close()
			var (
				reader  = wsutil.NewReader(conn, ws.StateServerSide)
				decoder = json.NewDecoder(reader)
			)
			// read from the connection forever and close when there is an error
			for {
				header, err := reader.NextFrame()
				if err != nil {
					em.Emit("connection.error", conn)
					return
				}
				if header.OpCode == ws.OpClose {
					em.Emit("connection.closed", conn)
					return
				}

				var msg Message
				msg.From = &conn
				msg.Data = make(map[string]interface{})

				if err := decoder.Decode(&msg.Data); err != nil {
					log.Printf("Error occured while parsing WebSocket: %s", err)
					return
				}

				em.Emit("connection.msg", msg)

			}
		}()
	})
	app.Start()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
