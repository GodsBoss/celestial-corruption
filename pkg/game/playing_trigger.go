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

func multipleDos(dos ...func(*playing)) func(*playing) {
  return func(p *playing) {
    for i := range dos {
      dos[i](p)
    }
  }
}

func doSetMessage(msg *message) func(*playing) {
  return func(p *playing) {
    p.message = msg
  }
}

func doAddTriggerFromMap(name string, m map[string]trigger) func(*playing) {
  return func(p *playing) {
    p.additionalTriggers = append(p.additionalTriggers, m[name])
  }
}
