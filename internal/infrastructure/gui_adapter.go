package infrastructure

import (
	"fmt"
	"image/color"
	"snake/internal/core"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app" // v2.5.4
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type GUIAdapter struct {
	gameService core.Game
	app         fyne.App
	window      fyne.Window
	content     *fyne.Container
	cellSize    float32
	inputChan   chan core.Direction
}

func NewGUIAdapter(game core.Game) *GUIAdapter {
	a := app.New()
	w := a.NewWindow("Snake Game")
	cellSize := float32(20)

	adapter := &GUIAdapter{
		gameService: game,
		app:         a,
		window:      w,
		cellSize:    cellSize,
		inputChan:   make(chan core.Direction, 1),
	}

	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyUp:
			adapter.inputChan <- core.Up
		case fyne.KeyDown:
			adapter.inputChan <- core.Down
		case fyne.KeyLeft:
			adapter.inputChan <- core.Left
		case fyne.KeyRight:
			adapter.inputChan <- core.Right
		case fyne.KeyEscape:
			adapter.inputChan <- core.Direction(-1)
		}
	})

	w.SetContent(adapter.createUI())
	w.Resize(fyne.NewSize(
		float32(game.GetState().Width)*cellSize,
		float32(game.GetState().Height)*cellSize+50,
	))

	return adapter
}

func (g *GUIAdapter) createUI() *fyne.Container {
	g.content = container.NewWithoutLayout()
	return g.content
}

func (g *GUIAdapter) Render(state core.GameState) {
	g.content.RemoveAll()

	// Draw score
	scoreText := canvas.NewText("Score: "+ fmt.Sprint(state.Score), color.White)
	scoreText.Move(fyne.NewPos(10, float32(state.Height)*g.cellSize+10))
	g.content.Add(scoreText)

	// Draw game elements
	for _, p := range state.Snake.Body {
		rect := canvas.NewRectangle(color.RGBA{R: 0, G: 255, B: 0, A: 255})
		rect.Resize(fyne.NewSize(g.cellSize, g.cellSize))
		rect.Move(fyne.NewPos(
			float32(p.X)*g.cellSize,
			float32(p.Y)*g.cellSize,
		))
		g.content.Add(rect)
	}

	// Draw food
	foodRect := canvas.NewRectangle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	foodRect.Resize(fyne.NewSize(g.cellSize, g.cellSize))
	foodRect.Move(fyne.NewPos(
		float32(state.Food.X)*g.cellSize,
		float32(state.Food.Y)*g.cellSize,
	))
	g.content.Add(foodRect)

	if state.GameOver {
		overText := canvas.NewText("Game Over! Press ESC to quit", color.RGBA{R: 255, A: 255})
		overText.Alignment = fyne.TextAlignCenter
		overText.Resize(fyne.NewSize(300, 30))
		overText.Move(fyne.NewPos(
			float32(state.Width)*g.cellSize/2-150,
			float32(state.Height)*g.cellSize/2-15,
		))
		g.content.Add(overText)
	}

	g.content.Refresh()
}

func (g *GUIAdapter) Run() {
	g.window.ShowAndRun()
}

func (g *GUIAdapter) HandleInput() core.Direction {
	return <-g.inputChan
}