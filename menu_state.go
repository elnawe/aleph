package main

import (
	"fmt"
)

type Menu_state struct {
	Game_manager
	menu_id string
}

func init_menu_state() *Menu_state {
	return &Menu_state{menu_id: "MENU"}
}

func (this *Menu_state) update() {}
func (this *Menu_state) render() {}

func (this *Menu_state) on_enter() bool {
	debug(fmt.Sprintf("Entering Menu_state: %s", this.menu_id))
	return true
}

func (this *Menu_state) on_exit() bool {
	debug(fmt.Sprintf("Exiting Menu_state: %s", this.menu_id))
	return true
}

func (this *Menu_state) get_state_id() string {
	return this.menu_id
}
