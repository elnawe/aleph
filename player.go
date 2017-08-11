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

func (this *Player) render(texture_manager *Texture_Manager) {
	if this.face == 1 {
		texture_manager.draw_frame(this.Game_Object, sdl.FLIP_HORIZONTAL)
	} else {
		texture_manager.draw_frame(this.Game_Object, sdl.FLIP_NONE)
	}
}

func (this *Player) update() {
	if is_key_down(sdl.SCANCODE_D) {
		this.velocity.set_x(7)
		this.face = 0
		this.is_moving = true
		this.Game_Object.update()
	}

	if is_key_down(sdl.SCANCODE_A) {
		this.velocity.set_x(-7)
		this.face = 1
		this.is_moving = true
		this.Game_Object.update()
	}

	// TODO: Animation
	if this.is_moving {
		if this.current_frame < this.frames-1 {
			this.current_frame++
		} else {
			this.reset_frame_count()
		}
	} else {
		this.reset_frame_count()
	}

	if this.is_grounded {
		if is_key_down(sdl.SCANCODE_W) {
			this.acceleration.set_y(-150)
			this.is_grounded = false
			this.Game_Object.update()
		}
	} else {
		this.acceleration.set_y(0)

		if this.position.y < 600-96 {
			this.acceleration.set_y(5)
		} else {
			this.is_grounded = true
		}

		this.Game_Object.update()
	}

	if key_pressed == nil {
		this.is_moving = false
	}

	this.velocity.reset()
}

func (this *Game_Object) reset_frame_count() {
	this.current_frame = 0
}
