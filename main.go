package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 900
)

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	// load object from bmp file
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		//panic out if error
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	// prevent memory leak
	defer img.Free()
	// use previously loaded spaceship to create texture
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		panic(fmt.Errorf("creating texture %v: %v", filename, err))
	}
	return tex
}

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

	// render new player
	plr, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("creating player:", err)
		return
	}

	var enemies []basicEnemy

	// render enemy troupe
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/5)*screenWidth + (basicEnemySize / 2.0)
			y := float64(j)*basicEnemySize + (basicEnemySize / 2.0)

			enemy, err := newBasicEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("creating basic enemy:", err)
				return
			}

			enemies = append(enemies, enemy)
		}
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

		// render all enemies
		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		// shows everything on renderer
		renderer.Present()
	}
}
