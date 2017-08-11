package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Texture_Manager struct {
	error       error
	game        *Game
	texture_map map[string]*sdl.Texture
	zero_point  *sdl.Point
}

func (this *Texture_Manager) load(file_name, id string) {
	temp_surface := load_image(file_name)

	this.texture_map[id] = create_texture_from_surface(temp_surface, this.game.renderer)

	defer temp_surface.Free()
}

func (this *Texture_Manager) draw(game_object Game_Object, flip sdl.RendererFlip) {
	var source_rect, destination_rect sdl.Rect

	source_rect.X, source_rect.Y = 0, 0
	source_rect.W, destination_rect.W = int32(game_object.scale.x), int32(game_object.scale.y)
	source_rect.H, destination_rect.H = int32(game_object.scale.y), int32(game_object.scale.y)
	destination_rect.X, destination_rect.Y = int32(game_object.position.x), int32(game_object.position.y)

	this.game.renderer.CopyEx(
		this.texture_map[game_object.texture_id],
		&source_rect,
		&destination_rect,
		0,
		this.zero_point,
		flip,
	)
}

func (this *Texture_Manager) draw_frame(game_object Game_Object, flip sdl.RendererFlip) {
	var source_rect, destination_rect sdl.Rect

	source_rect.X = int32(game_object.scale.x) * game_object.current_frame
	source_rect.Y = int32(game_object.scale.y) * (game_object.current_row - 1)
	destination_rect.X, destination_rect.Y = int32(game_object.position.x), int32(game_object.position.y)
	source_rect.W, source_rect.H = int32(game_object.scale.x), int32(game_object.scale.y)
	destination_rect.W, destination_rect.H = int32(game_object.scale.x), int32(game_object.scale.y)

	this.game.renderer.CopyEx(
		this.texture_map[game_object.texture_id],
		&source_rect,
		&destination_rect,
		0,
		this.zero_point,
		flip,
	)
}
