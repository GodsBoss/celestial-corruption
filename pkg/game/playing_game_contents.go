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
              // Maximum message width:
              // ----------------------------------------
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
              animation: animation{
                maxFrame: 3,
                msPerFrame: 145,
              },
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
              animation: animation{
                maxFrame: 3,
                msPerFrame: 150,
              },
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
              animation: animation{
                maxFrame: 3,
                msPerFrame: 155,
              },
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
            imageAnimation: animation{
              maxFrame: 3,
              msPerFrame: 100,
            },
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
            imageID: "larboratory",
            contents: lines(
              "Welcome to the Quantum Bomb Laboratories.",
              "Sorry that you had such a rough journey.",
              "The Quantum Bomb 9001 will be attached to",
              "to your ship immediately.",
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
            imageID: "larboratory",
            contents: lines(
              "You can now proceed. We will send you",
              "additional valuable information while",
              "you penetrate deeper into enemy",
              "territory. Take care!",
            ),
          },
        ),
        doAddTrigger("spawn_aliens"),
        doAddTrigger("stop_alien_spawn"),
        doAddTrigger("brainy_aliens_laboratory_message"),
        doAddTrigger("brainy_aliens_lead_to_madness"),
      ),
    ),
    "spawn_aliens": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemyAlien, spawnEnemyBrainy),
      maxEnemies: 10,
    },
    "brainy_aliens_laboratory_message": newConditionalTrigger(
      killedAtLeast("brainy", 1),
      doSetMessage(
        &message{
          duration: seconds(3),
          imageID: "larboratory",
          contents: lines(
            // ----------------------------------------
            "Here is the laboratory. These aliens seem",
            "to emanate some unknown kind of energy.",
            "We don't know wether this happens on",
            "purpose or if they are just that way.",
            "Be careful!",
          ),
        },
      ),
    ),
    "brainy_aliens_lead_to_madness": newConditionalTrigger(
      killedAtLeast("brainy", 1),
      doSetMadnessLevel(1),
    ),
    "stop_alien_spawn": newConditionalTrigger(
      allOf(
        killedAtLeast("alien", 2),
        killedAtLeast("brainy", 2),
      ),
      multipleDos(
        doRemoveTrigger("spawn_aliens"),
        doAddTrigger("wait_for_aliens_end"),
        doSetMessage(
          &message{
            duration: seconds(3),
            imageID: "TODO",
            contents: lines(
              "Aliens defeated,",
              "Nightmare ahead!",
            ),
          },
        ),
      ),
    ),
    "wait_for_aliens_end": newConditionalTrigger(
      allOf(
        enemiesAtMost(0),
        enemyShotsAtMost(0),
      ),
      multipleDos(
        doAddTrigger("spawn_nightmares"),
        doAddTrigger("stop_nightmare_spawn"),
      ),
    ),
    "spawn_nightmares": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemyNightmare1, spawnEnemyNightmare2),
      maxEnemies: 10,
    },
    "stop_nightmare_spawn": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 2),
        killedAtLeast("nightmare_2", 2),
      ),
      multipleDos(
        doRemoveTrigger("spawn_nightmares"),
        doAddTrigger("wait_for_nightmare_end"),
        doSetMessage(
          &message{
            duration: seconds(3),
            imageID: "TODO",
            contents: lines(
              "Nightmares defeated,",
              "Corrupted Earth Forces ahead!",
            ),
          },
        ),
      ),
    ),
    "wait_for_nightmare_end": newConditionalTrigger(
      allOf(
        enemiesAtMost(0),
        enemyShotsAtMost(0),
      ),
      multipleDos(
        doAddTrigger("spawn_corrupted_earth_forces"),
        doAddTrigger("stop_corrupted_earth_forces_spawn"),
      ),
    ),
    "spawn_corrupted_earth_forces": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemyCorruptedEarthForces1, spawnEnemyCorruptedEarthForces2),
      maxEnemies: 10,
    },
    "stop_corrupted_earth_forces_spawn": newConditionalTrigger(
      allOf(
        killedAtLeast("corrupted_earth_forces_1", 2),
        killedAtLeast("corrupted_earth_forces_2", 2),
      ),
      multipleDos(
        doRemoveTrigger("spawn_corrupted_earth_forces"),
        doAddTrigger("wait_for_corrupted_earth_forces_end"),
      ),
    ),
    "wait_for_corrupted_earth_forces_end": newConditionalTrigger(
      allOf(
        enemiesAtMost(0),
        enemyShotsAtMost(0),
      ),
      multipleDos(
        doSetMessage(
          &message{
            duration: seconds(5),
            imageID: "planet",
            contents: lines(
              // ----------------------------------------
              "The enemy's home planet is near. It looks",
              "like a twisted version of Earth, deeply",
              "corrupted. Luckily, the Quantum Bomb will",
              "end its existence. Humanity will be",
              "safe again!",
            ),
          },
        ),
        doAddTrigger("end_the_game"),
      ),
    ),
    "end_the_game": newConditionalTrigger(
      invertCheck(showsMessage()),
      doSetNextState("epilogue"),
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
      maxFrame: 7,
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
    control: &wave1Shooter{
      rm: randomMovement{
        speed: 40.0,
        switchTargetInterval: seconds(2),
      },
      bulletSpeed: 25,
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
      w: 16,
      h: 24,
    },
    animation: animation{
      maxFrame: 5,
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
      w: 16,
      h: 32,
    },
    animation: animation{
      maxFrame: 11,
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
      w: 36,
      h: 36,
    },
    animation: animation{
      maxFrame: 1,
      msPerFrame: 500,
    },
    control: &randomMovement{
      speed: 20.0,
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
      w: 16,
      h: 16,
    },
    animation: animation{
      maxFrame: 5,
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
      w: 72,
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
      w: 16,
      h: 16,
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
