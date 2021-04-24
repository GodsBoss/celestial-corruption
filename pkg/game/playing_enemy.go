package game

// enemy is a simple enemy, e.g. a ship, flying brain, etc.
type enemy struct {
  entity

  Type string
  health int

  // ramDamage is the damage dealt if the enemy collides with the player.
  ramDamage int

  animation

  control enemyControl
}

func (e *enemy) Tick(ms int) {
  e.control.control(e)
  e.entity.Tick(ms)
  e.animation.Tick(ms)
}

func (e *enemy) Alive() bool {
  return e.health > 0
}

type enemyControl interface {
  control(*enemy)
}

type nopEnemyControl struct{}

func (ctrl nopEnemyControl) control(_ *enemy) {}
