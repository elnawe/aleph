package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func load_texture(filename string, r *sdl.Renderer) (texture *sdl.Texture) {
	temp_surface := load_image(filename)

	if r != nil {
		texture = create_texture_from_surface(temp_surface, r)
	} else {
		texture = create_texture_from_surface(temp_surface, the_game.renderer)
	}

	free_surface(temp_surface)

	return
}
