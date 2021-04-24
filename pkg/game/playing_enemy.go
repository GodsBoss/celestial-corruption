package game

// enemy is a simple enemy, e.g. a ship, flying brain, etc.
type enemy struct {
  entity

  Type string
  health int
}

func (e *enemy) Tick(ms int) {}

func (e *enemy) Alive() bool {
  return e.health > 0
}
