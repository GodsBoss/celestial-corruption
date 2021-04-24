package game

import (
  "math"

  "github.com/GodsBoss/gggg/pkg/interaction"
)

type playing struct {
  spriteFactory *spriteFactory
  keyboardControl keyboardControl

  playership player
  playerShots []shot

  enemies []enemy

  message *message

  triggers []trigger

  // additionalTriggers are added to playing.triggers after the current triggers
  // have been processed. This avoids changing the triggers slice during its
  // iteration.
  additionalTriggers []trigger

  timers timers

  kills map[string]int
}

var _ state = &playing{}

func (p *playing) init() {
  p.keyboardControl = keyboardControl{}

  p.playership = player{
    entity: entity{
      w: 36,
      h: 12,
      x: 20,
      y: 0, // Needs h to be present.
    },
    health: playerMaxHealth,
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &p.keyboardControl,
  }
  p.playership.y = float64(gfxHeight) / 2 - p.playership.h / 2

  p.playerShots = nil
  p.enemies = nil
  p.message = nil
  p.triggers = []trigger{
    playingTriggers["init"],
  }
  p.additionalTriggers = nil
  p.timers = make(timers)
  p.kills = make(map[string]int)
}

func (p *playing) tick(ms int)  (next string) {
  p.tickAll(ms)
  p.removeGoneShots()
  p.handleEnemyShotCollisions()
  p.handleEnemyPlayerCollisions()
  p.removeOverMessage()

  p.handleTriggers()

  if !p.playership.Alive() {
    return "game_over"
  }

  p.playerShots = append(p.playerShots, p.playership.shots()...)

  return ""
}

func (p *playing) tickAll(ms int) {
  p.playership.Tick(ms)
  for i := range p.playerShots {
    p.playerShots[i].Tick(ms)
  }
  for i := range p.enemies {
    p.enemies[i].Tick(ms)
  }
  p.message.Tick(ms)
  p.timers.Tick(ms)
}

func (p *playing) removeGoneShots() {
  newShots := make([]shot, 0)
  for i := range p.playerShots {
    if !p.playerShots[i].Gone() {
      newShots = append(newShots, p.playerShots[i])
    }
  }
  p.playerShots = newShots
}

func (p *playing) handleEnemyShotCollisions() {
  newEnemies := make([]enemy, 0)
  for i := range p.enemies {
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
    } else {
      p.kills[p.enemies[i].Type]++
    }
  }
  p.enemies = newEnemies
}

func (p *playing) handleEnemyPlayerCollisions() {
  newEnemies := make([]enemy, 0)
  for i := range p.enemies {
    if _, collision := entityCollision(p.enemies[i].entity, p.playership.entity); collision {
      p.playership.health = max(0, p.playership.health - p.enemies[i].ramDamage)
      p.enemies[i].health = 0
      p.kills[p.enemies[i].Type]++
    } else {
      newEnemies = append(newEnemies, p.enemies[i])
    }
  }
  p.enemies = newEnemies
}

func (p *playing) removeOverMessage() {
  if p.message.Over() {
    p.message = nil
  }
}

func (p *playing) handleTriggers() {
  leftOverTriggers := make([]trigger, 0)
  for i := range p.triggers {
    if p.triggers[i].run(p) {
      leftOverTriggers = append(leftOverTriggers, p.triggers[i])
    }
  }
  p.triggers = append(leftOverTriggers, p.additionalTriggers...)
  p.additionalTriggers = nil
}

var playerSpeedDiagonalFactor = math.Sqrt(2.0)

func (p *playing) receiveKeyEvent(event interaction.KeyEvent) (next string){
  if event.Key == "t" {
    return "title"
  }
  p.keyboardControl.receiveKeyEvent(event)
  return ""
}

func (p *playing) renderable() renderable {
  result := renderables{
    p.spriteFactory.create("bg_playing", 0, 0, 0),
  }
  for i := range p.enemies {
    result = append(
      result,
      p.spriteFactory.create("enemy_"+ p.enemies[i].Type, int(p.enemies[i].x), int(p.enemies[i].y), p.enemies[i].Frame()),
    )
  }
  for i := range p.playerShots {
    result = append(
      result,
      p.spriteFactory.create("player_shot_1", int(p.playerShots[i].x), int(p.playerShots[i].y), p.playerShots[i].Frame()),
    )
  }
  result = append(result, p.spriteFactory.create("player_ship", int(p.playership.x), int(p.playership.y), p.playership.Frame()))
  if p.message != nil {
    result = append(result, p.message.renderable(p.spriteFactory))
  }
  return result
}

var boolInts = map[bool]int{
  false: 0,
  true: 1,
}
