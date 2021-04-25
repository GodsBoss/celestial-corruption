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
  "bg_epilogue": {
    x: 640,
    y: 600,
    w: 320,
    h: 200,
  },
  "player_ship": {
    x: 2,
    y: 2,
    w: 36,
    h: 12,
  },
  "q_bomb": {
    x: 2,
    y: 72,
    w: 36,
    h: 12,
  },
  "player_shot": {
    x: 2,
    y: 15,
    w: 4,
    h: 4,
  },
  "enemy_2_shot": {
    x: 100,
    y: 16,
    w: 8,
    h: 8,
  },
  "alien_shot": {
    x: 100,
    y: 25,
    w: 8,
    h: 8,
  },
  "void_shot": {
    x: 100,
    y: 34,
    w: 12,
    h: 12,
  },
  "enemy_1": {
    x: 2,
    y: 86,
    w: 24,
    h: 24,
  },
  "enemy_2": {
    x: 2,
    y: 112,
    w: 24,
    h: 24,
  },
  "cinematic_screen": {
    x: 150,
    y: 0,
    w: 48,
    h: 48,
  },
  "enemy_practice": {
    x: 198,
    y: 0,
    w: 24,
    h: 24,
  },
  "enemy_alien": {
    x: 198,
    y: 24,
    w: 16,
    h: 12,
  },
  "enemy_brainy": {
    x: 2,
    y: 138,
    w: 16,
    h: 32,
  },
  "enemy_nightmare_1": {
    x: 2,
    y: 174,
    w: 36,
    h: 36,
  },
  "enemy_nightmare_2": {
    x: 16,
    y: 211,
    w: 16,
    h: 16,
  },
  "enemy_corrupted_earth_forces_1": {
    x: 16,
    y: 248,
    w: 72,
    h: 24,
  },
  "enemy_corrupted_earth_forces_2": {
    x: 2,
    y: 275,
    w: 16,
    h: 16,
  },
  "message_container": {
    x: 340,
    y: 0,
    w: 300,
    h: 44,
  },
  "message_test": {
    x: 370,
    y: 86,
    w: 36,
    h: 36,
  },
  "message_target_practice": {
    x: 370,
    y: 122,
    w: 36,
    h: 36,
  },
  "message_starting_orders_1": {
    x: 370,
    y: 158,
    w: 36,
    h: 36,
  },
  "message_starting_orders_2": {
    x: 442,
    y: 122,
    w: 36,
    h: 36,
  },
  "message_larboratory": {
    x: 370,
    y: 194,
    w: 36,
    h: 36,
  },
  "message_astronaut": {
    x: 406,
    y: 194,
    w: 36,
    h: 36,
  },
  "message_planet": {
    x: 442,
    y: 194,
    w: 36,
    h: 36,
  },
  "madness_particles": {
    x: 100,
    y: 50,
    w: 16,
    h: 16,
  },
  "astronaut_1": {
    x: 165,
    y: 50,
    w: 12,
    h: 12,
  },
  "astronaut_2": {
    x: 177,
    y: 50,
    w: 12,
    h: 12,
  },
  "astronaut_3": {
    x: 189,
    y: 50,
    w: 12,
    h: 12,
  },
  "astronaut_4": {
    x: 201,
    y: 50,
    w: 12,
    h: 12,
  },
  "health_bar_full_1": {
    x: 165,
    y: 63,
    w: 8,
    h: 8,
  },
  "health_bar_full_2": {
    x: 177,
    y: 63,
    w: 8,
    h: 8,
  },
  "health_bar_full_3": {
    x: 189,
    y: 63,
    w: 8,
    h: 8,
  },
  "health_bar_full_4": {
    x: 201,
    y: 63,
    w: 8,
    h: 8,
  },
  "health_bar_empty_1": {
    x: 165,
    y: 72,
    w: 8,
    h: 8,
  },
  "health_bar_empty_2": {
    x: 177,
    y: 72,
    w: 8,
    h: 8,
  },
  "health_bar_empty_3": {
    x: 189,
    y: 72,
    w: 8,
    h: 8,
  },
  "health_bar_empty_4": {
    x: 201,
    y: 72,
    w: 8,
    h: 8,
  },
  "char_a": {
    x: 0,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_b": {
    x: 5,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_c": {
    x: 10,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_d": {
    x: 15,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_e": {
    x: 20,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_f": {
    x: 25,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_g": {
    x: 30,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_h": {
    x: 35,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_i": {
    x: 40,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_j": {
    x: 45,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_k": {
    x: 50,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_l": {
    x: 55,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_m": {
    x: 60,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_n": {
    x: 65,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_o": {
    x: 70,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_p": {
    x: 75,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_q": {
    x: 80,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_r": {
    x: 85,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_s": {
    x: 90,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_t": {
    x: 95,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_u": {
    x: 100,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_v": {
    x: 105,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_w": {
    x: 110,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_x": {
    x: 115,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_y": {
    x: 120,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_z": {
    x: 125,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_ ": {
    x: 130,
    y: 762,
    w: 5,
    h: 6,
  },
  "char_0": {
    x: 0,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_1": {
    x: 5,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_2": {
    x: 10,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_3": {
    x: 15,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_4": {
    x: 20,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_5": {
    x: 25,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_6": {
    x: 30,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_7": {
    x: 35,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_8": {
    x: 40,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_9": {
    x: 45,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_.": {
    x: 50,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_,": {
    x: 55,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_:": {
    x: 60,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_!": {
    x: 65,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_?": {
    x: 70,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_(": {
    x: 75,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_)": {
    x: 80,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_[": {
    x: 85,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_]": {
    x: 90,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_-": {
    x: 95,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_=": {
    x: 100,
    y: 769,
    w: 5,
    h: 6,
  },
  "char__": {
    x: 105,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_\"": {
    x: 110,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_'": {
    x: 115,
    y: 769,
    w: 5,
    h: 6,
  },
  "char_+": {
    x: 120,
    y: 769,
    w: 5,
    h: 6,
  },
}
