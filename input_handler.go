package main

import "github.com/veandco/go-sdl2/sdl"

type Input_Handler struct {
	key_pressed []uint8
}

var key_pressed []uint8

func is_key_down(key sdl.Scancode) bool {
	if len(key_pressed) != 0 {
		if key_pressed[key] == 1 {
			return true
		}
	}

	return false
}

func (ih *Input_Handler) update() {
	var event sdl.Event
	for event = poll_event; event != nil; event = poll_event() {
		switch event_type := event.(type) {
		case *sdl.QuitEvent:
			// TODO: Add sync_mutex superset
			sync_mutex.Lock()
			should_quit = true
			sync_mutex.Unlock()
		case *sdl.KeyDownEvent:
			ih.on_key_pressed()
			if is_debug_event_enabled() && event != nil {
				debug("Event key pressed:\t%v", "SDL", event_type.Keysym)
			}
		case *sdl.KeyUpEvent:
			ih.on_key_released()
			if is_debug_event_enabled() && event != nil {
				debug("Event key released:\t%v", "SDL", event_type.Keysym)
			}
		case *sdl.MouseMotionEvent:
			if is_debug_event_enabled() && event != nil {
				debug("Event mouse motion:\t%v", "SDL", event_type.X, event_type.Y)
			}
		}
	}
}

func (ih *Input_Handler) on_key_pressed() {
	key_pressed = get_keyboard_state()
}

func (ih *Input_Handler) on_key_released() {
	key_pressed = nil
}
