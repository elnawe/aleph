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

func (g *Game) init(title string, x, y, width, height int, fullscreen bool, texture_manager *Texture_Manager) {
	sdl_init()

	g.window = create_window(title, x, y, width, height, fullscreen)
	g.renderer = create_renderer(g.window, 0)

	g.game_state.change_state(init_menu_state())
}

func (g *Game) render(texture_manager *Texture_Manager) {
	clear_renderer(g.renderer)

	for i, _ := range g.game_objects {
		g.game_objects[i].render(texture_manager)
	}

	present_renderer(g.renderer)
}

func (g *Game) update() {
	for i, _ := range g.game_objects {
		g.game_objects[i].update()
	}

	g.game_state.update()
}

func (g *Game) input() {
	g.input_handler.update()
}
