package main

import (
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

func sdl_init() {
	runtime.LockOSThread()
	if sdl.Init(sdl.INIT_EVERYTHING) != nil {
		handle_error(sdl_get_error(), "failed to init SDL: %v", 100)
	}
	debug("SDL initialized. Goroutine: %d", runtime.NumGoroutine())
	debug("Allocated Memory: %v", runtime.GC)
}

func create_window(title string, x, y, w, h int, fullscreen bool) *sdl.Window {
	flags := sdl.WINDOW_OPENGL
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	window, error := sdl.CreateWindow(title, x, y, w, h, uint32(flags))
	handle_error(error, "failed to create window: %v", 101)
	debug("Window ID: %v has been created", window.GetID())

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

func poll_event() (event sdl.Event) {
	event = sdl.PollEvent()
	if is_debug_event_enabled() && event != nil {
		debug("Event triggered: %v", event)
	}

	return
}

func sdl_get_error() error {
	return sdl.GetError()
}

func sdl_get_ticks() uint32 {
	return sdl.GetTicks()
}

func do_delay_lock_fps(frameTime uint32) {
	if CAP_MAX_FPS && frameTime < DELAY_TIME {
		sdl.Delay(uint32(DELAY_TIME - frameTime))
	}
}

func sdl_quit(w *sdl.Window, r *sdl.Renderer) {
	sdl.Quit()
	debug("exitting SDL")
	if r != nil {
		debug("destroying renderer instace")
		r.Destroy()
	}
	if w != nil {
		debug("destroying window ID: %d instance", w.GetID())
		w.Destroy()
	}
	os.Exit(0)
}
