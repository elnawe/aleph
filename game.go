package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// TODO: Make error to handle game errors
type Game struct {
	currentFrame  int32
	errors        []error
	game_objects  []Game_Object_State
	input_handler Input_Handler
	renderer      *sdl.Renderer
	game_state    Game_State
	window        *sdl.Window
}

func (this *Game) init(title string, x, y, width, height int, fullscreen bool, texture_manager *Texture_Manager) {
	sdl_init()

	this.window = create_window(title, x, y, width, height, fullscreen)
	this.renderer = create_renderer(this.window, 0)

	this.game_state.change_state(init_menu_state())
}

func (this *Game) render(texture_manager *Texture_Manager) {
	clear_renderer(this.renderer)

	for i, _ := range this.game_objects {
		this.game_objects[i].render(texture_manager)
	}

	present_renderer(this.renderer)
}

func (this *Game) update() {
	for i, _ := range this.game_objects {
		this.game_objects[i].update()
	}

	this.game_state.update()
}

func (this *Game) input() {
	this.input_handler.update()
}
