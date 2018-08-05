package app

type Point struct {
	X, Y int
}

type Boat struct {
	Size     int
	Pos      Point
	Rotation int // 0, 1, 2 or 3
}

type Board struct {
	Hits  []Point
	Boats []Boat
}

type Player struct {
	Name  string `json:"name"`
	Board Board
}
