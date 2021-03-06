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

// TODO: This globals could be a problem. Redo this if needed.
var (
	should_quit     bool
	sync_mutex      sync.Mutex
	texture_manager Texture_Manager
	the_game        Game
)

func init() {
	texture_manager = Texture_Manager{
		game:        &the_game,
		texture_map: map[string]*sdl.Texture{},
	}

	the_game.init(
		WINDOW_TITLE,
		WINDOW_POSITION_X,
		WINDOW_POSITION_Y,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		WINDOW_FULLSCREEN,
		&texture_manager,
	)
}

func main() {
	main_loop()
	sdl_quit(the_game.window, the_game.renderer)
}

func main_loop() {
	var frameStart, frameTime uint32

	for {
		if should_quit {
			return
		}
		frameStart = sdl_get_ticks()
		the_game.input()
		the_game.update()
		the_game.render(&texture_manager)
		frameTime = sdl_get_ticks() - frameStart

		do_delay_lock_fps(frameTime)
	}
}
