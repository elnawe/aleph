package main

type Game_Console struct {
	Game_Object

	enabled bool
	height  int
}

func (g *Game_Console) render(texture_manager *Texture_Manager) {
	if g.enabled {
		//texture_manager.draw(g, sdl.FLIP_NONE)
	}
}

func (g *Game_Console) update() {

}
