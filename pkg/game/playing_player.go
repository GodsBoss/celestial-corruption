package game

import (
  "strconv"

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

  control playerControl

  hasQBomb bool

  madness int

  // madnessLevel goes from 0 (none) to 3 (maximum).
  madnessLevel int
}

func (p *player) Tick(ms int) {
  p.control.setSpeed(p)
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

func (p *player) shots() []shot {
  if p.reload > 0 {
    return nil
  }
  if !p.control.shouldShoot() {
    return nil
  }
  p.reload = playerReload
  _, y := p.Center()
  anim := &animation{
    maxFrame: 3,
    msPerFrame: 50,
  }
  anim.Randomize()
  return []shot{
    {
      Type: "player",
      entity: entity{
        x: p.Right(),
        y: y,
        w: 4,
        h: 4,
        dx: 200.0,
      },
      power: 100,
      animation: *anim,
      control: nopShotControl{},
    },
  }
}

func (p *player) Alive() bool {
  return p.health > 0
}

func (p *player) healthBar(sf *spriteFactory) renderable {
  m := strconv.Itoa(p.madnessLevel+1)
  r := renderables{
    sf.create("astronaut_" + m, 5, 5, 0),
  }
  for i := 0; i<10; i++ {
    typ := "empty"
    if p.health > i*playerMaxHealth/10{
      typ = "full"
    }
    r = append(
      r,
      sf.create("health_bar_" + typ + "_" + m, 20 + i * 10, 6, 0),
    )
  }
  return r
}

type keyboardControl struct {
  up bool
  down bool
  left bool
  right bool
  shoot bool
}

func (kc *keyboardControl) receiveKeyEvent(event interaction.KeyEvent) {
  if event.Type == interaction.KeyUp {
    kc.setByKey(event.Key, false)
  }
  if event.Type == interaction.KeyDown {
    kc.setByKey(event.Key, true)
  }
}

func (kc *keyboardControl) setByKey(key string, value bool) {
  switch key {
  case "w":
    kc.up = value
  case "s":
    kc.down = value
  case "a":
    kc.left = value
  case "d":
    kc.right = value
  case " ":
    kc.shoot = value
  }
}

func (kc *keyboardControl) combined() (int, int) {
  return boolInts[kc.right] - boolInts[kc.left], boolInts[kc.down] - boolInts[kc.up]
}

func (kc *keyboardControl) setSpeed(p *player) {
  dx, dy := kc.combined()
  pSpeed := playerSpeed
  if dx != 0 && dy != 0 {
    pSpeed = pSpeed / playerSpeedDiagonalFactor
  }
  p.dx = float64(dx) * pSpeed
  p.dy = float64(dy) * pSpeed
}

func (kc *keyboardControl) shouldShoot() bool {
  return kc.shoot
}

// Position of the player's ship in cinematics.
const (
  cinematicPlayerX = 50.0
  cinematicPlayerY = float64(gfxHeight) / 2.0
  cinematicDistanceThreshold = 2
)

type cinematicControl struct{}

func (cc *cinematicControl) setSpeed(p *player) {
  targetX, targetY := cinematicPlayerX - p.w / 2, cinematicPlayerY - p.h / 2
  pSpeed := playerSpeed / 2
  d := distance(p.x, p.y, targetX, targetY)
  if d < cinematicDistanceThreshold {
    p.dx = 0
    p.dy = 0
    return
  }
  p.dx = pSpeed * (targetX - p.x) / d
  p.dy = pSpeed * (targetY - p.y) / d
}

func (cc *cinematicControl) shouldShoot() bool {
  return false
}

type playerControl interface {
  setSpeed(p *player)
  shouldShoot() bool
}
