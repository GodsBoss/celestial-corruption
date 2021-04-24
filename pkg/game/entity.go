package game

import (
  "math"
)

type entity struct {
  x float64
  y float64
  w float64
  h float64
}

func (e entity) Left() float64 {
  return e.x
}

func (e entity) Right() float64 {
  return e.x + e.w
}

func (e entity) Top() float64 {
  return e.y
}

func (e entity) Bottom() float64 {
  return e.y + e.h
}

func (e entity) Center() (x, y float64) {
  return e.x + (e.w / 2), e.y + (e.h / 2)
}

func entityCollision(e1, e2 entity) (entity, bool) {
  left, right := e1, e2
  if left.x > right.x {
    left, right = right, left
  }

  if right.Left() > left.Right() {
    return entity{}, false
  }

  top, bottom := e1, e2
  if top.y > bottom.y {
    top, bottom = bottom, top
  }

  if bottom.Top() > top.Bottom() {
    return entity{}, false
  }

  return entity{
    x: right.x,
    y: bottom.y,
    w: math.Min(right.Right(), left.Right()),
    h: math.Min(bottom.Bottom(), top.Bottom()),
  }, true
}
