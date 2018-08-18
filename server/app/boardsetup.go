package app

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

// makes sure that there is a right amount of boats, of the right size, etc...
func validBoats(boats []Boat) error {
	if len(boats) != 5 {
		return fmt.Errorf("Invalid number of boats. Should have 5, got %d", len(boats))
	}
	// ensure we have the right amount of boats of the right size
	var sizes = make(map[int]int)
	var occupied []Point
	for _, boat := range boats {
		if boat.Size <= 1 {
			return fmt.Errorf("Invalid boat size: should be >1, got %d", boat.Size)
		}
		if boat.Pos.X < 0 || boat.Pos.X >= 10 || boat.Pos.Y < 0 || boat.Pos.Y >= 10 {
			return fmt.Errorf("Invalid boat origin point: should be 0 < x or y < 10, got %s",
				boat.Pos)
		}
		if (boat.Rot == 0 && boat.Pos.X+boat.Size >= 10) || (boat.Rot == 1 && boat.Pos.Y+boat.Size >= 10) {
			return fmt.Errorf("Invalid boat position: some of it is outside the map")
		}
		if boat.Rot != 0 && boat.Rot != 1 {
			return fmt.Errorf("Invlaid boat rotation: should be 0 or 1, got %d", boat.Rot)
		}
		fmt.Println(boat, boat.Pos.X >= 10)
		sizes[boat.Size]++
		var pt Point
		for i := 0; i < boat.Size; i++ {
			if boat.Rot == 0 {
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

func handleBoardSetup(players []Player, data map[string]interface{}) error {
	raw, ok := data["boats"]
	if !ok {
		return fmt.Errorf("Not boat setup in data: %s", data)
	}
	var boats []Boat
	mapstructure.Decode(&boats, raw)
	if err := validBoats(boats); err != nil {
		return err
	}
	return nil
}
