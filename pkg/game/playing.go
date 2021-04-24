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

  enemies []enemy
}

var _ state = &playing{}

func (p *playing) init() {
  p.playerControls = playerControls{}
  p.playership = player{
    entity: entity{
      w: 36,
      h: 12,
      x: 20,
      y: float64(gfxHeight) / 2 - p.playership.h / 2,
    },
    health: playerMaxHealth,
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
  }
  p.playerShots = []shot{}
  p.enemies = []enemy{
    {
      entity: entity{
        x: 200,
        y: 50,
        w: 24,
        h: 24,
      },
      health: 1000,
      Type: "1",
      ramDamage: 800,
    },
    {
      entity: entity{
        x: 180,
        y: 150,
        w: 24,
        h: 24,
      },
      health: 1000,
      Type: "2",
      ramDamage: 800,
    },
  }
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

  newEnemies := make([]enemy, 0)
  for i := range p.enemies {
    p.enemies[i].Tick(ms)
    newShots := make([]shot, 0)
    for j := range p.playerShots {
      if _, collision := entityCollision(p.enemies[i].entity, p.playerShots[j].entity); collision {
        p.enemies[i].health -= p.playerShots[j].power
      } else {
        newShots = append(newShots, p.playerShots[j])
      }
    }
    p.playerShots = newShots
    if p.enemies[i].Alive() {
      newEnemies = append(newEnemies, p.enemies[i])
    }
  }
  p.enemies = newEnemies

  newEnemies = make([]enemy, 0)
  for i := range p.enemies {
    if _, collision := entityCollision(p.enemies[i].entity, p.playership.entity); collision {
      p.playership.health = max(0, p.playership.health - p.enemies[i].ramDamage)
      p.enemies[i].health = 0
    } else {
      newEnemies = append(newEnemies, p.enemies[i])
    }
  }
  p.enemies = newEnemies

  if !p.playership.Alive() {
    return "game_over"
  }

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
    p.spriteFactory.create("bg_playing", 0, 0, 0),
  }
  for i := range p.enemies {
    result = append(
      result,
      p.spriteFactory.create("enemy_"+ p.enemies[i].Type, int(p.enemies[i].x), int(p.enemies[i].y), 0),
    )
  }
  for i := range p.playerShots {
    result = append(
      result,
      p.spriteFactory.create("player_shot_1", int(p.playerShots[i].x), int(p.playerShots[i].y), 0),
    )
  }
  result = append(result, p.spriteFactory.create("player_ship", int(p.playership.x), int(p.playership.y), p.playership.Frame()))
  return result
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

type shot struct {
  entity

  // power is the shot's strength. A higher value means less successful hits are
  // necessary before destroying enemies.
  power int

  speedX float64
  speedY float64
}

func (sh *shot) Tick(ms int) {
  sh.x += normalizedSpeed(sh.speedX, ms)
}

func (sh *shot) Gone() bool {
  return sh.x > float64(gfxWidth)+10
}
