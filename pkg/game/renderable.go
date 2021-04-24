package game

import (
  "github.com/GodsBoss/gggg/pkg/dom"
)

type renderable interface {
  Render(output *dom.Context2D)
}

type nopRenderable struct{}

func (r nopRenderable) Render(output *dom.Context2D) {}
