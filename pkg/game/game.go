package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type game struct{
  output *dom.Context2D
	scale int
}

func New() dominit.Game {
	return &game{
		scale: 1,
	}
}

func (g *game) TicksPerSecond() int {
	return 40
}

func (g *game) Tick(ms int) {}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) SetOutput(ctx2d *dom.Context2D) {
  g.output = ctx2d
}

func (g *game) Render() {}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
  wf, hf := (availableWidth - 20) / gfxWidth, (availableHeight - 20) / gfxHeight
  f := wf
  if hf < f {
    f = hf
  }
  if f < 1 {
    f = 1
  }
	g.scale = 1
	return f*gfxWidth, f*gfxHeight, float64(f), float64(f)
}

const (
  gfxWidth int = 320
  gfxHeight int = 200
)
