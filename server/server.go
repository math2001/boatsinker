package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/math2001/boatsinker/server/app"
	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
)

const port = 9999

// the server manages the "raw" connection. It doesn't know anything about the game
// It could be used in a whole different application.
// Every connection.* events is managed my him (either triggered or listened)

type connection struct {
	closed  bool
	raw     *net.Conn
	encoder *json.Encoder
	decoder *json.Decoder
	writer  *wsutil.Writer
	reader  *wsutil.Reader
}

func (c connection) String() string {
	var b strings.Builder
	fmt.Fprint(&b, "connection ")
	if c.closed {
		fmt.Fprintf(&b, "'closed'")
	} else {
		fmt.Fprintf(&b, "'open'")
	}
	return b.String()
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Listening on http://%s\n", addr)
	http.Handle("/", http.FileServer(http.Dir("dist")))

	var conns []connection

	em.On("connection.send", func(e interface{}) error {
		msg, ok := e.(utils.Message)
		if !ok {
			panic(fmt.Sprintf("Invalid data to send. Should have utils.Message, got %T", e))
		}
		for _, conn := range conns {
			if conn.raw == msg.From {
				if err := conn.encoder.Encode(msg.Data); err != nil {
					return fmt.Errorf("Couldn't encode message and write to socket\n%s",
						err)
				}
				if err := conn.writer.Flush(); err != nil {
					return fmt.Errorf("Couldn't flush the writer\n%s", err)
				}
				return nil
			}
		}
		// this shouldn't happen in any case.
		panic("Couldn't find connection to send message to.")
	})

	em.On("connection.close", func(e interface{}) error {
		conn, ok := e.(*net.Conn)
		if !ok {
			panic(fmt.Sprintf("Should have net.Conn, got %T", e))
		}
		return (*conn).Close()
	})

	em.On("connection.broadcast", func(e interface{}) error {
		data, ok := e.(map[string]interface{})
		if !ok {
			panic(fmt.Sprintf("Should have map[string]interface{}, got %T", e))
		}
		var errs []error
		for _, conn := range conns {
			if err := em.Emit("connection.send",
				utils.Message{From: conn.raw, Data: data}); err != nil {
				errs = append(errs, err)
			}
		}
		return utils.ErrorFrom(errs)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		raw, _, _, err := ws.UpgradeHTTP(r, w, nil)
		if err != nil {
			log.Printf("Couldn't upgrade HTTP to WebSocket: %s", err)
			return
		}
		go func() {
			defer raw.Close()
			// read from socket and emit an event as soon as there is a message
			conn := connection{
				raw:    &raw,
				writer: wsutil.NewWriter(raw, ws.StateServerSide, ws.OpText),
				reader: wsutil.NewReader(raw, ws.StateServerSide),
			}
			conn.encoder = json.NewEncoder(conn.writer)
			conn.decoder = json.NewDecoder(conn.reader)
			conns = append(conns, conn)
			em.Emit("connection.new", conn)
			// read from the connection forever and close when there is an error
			messagecount := 1
			for {
				header, err := conn.reader.NextFrame()
				if conn.closed {
					// we have manually closed the connection, just ignore the rest
					return
				}
				if err != nil {
					em.Emit("connection.error", utils.Error{Err: err, Data: conn})
					em.Emit("connection.closed", conn)
					return
				}
				if header.OpCode == ws.OpClose {
					em.Emit("connection.closed", conn)
					return
				}

				var msg utils.Message
				msg.From = conn.raw
				msg.Data = make(map[string]interface{})
				msg.Count = messagecount

				if err := conn.decoder.Decode(&msg.Data); err != nil {
					log.Printf("Error occured while parsing WebSocket: %s", err)
					return
				}

				if err := em.Emit("connection.msg", msg); err != nil {
					log.Printf("Error on 'connection.msg':\n%s", err)
				}

				messagecount++
			}
		}()
	})
	app.Start()
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
