package game

type shot struct {
  entity

  // power is the shot's strength. A higher value means less successful hits are
  // necessary before destroying enemies.
  power int

  animation
}

func (sh *shot) Tick(ms int) {
  sh.entity.Tick(ms)
  sh.animation.Tick(ms)
}

func (sh *shot) Gone() bool {
  return sh.x > float64(gfxWidth)+10
}
