package cell

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Cell struct {
	Alive bool
}

func (c *Cell) Update() {}

func (c Cell) Draw(dst *ebiten.Image, x, y, size int) {
	if c.Alive {
		vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.White, false)
	} else {
		vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.Black, false)
	}
}
