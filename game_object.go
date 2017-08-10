package main

import "github.com/veandco/go-sdl2/sdl"

// NOTE: This is a super early version of Game_Object.
// Game_Object should have everything to make it work.

type Game_Object_State interface {
	render() // render(g *Game, texture_manager *Texture_Manager). Make it to pass game and texture_manager
	update()
}

type Game_Object struct {
	id           int32
	name         string
	acceleration Vector2
	position     Vector2
	velocity     Vector2
}

func (this *Game_Object) render(r *sdl.Renderer) {
	// Refactor with texture manager
	r.Copy(this.texture, &this.source_rect, &this.position_rect)
}

func (this *Game_Object) update() {
	this.velocity.add(this.acceleration)
	this.position.add(this.velocity)
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
