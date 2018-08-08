package app

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
)

var boat_sizes_count = map[int]int{
	// boat size: count
	5: 1,
	4: 1,
	3: 2,
	2: 1,
}

// makes sure that there is a right amount of boats, of the right size, etc...
func validBoats(boats []Boat) error {
	if len(boats) != 5 {
		return fmt.Errorf("Invalid number of boats. Should have 5, got %d", len(boats))
	}
	// ensure we have the right amount of boats of the right size
	var sizes = make(map[int]int)
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
		sizes[boat.Size] += 1
	}
	if reflect.DeepEqual(sizes, boat_sizes_count) {
		return fmt.Errorf("Invalid boat sizes: should have %v, got %v", sizes, boat_sizes_count)
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
