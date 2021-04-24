package game

type timers map[string]int

func (t timers) Tick(ms int) {
  for name := range t {
    t[name] = max(0, t[name]-ms)
  }
}

// Set both adds and re-sets timers.
func (t timers) Set(name string, duration int) {
  t[name] = duration
}

func (t timers) Remove(name string) {
  if t.Exists(name) {
    delete(t, name)
  }
}

func (t timers) Exists(name string) bool {
  _, ok := t[name]
  return ok
}

func (t timers) Finished(name string) bool {
  dur, ok := t[name]
  return ok && dur <= 0
}
