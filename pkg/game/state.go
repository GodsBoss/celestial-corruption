package game

import (
  "github.com/GodsBoss/gggg/pkg/interaction"
)

type state interface {
  // init is called whenever the game switches to this state.
  init()

  tick(ms int) (next string)

  receiveKeyEvent(event interaction.KeyEvent) (next string)

  sprites() renderable
}
