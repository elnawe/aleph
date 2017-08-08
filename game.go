package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Add game_objects
// TODO: Make error to handle game errors
type Game struct {
	currentFrame int32
	errors       []error
	//game_objects  []Game_object
	input_handler Input_Handler
	renderer      *sdl.Renderer
	game_state    Game_State
	window        *sdl.Window
}

func (this *Game) init(title string, x, y, width, height int, fullscreen bool) {
	sdl_init()

	this.window = create_window(title, x, y, width, height, fullscreen)
	this.renderer = create_renderer(this.window, 0)
	texture_manager.renderer = this.renderer

	this.game_state.change_state(init_menu_state())
}

func (this *Game) update() {
	this.game_state.update()
}

func (this *Game) render() {
	clear_renderer(this.renderer)

	// TODO: Add logic to render (draw)

	present_renderer(this.renderer)
}

func (this *Game) handle_inputs() {
	for event := poll_event(); event != nil; event = poll_event() {
		switch event_type := event.(type) {
		case *sdl.QuitEvent:
			// TODO: Add sync_mutex superset
			sync_mutex.Lock()
			should_quit = true
			sync_mutex.Unlock()
		case *sdl.KeyDownEvent:
			if is_debug_event_enabled() && event != nil {
				debug("Event key pressed:\t%v", "SDL", event_type.Keysym)
			}
			// TODO: Add input handler to handle these
			// handle_keydown_event(event_type.Keysym.Scancode)
		case *sdl.KeyUpEvent:
			if is_debug_event_enabled() && event != nil {
				debug("Event key released:\t%v", "SDL", event_type.Keysym)
			}
			// TODO: Add input handler to handle these
			// handle_keyup_event(event_type.Keysim.Scancode)
		case *sdl.MouseMotionEvent:
			if is_debug_event_enabled() && event != nil {
				debug("Event mouse motion:\t%v", "SDL", event_type.X, event_type.Y)
			}
		}
	}
}
