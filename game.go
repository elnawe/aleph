package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Add game_objects and input_handler
type Game struct {
	currentFrame int32
	exception    error
	//game_objects  []Game_object
	//input_handler Input_handler
	renderer   *sdl.Renderer
	game_state Game_state
	window     *sdl.Window
}

// TODO: add texture_manager *Texture
func (this *Game) init(title string, x, y, w, h int, fullscreen bool) {
	init_sdl()

	this.window = create_window(title, x, y, w, h, fullscreen)
	this.renderer = create_renderer(this.window, 0)

	this.game_state.change_state(init_menu_state())
}

func (this *Game) render() {
	clear_renderer(this.renderer)

	// TODO: Add logic to render (draw)

	present_renderer(this.renderer)
}

func (this *Game) update() {
	this.game_state.update()
}

func (this *Game) handle_inputs() {
	event := poll_event()

	switch event.(type) {
	case *sdl.QuitEvent:
		should_quit = true
	default:
		// TODO: Input_handler
	}
}
