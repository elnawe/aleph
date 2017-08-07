package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

func init_sdl() {
	if sdl.Init(sdl.INIT_EVERYTHING) != nil {
		handle_error(get_sdl_error(), "failed to init SDL: %v", 100)
	}
	debug(fmt.Sprintf("SDL initialized. Goroutine: %d", runtime.NumGoroutine()))
	debug(fmt.Sprintf("Allocated Memory: %v", runtime.GC))
}

func create_window(title string, x, y, w, h int, fullscreen bool) *sdl.Window {
	flags := sdl.WINDOW_SHOWN
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	window, error := sdl.CreateWindow(title, x, y, w, h, uint32(flags))
	handle_error(error, "failed to create window: %v", 101)
	debug(fmt.Sprintf("Window ID: %v has been created", window.GetID()))

	return window
}

func create_renderer(w *sdl.Window, i int) *sdl.Renderer {
	renderer, error := sdl.CreateRenderer(w, i, 0)
	handle_error(error, "failed to create renderer: %v", 102)
	debug("Renderer has been created")
	renderer.SetDrawColor(255, 255, 255, 255)

	return renderer
}

func clear_renderer(r *sdl.Renderer) {
	r.Clear()
}

func present_renderer(r *sdl.Renderer) {
	r.Present()
}

func poll_event() sdl.Event {
	return sdl.PollEvent()
}

func get_sdl_error() error {
	return sdl.GetError()
}

func do_delay_lock_fps(frameTime uint32) {
	if CAP_MAX_FPS && frameTime < DELAY_TIME {
		sdl.Delay(uint32(DELAY_TIME - frameTime))
	}
}

func quit_sdl(w *sdl.Window, r *sdl.Renderer) {
	sdl.Quit()
	debug("exitting SDL")
	if r != nil {
		r.Destroy()
		debug("destroying renderer instace")
	}
	if w != nil {
		w.Destroy()
		debug("destroying window instance")
	}

	os.Exit(0)
}
