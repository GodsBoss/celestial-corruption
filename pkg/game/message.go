package game

type message struct {
  // duration is the duration (in ms) the message is shown.
  duration int

  // imageID is the image shown in the communication screen.
  imageID string

  imageAnimation animation
}

func (msg *message) Tick(ms int) {
  if msg == nil {
    return
  }
  msg.imageAnimation.Tick(ms)
  msg.duration -= ms
}

func (msg *message) Over() bool {
  return msg == nil || msg.duration <= 0
}

func (msg *message) renderable(sf *spriteFactory) renderable {
  return renderables{
    sf.create("message_container", 10, 146, 0),
    sf.create("message_" + msg.imageID, 14, 150, msg.imageAnimation.Frame()),
  }
}

const (
  messageX = 10
  messageY = 150
)
