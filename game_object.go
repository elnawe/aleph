package main

import "github.com/veandco/go-sdl2/sdl"

type Object_State interface {
	render()
	update()
}

type Game_Object struct {
	Object_State

	id            int
	name          string
	on_update     func(*Game_Object)
	position_rect sdl.Rect
	source_rect   sdl.Rect
	texture       *sdl.Texture
}

func (this *Game_Object) render(r *sdl.Renderer) {
	r.Copy(this.texture, &this.source_rect, &this.position_rect)
}

func (this *Game_Object) update() {
	if this.on_update != nil {
		this.on_update(this)
	}
}

func (this *Game_Object) find_game_object(name string) Game_Object {
	var found_game_object Game_Object

	for _, game_object := range the_game.game_objects {
		if game_object.name == name {
			found_game_object = game_object
		}
	}

	return found_game_object
}

// TODO: Find a better way, probably?
func new_game_object(name string, texture *sdl.Texture, source_rect sdl.Rect) {
	game_object := Game_Object{
		name:        name,
		source_rect: source_rect,
		texture:     texture,
	}

	debug("Creating new Game Object, %v", "Game_Object", game_object)

	the_game.game_objects = append(the_game.game_objects, game_object)
}
