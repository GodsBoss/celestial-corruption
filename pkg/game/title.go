package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type title struct {
  spriteFactory *spriteFactory
}

var _ state = &title{}

func (t *title) init() {}

func (t *title) tick(ms int) (next string) {
  return ""
}

func (t *title) receiveKeyEvent(event interaction.KeyEvent) (next string) {
  if event.Key == "p" {
    return "playing"
  }
  return ""
}

func (t *title) renderable() renderable {
  return renderables{
    t.spriteFactory.create("bg_title", 0, 0, 0),
    newText("Press [P] to start.", 5, 189).renderable(t.spriteFactory),
  }
}
