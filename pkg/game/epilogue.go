package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type epilogue struct{
  spriteFactory *spriteFactory
}

func (ep *epilogue) init() {}

func (ep *epilogue) tick(ms int) (next string) {
  return ""
}

func (ep *epilogue) receiveKeyEvent(event interaction.KeyEvent) (next string) {
  if event.Key == "t" {
    return "title"
  }
  return ""
}

func (ep *epilogue) renderable() renderable {
  return ep.spriteFactory.create("bg_epilogue", 0, 0, 0)
}
