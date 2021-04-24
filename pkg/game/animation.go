package game

type animation struct {
  maxFrame int
  msPerFrame int

  current int
}

func (anim *animation) Tick(ms int) {
  anim.current += ms
  if anim.Frame() > anim.maxFrame {
    anim.current -= anim.Frame() * anim.msPerFrame
  }
}

func (anim *animation) Frame() int {
  return anim.current / anim.msPerFrame
}
