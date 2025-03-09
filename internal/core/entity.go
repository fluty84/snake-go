package core

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Point struct {
	X, Y int
}

type Snake struct {
	Body      []Point
	Direction Direction
}

type GameState struct {
	Snake    Snake
	Food     Point
	Score    int
	GameOver bool
	Width    int
	Height   int
}

type Game interface {
	Move()
	ChangeDirection(Direction)
	GetState() GameState
	Reset()
}