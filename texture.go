package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Texture_manager struct {
	renderer *sdl.Renderer
}

func (this *Texture_manager) load(filename string) (texture *sdl.Texture) {
	temp_surface := load_image(filename)
	texture = create_texture_from_surface(temp_surface, this.renderer)

	free_surface(temp_surface)

	return
}
