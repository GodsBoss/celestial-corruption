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

}
