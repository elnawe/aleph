package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Add game_objects
type Game struct {
	currentFrame int32
	exception    error
	//game_objects  []Game_object
	input_handler Input_handler
	renderer      *sdl.Renderer
	game_state    Game_state
	window        *sdl.Window
}

// TODO: add texture_manager *Texture
func (this *Game) init(title string, x, y, w, h int, fullscreen bool) {
	sdl_init()

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

	switch event_type := event.(type) {
	case *sdl.QuitEvent:
		// TODO: Add sync_mutex superset
		sync_mutex.Lock()
		should_quit = true
		sync_mutex.Unlock()
	case *sdl.KeyDownEvent:
		// TODO: Add input handler to handle these
		// handle_keydown_event(event_type.Keysym.Scancode)
	case *sdl.KeyUpEvent:
		// TODO: Add input handler to handle these
		// handle_keyup_event(event_type.Keysim.Scancode)
	}
}
