package app

import (
	"testing"
)

// usage: newBoat("3234") -> Boat{3, Point{2, 2}, 4}
func newBoat(size, x, y, rot int) Boat {
	return Boat{size, Point{x, y}, rot}
}

func TestValidBoats(t *testing.T) {
	if err := validBoats([]Boat{}); err == nil {
		t.Errorf("Should have returned an error: length was 0")
	}
	if err := validBoats([]Boat{
		newBoat(2, 0, 0, 0), newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0),
		newBoat(4, 0, 3, 0),
	}); err == nil {
		t.Errorf("Should have returned an error: invalid sizes of boats")
	}
	if err := validBoats([]Boat{
		newBoat(3, 0, 0, 0), newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0),
		newBoat(5, 0, 3, 0),
	}); err == nil {
		t.Errorf("Should have returned an error: invalid sizes of boats")
	}
	if err := validBoats([]Boat{
		newBoat(2, 10, 0, 0), // this boat is outside
		newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0), newBoat(5, 0, 3, 0),
	}); err == nil {
		t.Errorf("Should have returned an error: origin of boats are out side")
	}
	if err := validBoats([]Boat{
		newBoat(2, 10, 10, 0), // this boat is outside
		newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0), newBoat(5, 0, 3, 0),
	}); err == nil {
		t.Errorf("Should have returned an error: origin of boats are out side")
	}
	if err := validBoats([]Boat{
		newBoat(2, 9, 0, 0), // this boat is outside
		newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0), newBoat(5, 0, 3, 0),
	}); err == nil {
		t.Errorf("Should have returned an error: parts of boats are out side")
	}
	if err := validBoats([]Boat{
		newBoat(2, 9, 0, 0), newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0), newBoat(4, 0, 2, 0),
		newBoat(5, 0, 5, 1), // this boat is outside
	}); err == nil {
		t.Errorf("Should have returned an error: parts of boats are out side")
	}
	if err := validBoats([]Boat{
		newBoat(2, 9, 0, 0), newBoat(3, 2, 0, 0), newBoat(3, 0, 1, 0),
		newBoat(4, 0, 2, 0), newBoat(5, 0, 5, 1), // this boat is outside
	}); err == nil {
		t.Errorf("Should have returned an error: parts of boats are out side")
	}
}
