package game

type message struct {
  // duration is the duration (in ms) the message is shown. Message does not
  // automatically fade away if this is negative.
  duration int

  // imageID is the image shown in the communication screen.
  imageID string

  imageAnimation animation

  contents string
}

func (msg *message) Tick(ms int) {
  if msg == nil {
    return
  }
  msg.imageAnimation.Tick(ms)
  if msg.duration > 0 {
    msg.duration = max(0, msg.duration - ms)
  }
}

func (msg *message) Over() bool {
  return msg == nil || msg.duration == 0
}

func (msg *message) renderable(sf *spriteFactory) renderable {
  return renderables{
    sf.create("message_container", 10, 146, 0),
    sf.create("message_" + msg.imageID, 14, 150, msg.imageAnimation.Frame()),
    newText(msg.contents, 54, 150).renderable(sf),
  }
}

const (
  messageX = 10
  messageY = 150
  forever = -1
)
