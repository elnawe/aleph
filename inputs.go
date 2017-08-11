package main

import "github.com/veandco/go-sdl2/sdl"

type Input_Handler struct{}

func is_key_down(key sdl.Scancode) bool {
	if len(key_pressed) != 0 {
		if key_pressed[key] == 1 {
			return true
		}
	}

	return false
}
