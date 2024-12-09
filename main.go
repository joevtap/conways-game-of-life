package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/joevtap/conways-game-of-life/grid"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth    = 500
	screenHeight   = 500
	updateInterval = 50 * time.Millisecond
	cellSize       = 8
)

type Game struct {
	grid             grid.Grid
	running          bool
	lastUpdated      time.Time
	mousePressed     bool
	spacePressed     bool
	backspacePressed bool
	rPressed         bool
	cPressed         bool
	colored          bool
}

func main() {
	game := &Game{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Conway's Game of Life")

	game.running = false
	game.grid = grid.New(screenWidth, screenHeight, cellSize)
	game.lastUpdated = time.Now()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		if !g.spacePressed {
			g.running = !g.running
			g.spacePressed = true
		}
	} else {
		g.spacePressed = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		if !g.rPressed && !g.running {
			g.grid.Clear()
			g.grid.Randomize()
			g.running = true
			g.rPressed = true
		}
	} else {
		g.rPressed = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		if !g.cPressed {
			g.colored = !g.colored
			g.cPressed = true
		}
	} else {
		g.cPressed = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyBackspace) {
		if !g.backspacePressed {
			g.grid.Clear()
			g.running = false
			g.backspacePressed = true
		}
	} else {
		g.backspacePressed = false
	}

	if !g.running {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			if !g.mousePressed {
				x, y := ebiten.CursorPosition()
				g.grid.ToggleCell(x, y)
				g.mousePressed = true
			}
		} else {
			g.mousePressed = false
		}
	}

	if g.running {
		now := time.Now()
		if now.Sub(g.lastUpdated) >= updateInterval {
			g.grid.Update()
			g.lastUpdated = now
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.grid.Draw(screen, g.colored)

	if g.running {
		text.Draw(screen, "Press SPACE to stop", basicfont.Face7x13, 10, 20, color.White)
		text.Draw(screen, "Simulation running...", basicfont.Face7x13, 10, screenHeight-20, color.White)
	} else {
		text.Draw(screen, "Press SPACE to start", basicfont.Face7x13, 10, 20, color.White)
		text.Draw(screen, "Click to toggle cells", basicfont.Face7x13, 10, 60, color.White)
		text.Draw(screen, "Press \"r\" to randomly generate cells", basicfont.Face7x13, 10, 80, color.White)
		text.Draw(screen, "Press \"c\" to toggle colored cells", basicfont.Face7x13, 10, 100, color.White)
	}

	text.Draw(screen, "Press BACKSPACE to clear", basicfont.Face7x13, 10, 40, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
