package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Make error to handle game errors
type Game struct {
	currentFrame  int32
	errors        []error
	game_objects  []Game_Object_State
	input_handler Input_Handler
	renderer      *sdl.Renderer
	game_state    Game_State
	window        *sdl.Window
}

func (this *Game) init(title string, x, y, width, height int, fullscreen bool, texture_manager *Texture_Manager) {
	sdl_init()

	this.window = create_window(title, x, y, width, height, fullscreen)
	this.renderer = create_renderer(this.window, 0)

	this.game_state.change_state(init_menu_state())
}

func (this *Game) render(texture_manager *Texture_Manager) {
	clear_renderer(this.renderer)

	for i, _ := range this.game_objects {
		this.game_objects[i].render(this, texture_manager)
	}

	present_renderer(this.renderer)
}

func (this *Game) update() {
	for i, _ := range this.game_objects {
		this.game_objects[i].update()
	}

	this.game_state.update()
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
