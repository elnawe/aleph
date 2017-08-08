package main

type Game_manager interface {
	update()
	render()

	on_enter() bool
	on_exit() bool

	get_state_id() string
}

type Game_state struct {
	game_managers []Game_manager
}

func (this *Game_state) push_state(state Game_manager) {
	this.game_managers = append(this.game_managers, state)
	this.game_managers[get_last_manager_index(this.game_managers)].on_enter()
}

func (this *Game_state) change_state(state Game_manager) {
	debug("Game managers: %v", "Game_state", this.game_managers)

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

func (this *Game_state) delete_state() {
	if len(this.game_managers) != 0 {
		lastManagerIndex := get_last_manager_index(this.game_managers)

		if this.game_managers[lastManagerIndex].on_exit() {
			temp, states := this.game_managers[lastManagerIndex], this.game_managers[:lastManagerIndex]

			this.game_managers = append(states, temp)
		}
	}
}

func (this *Game_state) update() {
	if len(this.game_managers) != 0 {
		this.game_managers[get_last_manager_index(this.game_managers)].update()
	}
}

func (this *Game_state) render() {
	if len(this.game_managers) != 0 {
		this.game_managers[get_last_manager_index(this.game_managers)].render()
	}
}

func get_last_manager_index(managers []Game_manager) int {
	return len(managers) - 1
}
