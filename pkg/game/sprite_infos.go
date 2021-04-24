package game

var spriteInfos = map[string]spriteInfo{
  "bg_title": {
    x: 640,
    y: 0,
    w: 320,
    h: 200,
  },
  "bg_playing": {
    x: 640,
    y: 200,
    w: 320,
    h: 200,
  },
  "bg_game_over": {
    x: 640,
    y: 400,
    w: 320,
    h: 200,
  },
  "player_ship": {
    x: 2,
    y: 2,
    w: 36,
    h: 12,
  },
  "player_shot_1": {
    x: 2,
    y: 15,
    w: 4,
    h: 4,
  },
  "enemy_1": {
    x: 2,
    y: 20,
    w: 24,
    h: 24,
  },
  "enemy_2": {
    x: 2,
    y: 46,
    w: 24,
    h: 24,
  },
  "cinematic_screen": {
    x: 150,
    y: 0,
    w: 48,
    h: 48,
  },
}
