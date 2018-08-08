package app

import (
	"fmt"
	"log"

	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
)

var players []Player

// the size of the board (10 by 10)
const mapsize = 10

var boat_sizes = map[int]int{
	// boat size: count
	5: 1,
	4: 1,
	3: 2,
	2: 1,
}

func Start() {
	em.On("connection.closed", func(e interface{}) error {
		// conn, ok := e.(net.Conn)
		// if !ok {
		// 	panic("Should have net.Conn")
		// }
		// when a connection is closed, we close the game, and shutdown
		// ... for now
		log.Fatal("A player left. Shutdown")
		return nil
	})
	em.On("connection.msg", func(e interface{}) error {
		msg, ok := e.(utils.Message)
		if !ok {
			panic(fmt.Sprintf("Should have utils.Message, got %T", e))
		}
		if msg.Count == 1 {
			var err error
			players, err = handleFirstMessage(players, msg)
			return err
		}
		// if this isn't the first message, and we haven't got 2 players, there's a problem
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
		if kind == "board ready" {
			if err := handleBoardSetup(players, msg.Data); err != nil {
				log.Fatalf("Invalid 'board ready' message: %s", err)
			}
			return nil
		}
		log.Fatalf("Unknown message kind: '%s' in message '%s'", kind, msg)
		return nil
	})
}
