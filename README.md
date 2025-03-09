# Snake Game

A classic snake game implementation using Go and Fyne GUI toolkit with hexagonal architecture.

## Project Structure

```
├── Makefile            # Build automation
├── build_and_run.sh    # Build script
├── go.mod              # Go dependencies
├── internal/
│   ├── core/           # Domain entities & interfaces
│   ├── application/    # Business logic (use cases)
│   └── infrastructure/ # Adapters (GUI, terminal)
└── main.go            # Application composition root
```

## Technology Stack

- Go 1.21
- Fyne GUI Toolkit v2.5.4
- Hexagonal Architecture
- Termbox-go (for keyboard input)

## Installation

1. Install Go 1.21+
2. Install Fyne requirements:
```bash
$ go get fyne.io/fyne/v2
$ go install fyne.io/fyne/cmd/fyne@latest
```

## How to Build & Run

```bash
# Build macOS application bundle
$ make build

# Run the game
$ make run

# Alternative: Use build script
$ chmod +x build_and_run.sh
$ ./build_and_run.sh
```

## Game Controls

- Arrow Keys: Change snake direction
- R Key: Restart game
- ESC: Quit game

## Architecture

Following hexagonal architecture principles:
1. **Core**: Game state, entities, and interfaces
2. **Application**: Game rules and use cases
3. **Infrastructure**: GUI implementation (Fyne adapter)

## Dependencies

- Fyne.io/v2 - Cross-platform GUI
- nsf/termbox-go - Terminal input handling
- Go modules for dependency management