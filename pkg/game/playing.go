package game

import (
  "math"

  "github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
  spriteFactory *spriteFactory
  playerControls playerControls

  playership player
  playerShots []shot
}

var _ state = &playing{}

func (p *playing) init() {
  p.playerControls = playerControls{}
  p.playership.w = 36
  p.playership.h = 12
  p.playership.x = 20
  p.playership.y =float64(gfxHeight) / 2 - p.playership.h / 2
}

func (p *playing) tick(ms int)  (next string) {
  p.playership.Tick(ms)

  newShots := make([]shot, 0)
  for i := range p.playerShots {
    p.playerShots[i].Tick(ms)
    if !p.playerShots[i].Gone() {
      newShots = append(newShots, p.playerShots[i])
    }
  }
  p.playerShots = newShots

  dx, dy := p.playerControls.combined()
  pSpeed := normalizedSpeed(playerSpeed, ms)
  if dx != 0 && dy != 0 {
    pSpeed = pSpeed / playerSpeedDiagonalFactor
  }
  p.playership.x += float64(dx) * pSpeed
  p.playership.y += float64(dy) * pSpeed

  if p.playership.x < 5 {
    p.playership.x = 5
  }
  if p.playership.x > float64(gfxWidth - int(p.playership.w) - 5) {
    p.playership.x = float64(gfxWidth - int(p.playership.w) - 5)
  }
  if p.playership.y < 5 {
    p.playership.y = 5
  }
  if p.playership.y > float64(gfxHeight - int(p.playership.h) - 5) {
    p.playership.y = float64(gfxHeight - int(p.playership.h) - 5)
  }

  if p.playerControls.shoot {
    p.playerShots = append(p.playerShots, p.playership.shoot()...)
  }

  return ""
}

var playerSpeedDiagonalFactor = math.Sqrt(2.0)

func (p *playing) receiveKeyEvent(event interaction.KeyEvent) (next string){
  if event.Key == "t" {
    return "title"
  }
  p.playerControls.receiveKeyEvent(event)
  return ""
}

func (p *playing) renderable() renderable {
  result := renderables{
    p.spriteFactory.create("player_ship", int(p.playership.x), int(p.playership.y), 0),
    p.spriteFactory.create("bg_playing", 0, 0, 0),
  }
  for i := range p.playerShots {
    result = append(
      result,
      p.spriteFactory.create("player_shot_1", int(p.playerShots[i].x), int(p.playerShots[i].y), 0),
    )
  }
  return result
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

var boolInts = map[bool]int{
  false: 0,
  true: 1,
}

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

const (
  // playerSpeed is the speed of the player in in-game pixels per second.
  playerSpeed = 100.0

  playerReload = 100
)

type player struct {
  entity

  // reload is the time (in ms) weapon needs to reload. Can shoot if zero.
  reload int
}

func (p *player) Tick(ms int) {
  p.reload = max(p.reload - ms, 0)
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
      },
      speedX: 200.0,
    },
  }
}

type shot struct {
  entity

  speedX float64
  speedY float64
}

func (sh *shot) Tick(ms int) {
  sh.x += normalizedSpeed(sh.speedX, ms)
}

func (sh *shot) Gone() bool {
  return sh.x > float64(gfxWidth)+10
}
