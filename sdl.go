package main

import (
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func sdl_init() {
	runtime.LockOSThread()
	if sdl.Init(sdl.INIT_EVERYTHING) != nil {
		handle_error(sdl_get_error(), "failed to init SDL: %v", "SDL", 100)
	}
	debug("initialized. Goroutine: %d", "SDL", runtime.NumGoroutine())
	debug("Allocated Memory: %v", "SDL", runtime.GC)
}

func create_window(title string, x, y, w, h int, fullscreen bool) (window *sdl.Window) {
	flags := sdl.WINDOW_OPENGL
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}
	window, error := sdl.CreateWindow(title, x, y, w, h, uint32(flags))
	handle_error(error, "failed to create window: %v", "SDL", 101)
	debug("Window ID: %v has been created", "SDL", window.GetID())

	return
}

func create_renderer(w *sdl.Window, i int) (renderer *sdl.Renderer) {
	renderer, error := sdl.CreateRenderer(w, i, 0)
	handle_error(error, "failed to create renderer: %v", "SDL", 102)
	debug("Renderer has been created", "SDL")
	renderer.SetDrawColor(255, 255, 255, 255)

	return
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

func load_image(f string) (s *sdl.Surface) {
	s, e := img.Load(f)
	handle_error(e, "failed to load image: %v", "SDL", 200)

	return
}

func free_surface(s *sdl.Surface) {
	s.Free()
}

func create_texture_from_surface(s *sdl.Surface, r *sdl.Renderer) (t *sdl.Texture) {
	t, e := r.CreateTextureFromSurface(s)
	handle_error(e, "failed to create texture from surface: %v", "SDL", 201)

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
	runtime.UnlockOSThread()
	debug("exitting", "SDL")
	if r != nil {
		debug("destroying renderer instace", "SDL")
		r.Destroy()
	}
	if w != nil {
		debug("destroying window ID: %d instance", "SDL", w.GetID())
		w.Destroy()
	}
	os.Exit(0)
}
