package game

var playingTriggers map[string]trigger = make(map[string]trigger)

func init() {
  doAddTrigger := func(name string) func(*playing) {
    return doAddTriggerFromMap(name, playingTriggers)
  }

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
        doAddTrigger("starting_orders_1"),
      ),
    ),
    "starting_orders_1": newConditionalTrigger(
      killedAtLeast("practice", 3),
      multipleDos(
        doSetMessage(
          &message{
            duration: -1,
            imageID: "starting_orders_1",
            contents: lines(
              "Very good! Now that you are an expert in",
              "handling both your vessel and enemies,",
              "here are your orders. Fight your way",
              "through the alien swarms to Timos-1.",
              "There is a weapons laboratory.",
            ),
          },
        ),
        doSetTimer("starting_orders_1", seconds(12)),
        doAddTrigger("starting_orders_2"),
      ),
    ),
    "starting_orders_2": newConditionalTrigger(
      timerFinished("starting_orders_1"),
      multipleDos(
        doRemoveTimer("starting_orders_1"),
        doSetMessage(
          &message{
            duration: seconds(12),
            imageID: "starting_orders_2",
            contents: lines(
              "They developed a revolutionary weapon,",
              "the Quantum Bomb 9001. It will be loaded",
              "onto your ship. After that, fly to the",
              "alien homeworld and drop it to destroy",
              "the whole planet.",
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
