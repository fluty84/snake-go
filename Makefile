.PHONY: clean build run

clean:
	rm -rf Snake.app

build:
	fyne package -os darwin -name Snake -icon icon.png -appID com.example.snake

run: build
	open Snake.app