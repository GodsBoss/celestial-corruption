package game

import (
  "github.com/GodsBoss/gggg/pkg/dom"
)

type spriteFactory struct {
  source *dom.Image
  scale *sharableInt
  infos map[string]spriteInfo
}

func (sf *spriteFactory) create(id string, x, y, frame int) renderable {
  info, ok := sf.infos[id]
  if !ok {
    return nopRenderable{}
  }
  return sprite{
    source: sf.source,
    scale: sf.scale,

    x: x,
    y: y,
    frame: frame,

    sx: info.x,
    sy: info.y,
    sw: info.w,
    sh: info.h,
  }
}

type sprite struct {
  source *dom.Image
  scale *sharableInt

  x int
  y int
  frame int

  sx int
  sy int
  sw int
  sh int
}

func (s sprite) Render(output *dom.Context2D) {
  scale := s.scale.Get()
  x := s.x * scale
  y := s.y * scale
  w := s.sw * scale
  h := s.sh * scale

  sx := s.sx + s.frame * s.sw

  output.DrawImage(s.source, sx, s.sy, s.sw, s.sh, x, y, w, h)
}

type spriteInfo struct {
  x int
  y int
  w int
  h int
}
