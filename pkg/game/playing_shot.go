package game

type shot struct {
  Type string

  entity

  // power is the shot's strength. A higher value means less successful hits are
  // necessary before destroying enemies.
  power int

  animation

  control shotControl
}

func (sh *shot) Tick(ms int) {
  sh.control.control(sh)
  sh.entity.Tick(ms)
  sh.animation.Tick(ms)
}

func (sh *shot) Gone() bool {
  return (sh.dx > 0 && sh.x > float64(gfxWidth)+10) || (sh.dy > 0 && sh.y > float64(gfxHeight)) || (sh.dx < 0 && sh.x < -10) || (sh.dy < 0 && sh.y < -10)
}

type shotControl interface {
  control(*shot)
}

type nopShotControl struct{}

func (ctrl nopShotControl) control(*shot) {}
