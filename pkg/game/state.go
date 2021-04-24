package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type state interface {
  // init is called whenever the game switches to this state.
  init()

  tick(ms int) (next string)

  receiveKeyEvent(event interaction.KeyEvent) (next string)

  sprites() map[string]sprite
}

type sprite interface {
  ID() string
  Frame() int
  X() int
  Y() int
}
