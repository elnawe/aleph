package main

import "github.com/veandco/go-sdl2/sdl"

// NOTE: This is a super early version of Game_Object.
// Game_Object should have everything to make it work.

type Game_Object_State interface {
	render(*Game, *Texture_Manager)
	update()
}

type Game_Object struct {
	acceleration  Vector2
	current_frame int32
	current_row   int32
	unique_id     int32
	name          string
	position      Vector2
	scale         Vector2
	texture_id    string
	velocity      Vector2
}

func (this *Game_Object) render(game *Game, texture_manager *Texture_Manager) {
	texture_manager.draw(*this, game, sdl.FLIP_NONE)
}

func (this *Game_Object) update() {
	this.velocity.add(this.acceleration)
	this.position.add(this.velocity)
}

func instantiate(name, texture_id string, width, height, x, y float64) *Game_Object {
	return &Game_Object{
		name:       name,
		position:   Vector2{x, y},
		scale:      Vector2{width, height},
		texture_id: texture_id,
	}
}
