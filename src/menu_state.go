package main

type Menu_State struct {
	Game_Manager
	menu_id string
}

func init_menu_state() *Menu_State {
	return &Menu_State{menu_id: "MENU"}
}

func (this *Menu_State) update() {}
func (this *Menu_State) render() {}

func (this *Menu_State) on_enter() bool {
	debug("Entering Menu_state: %s", "Menu_state", this.menu_id)
	return true
}

func (this *Menu_State) on_exit() bool {
	debug("Exiting Menu_state: %s", "Menu_state", this.menu_id)
	return true
}

func (this *Menu_State) get_state_id() string {
	return this.menu_id
}
