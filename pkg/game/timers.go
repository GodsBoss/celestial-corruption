package game

type timers map[string]int

func (t timers) Tick(ms int) {
  for name := range t {
    t[name] = max(0, t[name]-ms)
  }
}

func (t timers) Add(name string, duration int) {
  t[name] = duration
}

func (t timers) Exists(name string) bool {
  _, ok := t[name]
  return ok
}

func (t timers) Finished(name string) bool {
  dur, ok := t[name]
  return ok && dur <= 0
}
