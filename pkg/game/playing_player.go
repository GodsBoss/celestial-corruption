package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

const (
  // playerSpeed is the speed of the player in in-game pixels per second.
  playerSpeed = 100.0

  playerReload = 100

  playerMaxHealth = 1000
)

type player struct {
  entity

  // reload is the time (in ms) weapon needs to reload. Can shoot if zero.
  reload int

  health int

  animation
}

func (p *player) Tick(ms int) {
  p.reload = max(p.reload - ms, 0)
  p.animation.Tick(ms)
  p.entity.Tick(ms)

  if p.x < 5 {
    p.x = 5
  }
  if p.x > float64(gfxWidth - int(p.w) - 5) {
    p.x = float64(gfxWidth - int(p.w) - 5)
  }
  if p.y < 5 {
    p.y = 5
  }
  if p.y > float64(gfxHeight - int(p.h) - 5) {
    p.y = float64(gfxHeight - int(p.h) - 5)
  }
}

func (p *player) shoot() []shot {
  if p.reload > 0 {
    return nil
  }
  p.reload = playerReload
  _, y := p.Center()
  return []shot{
    {
      entity: entity{
        x: p.Right(),
        y: y,
        w: 4,
        h: 4,
        dx: 200.0,
      },
      power: 100,
    },
  }
}

func (p *player) Alive() bool {
  return p.health > 0
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

func (pc *playerControls) combined() (int, int) {
  return boolInts[pc.right] - boolInts[pc.left], boolInts[pc.down] - boolInts[pc.up]
}

func (pc *playerControls) setSpeed(p *player) {
  dx, dy := pc.combined()
  pSpeed := playerSpeed
  if dx != 0 && dy != 0 {
    pSpeed = pSpeed / playerSpeedDiagonalFactor
  }
  p.dx = float64(dx) * pSpeed
  p.dy = float64(dy) * pSpeed
}
