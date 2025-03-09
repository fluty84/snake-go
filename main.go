package main

import (
	"time"

	"snake/internal/application"
	"snake/internal/core"
	"snake/internal/infrastructure"
)

func main() {
	gameService := application.NewGameService(40, 20)
	adapter := infrastructure.NewGUIAdapter(gameService)

	input := make(chan core.Direction)
	go func() {
		for {
			input <- adapter.HandleInput()
		}
	}()

	go gameLoop(gameService, adapter, input)
	adapter.Run()
}

func gameLoop(service *application.GameService, adapter *infrastructure.GUIAdapter, input chan core.Direction) {
	for !service.GetState().GameOver {
		select {
		case dir := <-input:
			if dir == -2 {
				service.Reset()
				continue
			}
			service.ChangeDirection(dir)
		default:
		}

		service.Move()
		adapter.Render(service.GetState())
		time.Sleep(100 * time.Millisecond)
	}

	// Handle game over and restart
	for {
		dir := <-input
		if dir == -2 {
			service.Reset()
			gameLoop(service, adapter, input)
			return
		}
	}
}