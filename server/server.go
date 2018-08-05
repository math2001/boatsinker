package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/math2001/boatsinker/server/app"
	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
)

const PORT = 9999

// the server manages the "raw" connection. It doesn't know anything about the game
// It could be used in a whole different application.
// It triggers and listen for various events:
// - connection.new -> got new connection
// - connection.msg -> got new message
// - connection.error -> got error while reading
// To send a message, just emit the event 'connection.send' with a Message
// set .Data to be the data, and .From to be the connection you want to send to.
// .Count should be 0 (this is the default). As soon as the message is send, it is set 1.
// In order to make sure that the connection has been found and your message sent, you might want
// to make sure that the it actually is set to 1

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
				writer  = wsutil.NewWriter(conn, ws.StateServerSide, ws.OpText)
				reader  = wsutil.NewReader(conn, ws.StateServerSide)
				decoder = json.NewDecoder(reader)
				encoder = json.NewEncoder(writer)
			)
			em.On("connetion.send", func(e interface{}) error {
				msg, ok := e.(utils.Message)
				if !ok {
					log.Print("Invalid data to send. Should be a Message")
					panic(nil)
				}
				if msg.Count != 0 {
					log.Printf("Message .Count should be 0, got %d", msg.Count)
					panic(nil)
				}
				if msg.From == &conn {
					// this is our connection, *we* have the writer
					if err := encoder.Encode(msg.Data); err != nil {
						return fmt.Errorf("Couldn't encode message and write to connection with %v",
							msg.Data)
					}
				}
				return nil
			})
			// read from the connection forever and close when there is an error
			messagecount := 1
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

				var msg utils.Message
				msg.From = &conn
				msg.Data = make(map[string]interface{})
				msg.Count = messagecount

				if err := decoder.Decode(&msg.Data); err != nil {
					log.Printf("Error occured while parsing WebSocket: %s", err)
					return
				}

				em.Emit("connection.msg", msg)

				messagecount += 1
			}
		}()
	})
	app.Start()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
