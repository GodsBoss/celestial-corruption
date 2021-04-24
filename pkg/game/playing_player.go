package game

const (
  // playerSpeed is the speed of the player in in-game pixels per second.
  playerSpeed = 100.0

  playerReload = 100
)

type player struct {
  entity

  // reload is the time (in ms) weapon needs to reload. Can shoot if zero.
  reload int
}

func (p *player) Tick(ms int) {
  p.reload = max(p.reload - ms, 0)
}

func (p *player) shoot() []shot {
  if p.reload > 0 {
    return nil
  }
  p.reload = playerReload
  _, y := p.Center()
  return []shot{
    {
      entity: entity{
        x: p.Right(),
        y: y,
        w: 4,
        h: 4,
      },
      speedX: 200.0,
      power: 100,
    },
  }
}
