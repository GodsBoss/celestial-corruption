package game

import (
  "math/rand"
)

type randomSpawner struct {
  spawnInterval int
  waitForSpawn int
  maxEnemies int
  spawn func()enemy
}

func (spwn *randomSpawner) run(p *playing) (keep bool) {
  if spwn.waitForSpawn < 0 && len(p.enemies) < spwn.maxEnemies {
    spwn.waitForSpawn += spwn.spawnInterval
    enemy := spwn.spawn()
    enemy.playing = p
    p.enemies = append(p.enemies, enemy)
  }
  return true
}

func (spwn *randomSpawner) Tick(ms int) {
  spwn.waitForSpawn -= ms
}

func spawnOneEnemyTypeRandomly(fs ...func()enemy) func() enemy {
  if len(fs) == 0 {
    panic("cannot call spawnOneEnemyTypeRandomly with no spawn funcs")
  }
  return func() enemy {
    return fs[rand.Intn(len(fs))]()
  }
}
