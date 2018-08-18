package app

import (
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/math2001/boatsinker/server/utils"
	"github.com/mitchellh/mapstructure"
)

// makes sure that there is a right amount of boats, of the right size, etc...
func validBoats(boats []Boat) error {
	if len(boats) != len(boatsizes) {
		return fmt.Errorf("Invalid number of boats. Should have %d, got %d", len(boatsizes),
			len(boats))
	}
	// ensure we have the right amount of boats of the right size
	var sizes = make(map[int]int)
	var occupied []Point
	for _, boat := range boats {
		if boat.Size <= 1 {
			return fmt.Errorf("Invalid boat size: should be >1, got %d", boat.Size)
		}
		if boat.Pos.X < 0 || boat.Pos.X >= width || boat.Pos.Y < 0 || boat.Pos.Y >= height {
			return fmt.Errorf("Invalid boat origin point: should be 0 < x or y < 10, got %s",
				boat.Pos)
		}
		if (boat.Rotation == 0 && boat.Pos.X+boat.Size >= width) || (boat.Rotation == 1 && boat.Pos.Y+boat.Size >= height) {
			return fmt.Errorf("Invalid boat position: some of it is outside the map")
		}
		if boat.Rotation != 0 && boat.Rotation != 1 {
			return fmt.Errorf("Invlaid boat rotation: should be 0 or 1, got %d",
				boat.Rotation)
		}
		sizes[boat.Size]++
		var pt Point
		for i := 0; i < boat.Size; i++ {
			if boat.Rotation == 0 {
				pt = Point{boat.Pos.X + i, boat.Pos.Y}
			} else {
				pt = Point{boat.Pos.X, boat.Pos.Y + i}
			}
			for _, t := range occupied {
				if t == pt {
					return fmt.Errorf("Invalid boats position: they collide on the case %s", pt)
				}
			}
			occupied = append(occupied, pt)
		}
	}
	if !reflect.DeepEqual(sizes, boatsizes) {
		return fmt.Errorf("Invalid boat sizes: should have %v, got %v", sizes, boatsizes)
	}
	return nil
}

type boardsetup struct {
	Kind  string
	Boats []Boat
}

func handleBoardSetup(players map[*net.Conn]*Player, msg utils.Message) error {
	var b boardsetup
	if err := mapstructure.Decode(msg.Data, &b); err != nil {
		log.Fatalf("Couldn't convert msg to boardsetup: %s", err)
	}
	if err := validBoats(b.Boats); err != nil {
		return err
	}
	players[msg.From].Board = Board{
		Boats: b.Boats,
	}
	return nil
}
