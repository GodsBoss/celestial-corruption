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

  // p.playerControls.setSpeed(&p.playership)
  (&cinematicControl{}).setSpeed(&p.playership)

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

type shot struct {
  entity

  // power is the shot's strength. A higher value means less successful hits are
  // necessary before destroying enemies.
  power int
}

func (sh *shot) Tick(ms int) {
  sh.entity.Tick(ms)
}

func (sh *shot) Gone() bool {
  return sh.x > float64(gfxWidth)+10
}
