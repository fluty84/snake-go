package application

import (
	"math/rand"
	"time"

	"snake/internal/core"
)

type GameService struct {
	state     core.GameState
	rand      *rand.Rand
}

func NewGameService(width, height int) *GameService {
	return &GameService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		state: core.GameState{
			Width:    width,
			Height:   height,
			Snake: core.Snake{
				Body: []core.Point{{X: width/2, Y: height/2}},
				Direction: core.Right,
			},
		},
	}
}

func (g *GameService) Move() {
	head := g.state.Snake.Body[0]
	newHead := core.Point{
		X: head.X + directionToVector(g.state.Snake.Direction).X,
		Y: head.Y + directionToVector(g.state.Snake.Direction).Y,
	}

	g.state.Snake.Body = append([]core.Point{newHead}, g.state.Snake.Body...)

	if newHead == g.state.Food {
		g.state.Score++
		g.spawnFood()
	} else {
		g.state.Snake.Body = g.state.Snake.Body[:len(g.state.Snake.Body)-1]
	}

	g.checkCollisions()
}

func (g *GameService) ChangeDirection(d core.Direction) {
	current := g.state.Snake.Direction
	if (current == core.Up && d != core.Down) || 
	   (current == core.Down && d != core.Up) ||
	   (current == core.Left && d != core.Right) ||
	   (current == core.Right && d != core.Left) {
		g.state.Snake.Direction = d
	}
}

func (g *GameService) GetState() core.GameState {
	return g.state
}

func (g *GameService) Reset() {
	g.state = core.GameState{
		Width:    g.state.Width,
		Height:   g.state.Height,
		Snake: core.Snake{
			Body: []core.Point{{X: g.state.Width/2, Y: g.state.Height/2}},
			Direction: core.Right,
		},
	}
	g.spawnFood()
}

func (g *GameService) checkCollisions() {
	head := g.state.Snake.Body[0]
	g.state.GameOver = head.X < 0 || head.X >= g.state.Width ||
		head.Y < 0 || head.Y >= g.state.Height

	for _, p := range g.state.Snake.Body[1:] {
		if head == p {
			g.state.GameOver = true
			break
		}
	}
}

func (g *GameService) spawnFood() {
	for {
		g.state.Food = core.Point{
			X: g.rand.Intn(g.state.Width),
			Y: g.rand.Intn(g.state.Height),
		}

		collision := false
		for _, p := range g.state.Snake.Body {
			if g.state.Food == p {
				collision = true
				break
			}
		}

		if !collision {
			break
		}
	}
}

func directionToVector(d core.Direction) core.Point {
	switch d {
	case core.Up:
		return core.Point{Y: -1}
	case core.Down:
		return core.Point{Y: 1}
	case core.Left:
		return core.Point{X: -1}
	default:
		return core.Point{X: 1}
	}
}