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
  return renderables{
    gov.spriteFactory.create("bg_game_over", 0, 0, 0),
    newText(
      lines(
        // --------------------------------------------------
        "You failed to halt the alien invasion. Your last",
        "thoughts are a nightmare of Earth overrun by a force",
        "of evil only to wipe out humanity, with no trace of",
        "it having ever existed. As your consciousness",
        "slowly vanishes, you accept the mercy of Death.",
        "",
        "At least your pitiful failure",
        "will be forgotten as well.",
      ),
      5,
      5,
    ).renderable(gov.spriteFactory),
    newText("Press [T] to return to title screen.", 5, 189).renderable(gov.spriteFactory),
  }
}
