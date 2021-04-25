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
              health: 1000,
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
              health: 1000,
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
              health: 1000,
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
        doSetTimer("starting_orders_1", seconds(8)),
        doAddTrigger("starting_orders_2"),
      ),
    ),
    "starting_orders_2": newConditionalTrigger(
      timerFinished("starting_orders_1"),
      multipleDos(
        doRemoveTimer("starting_orders_1"),
        doSetMessage(
          &message{
            duration: seconds(8),
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
        killedAtLeast("1", 20),
        killedAtLeast("2", 20),
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
            duration: seconds(8),
            imageID: "laboratory",
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
            duration: seconds(8),
            imageID: "laboratory",
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
      ),
    ),
    "spawn_aliens": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemyAlien, spawnEnemyBrainy),
      maxEnemies: 10,
    },
    "brainy_aliens_laboratory_message": newConditionalTrigger(
      killedAtLeast("brainy", 5),
      multipleDos(
        doSetMessage(
          &message{
            duration: seconds(8),
            imageID: "laboratory",
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
        doSetMadnessLevel(1),
      ),
    ),
    "stop_alien_spawn": newConditionalTrigger(
      allOf(
        killedAtLeast("alien", 20),
        killedAtLeast("brainy", 20),
      ),
      multipleDos(
        doRemoveTrigger("spawn_aliens"),
        doAddTrigger("wait_for_aliens_end"),
        doSetMessage(
          &message{
            duration: seconds(8),
            imageID: "astronaut",
            contents: lines(
              // ----------------------------------------
              "Your instruments catch very strange",
              "readings... it is as if entities are ahead",
              "which don't really belong to the material",
              "realm. Will the cannons even hurt them?",
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
        doAddTrigger("nightmares_lead_to_madness"),
        doAddTrigger("message_laboratory_warning"),
        doAddTrigger("what_are_these"),
        doAddTrigger("whisperings_1"),
        doAddTrigger("whisperings_2"),
      ),
    ),
    "spawn_nightmares": &randomSpawner{
      spawnInterval: seconds(1),
      spawn: spawnOneEnemyTypeRandomly(spawnEnemyNightmare1, spawnEnemyNightmare2),
      maxEnemies: 10,
    },
    "what_are_these": newConditionalTrigger(
      oneOf(
        killedAtLeast("nightmare_1", 1),
        killedAtLeast("nightmare_2", 1),
      ),
      doSetMessage(
        &message{
          duration: seconds(1),
          imageID: "astronaut",
          contents: "What are these?",
        },
      ),
    ),
    "nightmares_lead_to_madness": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 10),
        killedAtLeast("nightmare_2", 10),
      ),
      multipleDos(
        doSetMadnessLevel(2),
        doSetMessage(
          &message{
            duration: seconds(8),
            imageID: "astronaut",
            contents: lines(
              // ----------------------------------------
              "Luckily, you checked your instruments.",
              "Seems like your course wasn't completely",
              "correct. It is now and you will reach the",
              "enemy's home planet in time!",
            ),
          },
        ),
      ),
    ),
    "message_laboratory_warning": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 15),
        killedAtLeast("nightmare_2", 15),
      ),
      multipleDos(
        doSetMessage(
          &message{
            duration: seconds(6),
            imageID: "laboratory",
            contents: lines(
              // ----------------------------------------
              "Emanations... dangerous... etheral...",
              "... without protection ... no chance ...",
              "... dark evil alien god ...",
              "(static)",
              "... your immediate return!",
            ),
          },
        ),
        doAddTrigger("after_message_laboratory_warning"),
        doSetMadnessLevel(3),
      ),
    ),
    "after_message_laboratory_warning": newConditionalTrigger(
      invertCheck(showsMessage()),
      doSetMessage(
        &message{
          duration: seconds(8),
          imageID: "astronaut",
          contents: lines(
            // ----------------------------------------
            "The message came in scrambled. Why should",
            "you return? Victory is close! Maybe the",
            "aliens have overrun the laboratory and try",
            "to stop you. That's it! They become",
            "desperate. Good!",
          ),
        },
      ),
    ),
    "whisperings_1": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 20),
        killedAtLeast("nightmare_2", 20),
      ),
      doSetMessage(
        &message{
          duration: seconds(8),
          imageID: "astronaut",
          contents: lines(
            // ----------------------------------------
            "It feels like you can hear whispers. Are",
            "these coming from those strange enemies?",
            "Whatever tricks they have, they cannot",
            "stop you from completing your mission.",
          ),
        },
      ),
    ),
    "whisperings_2": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 25),
        killedAtLeast("nightmare_2", 25),
      ),
      doSetMessage(
        &message{
          duration: seconds(8),
          imageID: "astronaut",
          contents: lines(
            // ----------------------------------------
            "You have a nagging feeling in the back of",
            "your head that something is wrong. Pretty",
            "sure those whisperings have the purpose",
            "to let you abort the mission. Ha!",
          ),
        },
      ),
    ),
    "stop_nightmare_spawn": newConditionalTrigger(
      allOf(
        killedAtLeast("nightmare_1", 30),
        killedAtLeast("nightmare_2", 30),
      ),
      multipleDos(
        doRemoveTrigger("spawn_nightmares"),
        doAddTrigger("wait_for_nightmare_end"),
        doSetMessage(
          &message{
            duration: seconds(8),
            imageID: "astronaut",
            contents: lines(
              // ----------------------------------------
              "Insane! The aliens have copied the designs",
              "of Earth ships and twisted them to make",
              "them more powerful. You should still be",
              "able to defeat them, at least they use the",
              "same tactics, known to you.",
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
        killedAtLeast("corrupted_earth_forces_1", 20),
        killedAtLeast("corrupted_earth_forces_2", 20),
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
            duration: seconds(8),
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
      speed: 40.0,
      switchTargetInterval: seconds(1),
    },
    health: 800,
    ramDamage: 75,
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
        speed: 30.0,
        switchTargetInterval: seconds(2),
      },
      bulletSpeed: 80,
    },
    health: 600,
    ramDamage: 75,
  }
}

func spawnEnemyAlien() enemy {
  return enemy{
    Type: "alien",
    entity: entity {
      x: 330,
      y: rand.Float64() * float64(gfxHeight),
      w: 16,
      h: 12,
    },
    animation: animation{
      maxFrame: 5,
      msPerFrame: 100,
    },
    control: &alienControl{
      targetX: 200.0 + rand.Float64() * 100,
      up: rand.Float64() > 0.5,
      dySwitchInterval: seconds(2),
      dySwitchChance: 0.5,
    },
    health: 500,
    ramDamage: 40,
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
    control: &brainControl{
      targetX: 150.0 + rand.Float64() * 50,
      up: rand.Float64() > 0.5,
      dySwitchInterval: seconds(1),
      dySwitchChance: 0.5,
    },
    health: 1000,
    ramDamage: 80,
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
    health: 1000,
    ramDamage: 100,
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
    health: 500,
    ramDamage: 50,
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
    control: &corruptedWarshipControl{
      rm: randomMovement{
        speed: warshipSpeed / 2.0,
        switchTargetInterval: seconds(2),
      },
      recovery: warshipRecovery,
    },
    health: 1800,
    ramDamage: 500,
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
    control: &fighterControl{
      rm: randomMovement{
        speed: 40.0,
        switchTargetInterval: seconds(2),
      },
      shotRecovery: fightShotRecovery,
    },
    health: 800,
    ramDamage: 50,
  }
}


func seconds(s int) (ms int) {
  return s * 1000
}
