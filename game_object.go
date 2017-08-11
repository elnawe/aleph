package main

import "github.com/veandco/go-sdl2/sdl"

// NOTE: This is a super early version of Game_Object.
// Game_Object should have everything to make it work.

type Game_Object_State interface {
	render(*Texture_Manager)
	update()
}

type Game_Object struct {
	acceleration  Vector2
	current_frame int32
	current_row   int32
	frames        int32
	face          int
	name          string
	position      Vector2
	scale         Vector2
	texture_id    string
	velocity      Vector2
}

func (this *Game_Object) render(texture_manager *Texture_Manager) {
	texture_manager.draw(*this, sdl.FLIP_NONE)
}

func (this *Game_Object) update() {
	this.velocity.add(this.acceleration)
	this.position.add(this.velocity)
}
