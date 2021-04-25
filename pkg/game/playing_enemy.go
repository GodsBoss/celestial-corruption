package game

import (
  "math/rand"
)

// enemy is a simple enemy, e.g. a ship, flying brain, etc.
type enemy struct {
  entity

  Type string
  health int

  // ramDamage is the damage dealt if the enemy collides with the player.
  ramDamage int

  animation

  control enemyControl

  // playing is a back reference to the playing state. We need it so the control
  // can utilize all of the state.
  playing *playing
}

func (e *enemy) Tick(ms int) {
  e.control.control(ms, e)
  e.entity.Tick(ms)
  e.animation.Tick(ms)
}

func (e *enemy) Alive() bool {
  return e.health > 0
}

type enemyControl interface {
  control(ms int, e *enemy)
}

type nopEnemyControl struct{}

func (ctrl nopEnemyControl) control(_ int, _ *enemy) {}

type enemyControlFunc func(ms int, e *enemy)

func (f enemyControlFunc) control(ms int, e *enemy) {
  f(ms, e)
}

type randomMovement struct {
  targetX float64
  targetY float64
  speed float64

  waitForNextTarget int
  switchTargetInterval int
}

func (mv *randomMovement) control(ms int, e *enemy) {
  mv.waitForNextTarget -= ms

  if mv.waitForNextTarget <= 0 {
    mv.waitForNextTarget += mv.switchTargetInterval
    mv.targetX, mv.targetY = rand.Float64() * float64(gfxWidth), rand.Float64() * float64(gfxHeight)
  }

  e.dx = 0
  e.dy = 0

  d := distance(mv.targetX, mv.targetY, e.x, e.y)

  if d > 0.1 {
    e.dx = (mv.targetX - e.x) * mv.speed / d
    e.dy = (mv.targetY - e.y) * mv.speed / d
  }
}

type wave1Shooter struct {
  rm randomMovement

  recovering int
  loadingShot int
  bulletSpeed float64
}

func (ws *wave1Shooter) control(ms int, e *enemy) {
  if ws.recovering <= 0 {
    ws.loadingShot -= ms
    if ws.loadingShot < 0 {
      ws.recovering = seconds(3)
      ws.loadingShot = seconds(1)

      cx, cy := e.Center()
      d := distance(e.playing.playership.x, e.playing.playership.y, cx, cy)

      e.playing.enemyShots = append(
        e.playing.enemyShots,
        shot{
          Type: "enemy_2",
          entity: entity{
            x: cx - 4,
            y: cy - 4,
            dx: (e.playing.playership.x - cx + 4) * ws.bulletSpeed / d,
            dy: (e.playing.playership.y - cy + 4) * ws.bulletSpeed / d,
          },
          power: 50,
          animation: animation{
            maxFrame: 3,
            msPerFrame: 50,
          },
          control: nopShotControl{},
        },
      )
    }
    return
  }
  ws.recovering -= ms
  ws.rm.control(ms, e)
}

type brainControl struct {
  targetX float64

  up bool
  dySwitchInterval int
  dySwitchChance float64
  dySwitchRecover int
}

func (bc *brainControl) control(ms int, e *enemy) {
  e.dx, e.dy = 0, 0
  if e.x > bc.targetX {
    e.dx = -brainSpeed
    return
  }
  bc.dySwitchRecover -= ms
  if bc.dySwitchRecover <= 0 {
    if rand.Float64() < bc.dySwitchChance {
      bc.up = !bc.up
    }
    bc.dySwitchRecover += bc.dySwitchInterval
  }
  if e.y < 0 {
    bc.up = false
  }
  if e.y > float64(gfxHeight) - e.h {
    bc.up = true
  }
  e.dy = brainSpeed
  if bc.up {
    e.dy *= -1
  }
}

const (
  brainSpeed = 50.0
)

type alienControl struct {
  targetX float64

  up bool
  dySwitchInterval int
  dySwitchChance float64
  dySwitchRecover int
}

func (ac *alienControl) control(ms int, e *enemy) {
  e.dx, e.dy = 0, 0
  if e.x > ac.targetX {
    e.dx = -alienSpeed
    return
  }
  ac.dySwitchRecover -= ms
  if ac.dySwitchRecover <= 0 {
    if rand.Float64() < ac.dySwitchChance {
      ac.up = !ac.up
    }
    ac.dySwitchRecover += ac.dySwitchInterval
    cx, cy := e.Center()
    e.playing.enemyShots = append(
      e.playing.enemyShots,
      shot{
        Type: "alien",
        entity: entity{
          x: cx - 5,
          y: cy,
          dx: -80,
          dy: 0,
        },
        power: 50,
        animation: animation{
          maxFrame: 3,
          msPerFrame: 50,
        },
        control: nopShotControl{},
      },
    )
  }
  if e.y < 0 {
    ac.up = false
  }
  if e.y > float64(gfxHeight) - e.h {
    ac.up = true
  }
  e.dy = alienSpeed
  if ac.up {
    e.dy *= -1
  }
}

const (
  alienSpeed = 35.0
)

type corruptedWarshipControl struct{
  rm randomMovement

  destSet bool
  destX float64
  destY float64

  recovery int
  shotRecovery int
  remainingShots int
}

func (ctrl *corruptedWarshipControl) control(ms int, e *enemy) {
  ctrl.recovery -= ms

  if ctrl.recovery > 0 {
    ctrl.rm.control(ms, e)
    return
  }

  if !ctrl.destSet {
    ctrl.destSet = true
    ctrl.destX = 250.0
    ctrl.destY = rand.Float64() * float64(gfxHeight - 40)
    ctrl.remainingShots = warshipShots
    ctrl.shotRecovery = 0
  }

  d := distance(ctrl.destX, ctrl.destY, e.x, e.y)
  if d > 5.0 {
    e.dx = (ctrl.destX - e.x) * warshipSpeed / d
    e.dy = (ctrl.destY - e.y) * warshipSpeed / d
    return
  }
  e.dx = 0
  e.dy = 0

  ctrl.shotRecovery -= ms
  if ctrl.shotRecovery > 0 {
    return
  }

  cx, cy := e.Center()
  e.playing.enemyShots = append(
    e.playing.enemyShots,
    shot{
      Type: "void",
      entity: entity{
        x: cx - 5,
        y: cy,
        dx: -80,
        dy: 0,
        w: 12,
        h: 12,
      },
      power: 50,
      animation: animation{
        maxFrame: 1,
        msPerFrame: 100,
      },
      control: nopShotControl{},
    },
  )

  ctrl.remainingShots--
  if ctrl.remainingShots <= 0 {
    // No more shots, back to recovery.
    ctrl.recovery = warshipRecovery
    ctrl.destSet = false
    return
  }

  ctrl.shotRecovery = warshipShotRecovery
}

const (
  warshipRecovery = 10000
  warshipShotRecovery = 250
  warshipShots = 4
  warshipSpeed = 30.0
)

type fighterControl struct {
  rm randomMovement

  shotRecovery int
}

func (fc *fighterControl) control(ms int, e *enemy) {
  fc.shotRecovery -= ms
  if fc.shotRecovery < 0 {
    fc.shotRecovery += fightShotRecovery

    cx, cy := e.Center()
    e.playing.enemyShots = append(
      e.playing.enemyShots,
      shot{
        Type: "mini_void",
        entity: entity{
          x: cx,
          y: cy,
          dx: -60,
          dy: 0,
          w: 4,
          h: 4,
        },
        power: 10,
        animation: animation{
          maxFrame: 3,
          msPerFrame: 50,
        },
        control: nopShotControl{},
      },
      shot{
        Type: "mini_void",
        entity: entity{
          x: cx,
          y: cy,
          dx: 60,
          dy: 0,
          w: 4,
          h: 4,
        },
        power: 10,
        animation: animation{
          maxFrame: 3,
          msPerFrame: 50,
        },
        control: nopShotControl{},
      },
      shot{
        Type: "mini_void",
        entity: entity{
          x: cx,
          y: cy,
          dx: 0,
          dy: -60,
          w: 4,
          h: 4,
        },
        power: 10,
        animation: animation{
          maxFrame: 3,
          msPerFrame: 50,
        },
        control: nopShotControl{},
      },
      shot{
        Type: "mini_void",
        entity: entity{
          x: cx,
          y: cy,
          dx: 0,
          dy: 60,
          w: 4,
          h: 4,
        },
        power: 10,
        animation: animation{
          maxFrame: 3,
          msPerFrame: 50,
        },
        control: nopShotControl{},
      },
    )
  }
  fc.rm.control(ms, e)
}

const (
  fightShotRecovery = 3000
)
