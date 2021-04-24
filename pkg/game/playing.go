package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
  spriteFactory *spriteFactory
  playerControls playerControls

  playership entity
}

var _ state = &playing{}

func (p *playing) init() {
  p.playerControls = playerControls{}
  p.playership.w = 36
  p.playership.h = 12
}

func (p *playing) tick(ms int)  (next string) {
  return ""
}

func (p *playing) receiveKeyEvent(event interaction.KeyEvent) (next string){
  if event.Key == "t" {
    return "title"
  }
  p.playerControls.receiveKeyEvent(event)
  return ""
}

func (p *playing) renderable() renderable {
  return renderables{
    p.spriteFactory.create("player_ship", p.playership.x, p.playership.y, 0),
    p.spriteFactory.create("bg_playing", 0, 0, 0),
  }
}

type playerControls struct {
  up bool
  down bool
  left bool
  right bool
  shoot bool
}

func (pc *playerControls) receiveKeyEvent(event interaction.KeyEvent) {
  if event.Type == interaction.KeyUp {
    pc.setByKey(event.Key, false)
  }
  if event.Type == interaction.KeyDown {
    pc.setByKey(event.Key, true)
  }
}

func (pc *playerControls) setByKey(key string, value bool) {
  switch key {
  case "w":
    pc.up = value
  case "s":
    pc.down = value
  case "a":
    pc.left = value
  case "d":
    pc.right = value
  case " ":
    pc.shoot = value
  }
}

type entity struct {
  x int
  y int
  w int
  h int
}

func (e entity) Left() int {
  return e.x
}

func (e entity) Right() int {
  return e.x + e.w
}

func (e entity) Top() int {
  return e.y
}

func (e entity) Bottom() int {
  return e.y + e.h
}

func (e entity) Center() (x, y int) {
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
    w: min(right.Right(), left.Right()),
    h: min(bottom.Bottom(), top.Bottom()),
  }, true
}
