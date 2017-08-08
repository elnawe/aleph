package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func load_texture(filename string, renderer *sdl.Renderer) (texture *sdl.Texture) {
	temp_surface := load_image(filename)
	texture = create_texture_from_surface(temp_surface, renderer)

	free_surface(temp_surface)

	return
}
