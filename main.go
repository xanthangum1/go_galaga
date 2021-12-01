package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 1200
)

func main() {
	// initializing all tools in sdl --standard
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}
	// create a window to play our game
	window, err := sdl.CreateWindow(
		"Gaming in Go",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	// catch error while creating window
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	// prevent memory leak
	defer window.Destroy()

	// create rendering object for our units
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	// catch error for renderer
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	// prevent memory leak after we are done
	defer renderer.Destroy()

	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	enemy, err := newBasicEnemy(renderer, screenWidth/2.0, screenHeight/2.0)
	if err != nil {
		fmt.Println("Initializing enemy:", err)
		return
	}

	for {
		// for loop to catch user exiting with alt + f4
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		// background color
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		// render player spaceship
		plr.draw(renderer)

		// update player position
		plr.update()

		enemy.draw(renderer)

		// shows everything on renderer
		renderer.Present()
	}
}
