package game

var playingTriggers map[string]trigger = make(map[string]trigger)

func init() {
  pTriggers := map[string]trigger{
    "init": newConditionalTrigger(
      alwaysOK,
      multipleDos(
        doSetMessage(
          &message{
            duration: forever,
            imageID: "target_practice",
            contents: lines(
              "Hello, commander! We don't have much time.",
              "Shoot the targets for practice!",
              "We continue when they are destroyed.",
              "Do not touch them! They'll kill you.",
            ),
            imageAnimation: animation{
              maxFrame: 1,
              msPerFrame: 1000,
            },
          },
        ),
        doAddEnemies(
          []enemy{
            {
              Type: "practice",
              health: 1000,
              ramDamage: 1000000, // These guys are *strong*.
              entity: entity{
                x: 250,
                y: 30,
                w: 24,
                h: 24,
              },
            },
            {
              Type: "practice",
              health: 1000,
              ramDamage: 1000000,
              entity: entity{
                x: 270,
                y: 80,
                w: 24,
                h: 24,
              },
            },
            {
              Type: "practice",
              health: 1000,
              ramDamage: 1000000,
              entity: entity{
                x: 220,
                y: 100,
                w: 24,
                h: 24,
              },
            },
          },
        ),
        doAddTriggerFromMap("starting_orders", playingTriggers),
      ),
    ),
    "starting_orders": newConditionalTrigger(
      killedAtLeast("practice", 3),
      multipleDos(
        doSetMessage(
          &message{
            duration: seconds(30),
            imageID: "",
            contents: lines(
              "Very good!",
            ),
          },
        ),
      ),
    ),
  }
  for tID := range pTriggers {
    playingTriggers[tID] = pTriggers[tID]
  }
}

func seconds(s int) (ms int) {
  return s * 1000
}
