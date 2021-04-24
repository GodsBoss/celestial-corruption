package game

import (
  "math/rand"
)

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
              health: 1,
              ramDamage: 1000000, // These guys are *strong*.
              entity: entity{
                x: 250,
                y: 30,
                w: 24,
                h: 24,
              },
              control: nopEnemyControl{},
            },
            {
              Type: "practice",
              health: 1,
              ramDamage: 1000000,
              entity: entity{
                x: 270,
                y: 80,
                w: 24,
                h: 24,
              },
              control: nopEnemyControl{},
            },
            {
              Type: "practice",
              health: 1,
              ramDamage: 1000000,
              entity: entity{
                x: 220,
                y: 100,
                w: 24,
                h: 24,
              },
              control: nopEnemyControl{},
            },
          },
        ),
        doAddTrigger("player_dies"),
        doAddTrigger("starting_orders_1"),
      ),
    ),
    "player_dies": newConditionalTrigger(
      playerIsDead(),
      doSetNextState("game_over"),
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
        doSetTimer("starting_orders_1", seconds(1)),
        doAddTrigger("starting_orders_2"),
      ),
    ),
    "starting_orders_2": newConditionalTrigger(
      timerFinished("starting_orders_1"),
      multipleDos(
        doRemoveTimer("starting_orders_1"),
        doSetMessage(
          &message{
            duration: seconds(1),
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
        doAddTrigger("spawn_first_wave"),
        doAddTrigger("stop_first_wave"),
      ),
    ),
    "spawn_first_wave": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemy1, spawnEnemy2),
      maxEnemies: 10,
    },
    "stop_first_wave": newConditionalTrigger(
      allOf(
        killedAtLeast("1", 2),
        killedAtLeast("2", 2),
      ),
      multipleDos(
        doRemoveTrigger("spawn_first_wave"),
        doAddTrigger("wait_for_first_wave_end"),
      ),
    ),
    "wait_for_first_wave_end": newConditionalTrigger(
      allOf(
        enemiesAtMost(0),
        enemyShotsAtMost(0),
      ),
      multipleDos(
        doSetCinematicControl(),
        doSetMessage(
          &message{
            duration: seconds(3),
            imageID: "starting_orders_1",
            contents: lines(
              "Welcome to the Quantum Bomb Laboratories.",
              "Sorry that you had such a rough journey.",
              "The Quantum Bomb 9001 is already",
              "attached to your ship.",
            ),
          },
        ),
        doAddTrigger("wait_for_q_bomb_labs_msg_fading"),
      ),
    ),
    "wait_for_q_bomb_labs_msg_fading": newConditionalTrigger(
      invertCheck(showsMessage()),
      multipleDos(
        doSetKeyboardControl(),
        doSetQBomb(true),
        doSetMessage(
          &message{
            duration: seconds(3),
            imageID: "starting_orders_2",
            contents: lines(
              "You can now proceed. We will send you",
              "additional valuable information while",
              "you penetrate deeper into enemy",
              "territory. Take care!",
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

func spawnEnemy1() enemy {
  return enemy{
    Type: "1",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 50.0,
      switchTargetInterval: seconds(1),
    },
    health: 1,
    ramDamage: 100,
  }
}

func spawnEnemy2() enemy {
  return enemy{
    Type: "2",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyAlien() enemy {
  return enemy{
    Type: "alien",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyBrainy() enemy {
  return enemy{
    Type: "brainy",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyNightmare1() enemy {
  return enemy{
    Type: "nightmare_1",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyNightmare2() enemy {
  return enemy{
    Type: "nightmare_2",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyCorruptedEarthForces1() enemy {
  return enemy{
    Type: "corrupted_earth_forces_1",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}

func spawnEnemyCorruptedEarthForces2() enemy {
  return enemy{
    Type: "corrupted_earth_forces_2",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 24,
      h: 24,
    },
    animation: animation{
      maxFrame: 3,
      msPerFrame: 100,
    },
    control: &randomMovement{
      speed: 40.0,
      switchTargetInterval: seconds(2),
    },
    health: 1,
    ramDamage: 200,
  }
}


func seconds(s int) (ms int) {
  return s * 1000
}
