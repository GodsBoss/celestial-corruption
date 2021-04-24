package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type gameOver struct {
  spriteFactory *spriteFactory
}

func (gov *gameOver) init() {}

func (gov *gameOver) tick(ms int) (next string) {
  return ""
}

func (gov *gameOver) receiveKeyEvent(event interaction.KeyEvent) (next string) {
  if event.Key == "t" {
    return "title"
  }
  return ""
}

func (gov *gameOver) renderable() renderable {
  return gov.spriteFactory.create("bg_game_over", 0, 0, 0)
}
