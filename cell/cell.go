package cell

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Cell struct {
	Alive bool
	Dying bool
	Born  bool
}

func (c *Cell) Update() {}

func (c Cell) Draw(dst *ebiten.Image, x, y, size int, colored bool) {
	if colored {

		if c.Dying {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.RGBA{255, 0, 255, 255}, false)
		} else if c.Born {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.RGBA{0, 0, 255, 255}, false)
		} else if c.Alive {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.RGBA{0, 255, 255, 255}, false)
		} else {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.Black, false)
		}
	} else {
		if c.Alive {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.White, false)
		} else {
			vector.DrawFilledRect(dst, float32(x*size), float32(y*size), float32(size), float32(size), color.Black, false)
		}
	}

}
