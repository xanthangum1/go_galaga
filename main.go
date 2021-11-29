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
	// catch error for rendere
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	// prevent memory leak after we are done
	defer renderer.Destroy()

	// load player spaceship for use
	img, err := sdl.LoadBMP("sprites/player.bmp")
	if err != nil {
		fmt.Println("loading player sprite:", err)
		return
	}
	// prevent memory leak
	defer img.Free()
	// use previously loaded spaceship to create texture
	playerTexture, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("creating player texture:", err)
	}
	// prevent memory leak
	defer playerTexture.Destroy()

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

		// Create player spaceship in game
		renderer.Copy(
			playerTexture,
			&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
			&sdl.Rect{X: 40, Y: 20, W: 105, H: 105},
		)

		// shows everything on renderer
		renderer.Present()
	}
}
