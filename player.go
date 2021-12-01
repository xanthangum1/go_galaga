package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 4
	playerSize  = 105
)

// Creating a player spaceship class
type player struct {
	tex  *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	// create player texture
	p.tex = textureFromBMP(renderer, "sprites/player.bmp")

	// set default position of player
	p.x = screenWidth / 2.0
	p.y = screenHeight - playerSize*2.5

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	// convert player coordinates to top left of sprite
	x := p.x - playerSize/2.0
	y := p.y - playerSize/2.0

	// Create player spaceship in game
	renderer.Copy(
		p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
	)
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// Move player left
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move player right
		p.x += playerSpeed
	}

	if keys[sdl.SCANCODE_UP] == 1 {
		// Move player up
		p.y -= playerSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		// move player down
		p.y += playerSpeed
	}

	// listen for shooting bullets
	if keys[sdl.SCANCODE_SPACE] == 1 {
		if bul, ok := bulletFromPool(); ok {
			bul.active = true
			bul.x = p.x
			bul.y = p.y
		}
	}
}
