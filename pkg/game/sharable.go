package game

type sharableInt struct {
  value int
}

func (sh *sharableInt) Get() int {
  return sh.value
}

func (sh *sharableInt) Set(value int) {
  sh.value = value
}
