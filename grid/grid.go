package grid

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/joevtap/conways-game-of-life/cell"
)

type Grid struct {
	Width    int
	Height   int
	CellSize int
	Cells    [][]cell.Cell
}

func New(screenWidth, screenHeight, cellSize int) Grid {

	width := screenWidth / cellSize
	height := screenHeight / cellSize

	cells := make([][]cell.Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]cell.Cell, width)
	}

	return Grid{
		Width:    width,
		CellSize: cellSize,
		Height:   height,
		Cells:    cells,
	}
}

func (g Grid) Draw(dst *ebiten.Image, colored, showGrid bool) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := g.Cells[y][x]
			cell.Draw(dst, x, y, g.CellSize, colored)
		}
	}

	for x := 0; x < g.Width; x++ {
		if showGrid {
			vector.StrokeLine(dst, float32(x*g.CellSize), 0, float32(x*g.CellSize), float32(g.Height*g.CellSize), 1, color.RGBA{20, 20, 20, 255}, false)
		} else {
			vector.StrokeLine(dst, float32(x*g.CellSize), 0, float32(x*g.CellSize), float32(g.Height*g.CellSize), 1, color.RGBA{0, 0, 0, 255}, false)
		}
	}

	for y := 0; y < g.Height; y++ {
		if showGrid {
			vector.StrokeLine(dst, 0, float32(y*g.CellSize), float32(g.Width*g.CellSize), float32(y*g.CellSize), 1, color.RGBA{20, 20, 20, 255}, false)
		} else {
			vector.StrokeLine(dst, 0, float32(y*g.CellSize), float32(g.Width*g.CellSize), float32(y*g.CellSize), 1, color.RGBA{0, 0, 0, 255}, false)

		}
	}

}

func (g *Grid) Update() {
	newCells := make([][]cell.Cell, g.Height)
	for y := 0; y < g.Height; y++ {
		newCells[y] = make([]cell.Cell, g.Width)
		for x := 0; x < g.Width; x++ {
			newCells[y][x] = g.Cells[y][x]
			aliveNeighbors := g.countAliveNeighbors(x, y)
			if g.Cells[y][x].Alive {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					newCells[y][x].Alive = false
					newCells[y][x].Dying = true
					newCells[y][x].Born = false
				} else {
					newCells[y][x].Dying = false
					newCells[y][x].Born = false
				}
			} else {
				if aliveNeighbors == 3 {
					newCells[y][x].Alive = true
					newCells[y][x].Born = true
					newCells[y][x].Dying = false
				} else {
					newCells[y][x].Dying = false
					newCells[y][x].Born = false
				}
			}
		}
	}
	g.Cells = newCells
}

func (g *Grid) countAliveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			ny, nx := y+dy, x+dx
			if nx >= 0 && nx < g.Width && ny >= 0 && ny < g.Height {
				if g.Cells[ny][nx].Alive {
					count++
				}
			}
		}
	}
	return count
}

func (g *Grid) ToggleCell(x, y int) {
	g.Cells[y/g.CellSize][x/g.CellSize].Alive = !g.Cells[y/g.CellSize][x/g.CellSize].Alive
}

func (g *Grid) Clear() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			g.Cells[y][x].Alive = false
			g.Cells[y][x].Dying = false
			g.Cells[y][x].Born = false
		}
	}
}

func (g *Grid) Randomize() {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			if rand.Intn(2) == 0 {
				g.Cells[y][x].Alive = true
			} else {
				g.Cells[y][x].Alive = false
			}
		}
	}
}
