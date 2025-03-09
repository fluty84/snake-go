package infrastructure

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"snake/internal/core"
)

type TerminalAdapter struct {
	gameService core.Game
}

func NewTerminalAdapter(game core.Game) *TerminalAdapter {
	return &TerminalAdapter{gameService: game}
}

func (t *TerminalAdapter) Render(state core.GameState) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Draw score
	scoreStr := fmt.Sprintf("Score: %d", state.Score)
	for i, ch := range scoreStr {
		termbox.SetCell(i, state.Height, ch, termbox.ColorWhite, termbox.ColorDefault)
	}

	// Draw snake
	for _, p := range state.Snake.Body {
		termbox.SetCell(p.X, p.Y, '■', termbox.ColorGreen, termbox.ColorDefault)
	}

	// Draw food
	termbox.SetCell(state.Food.X, state.Food.Y, '●', termbox.ColorRed, termbox.ColorDefault)

	if state.GameOver {
		gameOverMsg := "Game Over! Press ESC to quit"
		startX := state.Width/2 - len(gameOverMsg)/2
		for i, ch := range gameOverMsg {
			termbox.SetCell(startX+i, state.Height/2, ch, termbox.ColorRed, termbox.ColorDefault)
		}
	}

	termbox.Flush()
}

func (t *TerminalAdapter) HandleInput() core.Direction {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			return core.Up
		case termbox.KeyArrowDown:
			return core.Down
		case termbox.KeyArrowLeft:
			return core.Left
		case termbox.KeyArrowRight:
			return core.Right
		case termbox.KeyEsc:
			return core.Direction(-1)
		}
	}
	return core.Direction(-2)
}