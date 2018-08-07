package app

import (
	"fmt"
	"log"

	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
)

var players []Player

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
		return nil
	})
}
