package app

import (
	"fmt"
	"log"

	"github.com/math2001/boatsinker/server/em"
	"github.com/math2001/boatsinker/server/utils"
	"github.com/mitchellh/mapstructure"
)

// handle new connections (first message)

type firstmessage struct {
	Kind string
	Name string
}

func handleFirstMessage(players []Player, msg utils.Message) ([]Player, error) {
	if len(players) == 2 {
		// we have enough players.
		if err := em.Emit("connection.send", utils.NewMessage(msg.From,
			"kind", "state change", "state", "enough players")); err != nil {
			log.Printf("Couldn't send refusal to extra client: %s", err)
		}
		if err := em.Emit("connection.close", msg.From); err != nil {
			log.Printf("Couldn't close connection with extra client: %s", err)
		}
	}
	var data firstmessage
	mapstructure.Decode(msg.Data, &data)
	players = append(players, Player{Name: data.Name})
	if data.Kind != "request" {
		log.Printf("Got invalid first message. Should have 'request', got '%s'", data.Kind)
		if err := em.Emit("connection.close", msg.From); err != nil {
			fmt.Printf("Couldn't close connection with invalid first message client:\n%s",
				err)
		}
	}
	if len(players) == 2 {
		if err := em.Emit("connection.broadcast",
			utils.MakeMap("kind", "state change", "state", "setup",
				"boat_sizes", boat_sizes, "size", mapsize)); err != nil {
			fmt.Printf("Error while broadcasting 'setup' message:\n%s", err)
		}
	}
	return players, nil
}
