package game

import (
	"github.com/GodsBoss/gggg/pkg/dom"
	"github.com/GodsBoss/gggg/pkg/dominit"
	"github.com/GodsBoss/gggg/pkg/interaction"
)

type game struct{
	img *dom.Image
  output *dom.Context2D
	scale *sharableInt

	states map[string]state
	currentStateID string
}

func New(img *dom.Image) dominit.Game {
	g := &game{
		img: img,
		scale: &sharableInt{},
		states: map[string]state{
			"title": &title{},
			"playing": &playing{},
		},
	}
	g.scale.Set(1)
	g.nextState("title")
	return g
}

// nextState switches the state. Does nothing if id is an empty string.
func (g *game) nextState(id string) {
	if id == "" {
		return
	}
	g.currentStateID = id
	g.states[id].init()
}

func (g *game) currentState() state {
	return g.states[g.currentStateID]
}

func (g *game) TicksPerSecond() int {
	return tps
}

func (g *game) Tick(ms int) {
	g.nextState(g.currentState().tick(ms))
}

func (g *game) ReceiveKeyEvent(event interaction.KeyEvent) {
	g.nextState(g.currentState().receiveKeyEvent(event))
}

func (g *game) ReceiveMouseEvent(event interaction.MouseEvent) {}

func (g *game) SetOutput(ctx2d *dom.Context2D) {
  g.output = ctx2d
}

func (g *game) Render() {
	scale := g.scale.Get()
	g.output.ClearRect(0, 0, gfxWidth*scale, gfxHeight*scale)
	fillStyle, _ := dom.NewColorCanvasFillStyle("#111111")
	g.output.SetFillStyle(fillStyle)
	g.output.FillRect(0, 0, gfxWidth*scale, gfxHeight*scale)
}

func (g *game) Scale(availableWidth, availableHeight int) (realWidth, realHeight int, scaleX, scaleY float64) {
  wf, hf := (availableWidth - 20) / gfxWidth, (availableHeight - 20) / gfxHeight
  f := wf
  if hf < f {
    f = hf
  }
  if f < 1 {
    f = 1
  }
	g.scale.Set(f)
	return f*gfxWidth, f*gfxHeight, float64(f), float64(f)
}

const (
  gfxWidth int = 320
  gfxHeight int = 200

	tps = 40
)
