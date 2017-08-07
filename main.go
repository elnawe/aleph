package main

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	DEBUG_MODE        bool   = true
	DEBUG_EVENTS      bool   = false
	DELAY_TIME        uint32 = uint32(1000.0 / FPS)
	FPS               uint32 = 60
	CAP_MAX_FPS       bool   = true
	WINDOW_TITLE      string = "Aleph"
	WINDOW_POSITION_X int    = sdl.WINDOWPOS_CENTERED
	WINDOW_POSITION_Y int    = sdl.WINDOWPOS_CENTERED
	WINDOW_WIDTH      int    = 800
	WINDOW_HEIGHT     int    = 600
	WINDOW_FULLSCREEN bool   = false
)

var (
	game        Game
	should_quit bool
	sync_mutex  sync.Mutex
)

func main() {
	init_game()
	main_loop()

	quit_sdl(game.window, game.renderer)
}

func init_game() {
	game.init(
		WINDOW_TITLE,
		WINDOW_POSITION_X,
		WINDOW_POSITION_Y,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		WINDOW_FULLSCREEN,
	)
}

func main_loop() {
	var frameStart, frameTime uint32

	for {
		if should_quit {
			return
		}

		frameStart = sdl.GetTicks()
		game.handle_inputs()
		game.update()
		game.render()
		frameTime = sdl.GetTicks() - frameStart

		do_delay_lock_fps(frameTime)
	}

	quit_sdl(game.window, game.renderer)
}
