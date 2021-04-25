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
  return renderables{
    ep.spriteFactory.create("bg_epilogue", 0, 0, 0),
    newText(
      lines(
        // --------------------------------------------------
        "Having descended deep into madness, you cannot",
        "appreciate the Dark Alien God tricking you into",
        "attacking Earth yourself. As the unstoppable Quantum",
        "fire burns through your own home planet, slowly",
        "consuming it and with it every single human,",
        "your maniacal laughter echoes",
        "within the otherwise empty",
        "chambers of your space vessel.",
      ),
      5,
      5,
    ).renderable(ep.spriteFactory),  }
}
