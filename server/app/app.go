package app

import (
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
)

// list of every players
var players = make(map[*net.Conn]*Player)

// the player who has to play. By default, it's nil
var turn *net.Conn

// the size of the board
const width = 10
const height = 10

var boatsizes = map[int]int{
	// boat size: count
	5: 1,
	// 4: 1,
	// 3: 2,
	// 2: 1,
}

func playersReady() bool {
	for _, p := range players {
		// if p.Board is empty, this means we are waiting for the player to
		// send his board set up
		if reflect.DeepEqual((Board{}), p.Board) {
			return false
		}
	}
	return true
}

// Start listen to connection events. They're the own who are going to actually
// start up the game
func Start() {

	em.On("connection.closed", func(e interface{}) error {
		// when a connection is closed, we close the game, and shutdown for now
		log.Fatal("A player left. Shutdown")
		return nil
	})

	em.On("connection.msg", func(e interface{}) error {
		msg, ok := e.(utils.Message)
		if !ok {
			panic(fmt.Sprintf("Should have utils.Message, got %T", e))
		}
		if msg.Count == 1 {
			return handleFirstMessage(players, msg)
		}
		// if this isn't the first message, and we haven't got 2 players,
		// there's a problem. This shouldn't happen
		if len(players) != 2 {
			log.Fatalf("Got message second %s, but haven't got a second player. Shutdown",
				msg)
		}

		kind, ok := msg.Data["kind"]
		if !ok {
			log.Print("No 'kind' field in message.")
			if err := em.Emit("connection.close", msg.From); err != nil {
				log.Print("Couldn't close connection (message didn't have 'kind' field).")
			}
			return nil
		}
		if kind == "board setup" {
			if err := handleBoardSetup(players, msg); err != nil {
				log.Fatalf("Invalid 'board ready' message: %s", err)
			}
			if playersReady() {
				// tell one to play and the other to wait
				count := 0
				for conn := range players {
					if count == 0 {
						em.Emit("connection.send", utils.NewMessage(conn, "kind", "hit"))
					} else {
						em.Emit("connection.send", utils.NewMessage(conn, "kind", "wait"))
					}
					count++
				}
			}
			return nil
		}
		log.Fatalf("Unknown message kind: '%s' in message '%s'", kind, msg)
		return nil
	})
}
