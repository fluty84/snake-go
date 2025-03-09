package main

import (
	"time"

	"github.com/nsf/termbox-go"
	"snake/internal/application"
	"snake/internal/core"
	"snake/internal/infrastructure"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	gameService := application.NewGameService(40, 20)
	adapter := infrastructure.NewTerminalAdapter(gameService)

	input := make(chan core.Direction)
	go func() {
		for {
			input <- adapter.HandleInput()
		}
	}()

	gameLoop(gameService, adapter, input)
}

func gameLoop(service *application.GameService, adapter *infrastructure.TerminalAdapter, input chan core.Direction) {
	for !service.GetState().GameOver {
		select {
		case dir := <-input:
			if dir == -1 {
				return
			}
			service.ChangeDirection(dir)
		default:
		}

		service.Move()
		adapter.Render(service.GetState())
		time.Sleep(100 * time.Millisecond)
	}

	// Handle game over
	adapter.Render(service.GetState())
	waitForEsc()
}

func waitForEsc() {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}