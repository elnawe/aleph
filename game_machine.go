package main

type Game_Manager interface {
	update()
	render()

	on_enter() bool
	on_exit() bool

	get_state_id() string
}

type Game_State struct {
	game_managers []Game_Manager
}

func (this *Game_State) push_state(state Game_Manager) {
	this.game_managers = append(this.game_managers, state)
	this.game_managers[get_last_manager_index(this.game_managers)].on_enter()
}

func (this *Game_State) change_state(state Game_Manager) {
	debug("Game Managers: %v", "Game_State", this.game_managers)

	if len(this.game_managers) != 0 {
		lastManagerIndex := get_last_manager_index(this.game_managers)

		if this.game_managers[lastManagerIndex].get_state_id() == state.get_state_id() {
			return
		}

		if this.game_managers[lastManagerIndex].on_exit() {
			temp, states := this.game_managers[lastManagerIndex], this.game_managers[:lastManagerIndex]

			this.game_managers = append(states, temp)
		}
	}
}

func (this *Game_State) delete_state() {
	if len(this.game_managers) != 0 {
		lastManagerIndex := get_last_manager_index(this.game_managers)

		if this.game_managers[lastManagerIndex].on_exit() {
			temp, states := this.game_managers[lastManagerIndex], this.game_managers[:lastManagerIndex]

			this.game_managers = append(states, temp)
		}
	}
}

func (this *Game_State) update() {
	if len(this.game_managers) != 0 {
		this.game_managers[get_last_manager_index(this.game_managers)].update()
	}
}

func (this *Game_State) render() {
	if len(this.game_managers) != 0 {
		this.game_managers[get_last_manager_index(this.game_managers)].render()
	}
}

func get_last_manager_index(managers []Game_Manager) int {
	return len(managers) - 1
}
