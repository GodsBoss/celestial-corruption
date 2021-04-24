package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type game struct{}

func New() dominit.Game {
	return &game{}
}

func (g *game) TicksPerSecond() int {
	return 40
}

func (g *game) Tick(ms int) {}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) SetOutput(ctx2d *dom.Context2D) {}

func (g *game) Render() {}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
	return 640, 400, 2, 2
}
