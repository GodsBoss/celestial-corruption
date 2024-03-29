package game

type trigger interface {
  // run runs the trigger. If it returns false, it is removed.
  run(*playing) (keep bool)
}

type triggerFunc func(*playing) (keep bool)

func (f triggerFunc) run(p *playing) (keep bool){
  return f(p)
}

type conditionalTrigger struct {
  check func(*playing) bool
  do func(*playing)
}

func newConditionalTrigger(check func(*playing) bool, do func(*playing)) trigger {
  return conditionalTrigger{
    check: check,
    do: do,
  }
}

func (t conditionalTrigger) run(p *playing) (keep bool) {
  if t.check(p) {
    t.do(p)
    return false
  }
  return true
}

func constCheckFunc(b bool) func(*playing) bool {
  return func(_ *playing) bool {
    return b
  }
}

var (
  alwaysOK = constCheckFunc(true)
  never = constCheckFunc(false)
)

func killedAtLeast(typ string, count int) func(*playing) bool {
  return func(p *playing) bool {
    return p.kills[typ] >= count
  }
}

func enemyShotsAtMost(count int) func(*playing) bool {
  return func(p *playing) bool {
    return len(p.enemyShots) <= count
  }
}

func enemiesAtMost(count int) func(*playing) bool {
  return func(p *playing) bool {
    return len(p.enemies) <= count
  }
}

func timerFinished(name string) func(*playing) bool {
  return func(p *playing) bool {
    return p.timers.Finished(name)
  }
}

func playerIsDead() func(*playing) bool {
  return func(p *playing) bool {
    return !p.playership.Alive()
  }
}

func showsMessage() func(*playing) bool {
  return func(p *playing) bool {
    return p.message != nil
  }
}

func madnessAtLeast(amount int) func(*playing) bool {
  return func(p *playing) bool {
    return p.playership.madness >= amount
  }
}

func invertCheck(f func(*playing) bool) func(*playing) bool {
  return func(p *playing) bool {
    return !f(p)
  }
}

func allOf(fs ...func(*playing) bool) func(*playing) bool {
  return func(p *playing) bool {
    for i := range fs {
      if !fs[i](p) {
        return false
      }
    }
    return true
  }
}

func oneOf(fs ...func(*playing) bool) func(*playing) bool {
  return func(p *playing) bool {
    for i := range fs {
      if fs[i](p) {
        return true
      }
    }
    return false
  }
}

func multipleDos(dos ...func(*playing)) func(*playing) {
  return func(p *playing) {
    for i := range dos {
      dos[i](p)
    }
  }
}

func doSetMessage(msg *message) func(*playing) {
  return func(p *playing) {
    cp := *msg
    p.message = &cp
  }
}

func doAddTriggerFromMap(name string, m map[string]trigger) func(*playing) {
  return func(p *playing) {
    p.additionalTriggers[name] = m[name]
  }
}

func doRemoveTrigger(name string) func(*playing) {
  return func(p *playing) {
    p.triggersToRemove = append(p.triggersToRemove, name)
  }
}

func doAddEnemies(enemies []enemy) func(*playing){
  return func(p *playing) {
    for i := range enemies{
      enemies[i].playing = p
    }
    p.enemies = append(p.enemies, enemies...)
  }
}

func doSetTimer(name string, duration int) func(*playing) {
  return func(p *playing) {
    p.timers.Set(name, duration)
  }
}

func doRemoveTimer(name string) func(*playing) {
  return func(p *playing) {
    p.timers.Remove(name)
  }
}

func doSetNextState(nextState string) func(*playing) {
  return func(p *playing) {
    p.nextState = nextState
  }
}

func doSetCinematicControl() func(*playing) {
  return func(p *playing) {
    p.playership.control = &cinematicControl{}
  }
}

func doSetKeyboardControl() func(*playing) {
  return func(p *playing) {
    p.playership.control = &p.keyboardControl
  }
}

func doSetQBomb(hasQBomb bool) func(*playing) {
  return func(p *playing) {
    p.playership.hasQBomb = hasQBomb
  }
}

func doSetMadnessLevel(level int) func(*playing) {
  return func(p *playing) {
    p.playership.madnessLevel = level
  }
}
