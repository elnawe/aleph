package main

import "github.com/veandco/go-sdl2/sdl"

type Player struct {
	Game_Object

	is_moving   bool
	is_grounded bool
}

func init() {
	name := "Johnny"
	position := Vector2{0, 600 - 96}
	velocity := Vector2{0, 0}
	acceleration := Vector2{0, 0}
	scale := Vector2{96, 96}
	texture_manager.load("sprite-sheet.png", "sonic")

	the_game.game_objects = append(the_game.game_objects, &Player{
		Game_Object{acceleration, 0, 1, 5, 0, name, position, scale, "sonic", velocity},
		false,
		true,
	})
}

func (p *Player) render(texture_manager *Texture_Manager) {
	if p.face == 1 {
		texture_manager.draw_frame(p.Game_Object, sdl.FLIP_HORIZONTAL)
	} else {
		texture_manager.draw_frame(p.Game_Object, sdl.FLIP_NONE)
	}
}

func (p *Player) update() {
	if is_key_down(sdl.SCANCODE_D) {
		p.velocity.set_x(7)
		p.face = 0
		p.is_moving = true
		p.Game_Object.update()
	}

	if is_key_down(sdl.SCANCODE_A) {
		p.velocity.set_x(-7)
		p.face = 1
		p.is_moving = true
		p.Game_Object.update()
	}

	// TODO: Animation
	if p.is_moving {
		if p.current_frame < p.frames-1 {
			p.current_frame++
		} else {
			p.current_frame = 1
		}
	} else {
		p.current_frame = 0
	}

	if p.is_grounded {
		if is_key_down(sdl.SCANCODE_W) {
			p.acceleration.set_y(-150)
			p.is_grounded = false
			p.Game_Object.update()
		}
	} else {
		p.acceleration.set_y(0)

		if p.position.y < 600-96 {
			p.acceleration.set_y(5)
		} else {
			p.is_grounded = true
		}

		p.Game_Object.update()
	}

	if key_pressed == nil {
		p.is_moving = false
	}

	if !p.is_moving {
		p.velocity.reset()
	}
}
