package app

import (
	"fmt"
	"log"

	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
	"github.com/mitchellh/mapstructure"
)

var players []Player

type firstmessage struct {
	Kind string
	Name string
}

func Start() {
	em.On("connection.msg", func(e interface{}) error {
		msg, ok := e.(utils.Message)
		if !ok {
			panic(fmt.Sprintf("Should have utils.Message, got %T", e))
		}
		if msg.Count == 1 {
			// this the player's first message
			if len(players) == 2 {
				// we have enough players.
				em.Emit("connection.send", utils.NewMessage(msg.From, "kind", "state change",
					"state", "enough players"))
				if errs := em.Emit("connection.close", msg.From); len(errs) != 0 {
					log.Print("Couldn't close connection with extra client")
				}
			}
			var data firstmessage
			mapstructure.Decode(msg.Data, &data)
			players = append(players, Player{Name: data.Name})
			if data.Kind != "request" {
				log.Printf("Got invalid first message. Should have 'request', got '%s'", data.Kind)
				em.Emit("connection.close", msg.From)
			}
			if len(players) == 2 {
				em.Emit("connection.broadcast", utils.MakeMap("kind", "state change",
					"state", "setup"))
			}
		}
		return nil
	})
}
