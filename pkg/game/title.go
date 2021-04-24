package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type title struct {}

var _ state = &title{}

func (t *title) init() {}

func (t *title) tick(ms int) (next string) {
  return ""
}

func (t *title) receiveKeyEvent(event interaction.KeyEvent) (next string) {
  return ""
}

func (t *title) sprites() map[string]sprite {
  return map[string]sprite{}
}
