package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 4
	playerSize  = 105

	playerShotCooldown = time.Millisecond * 250
)

// Creating a player spaceship class
type player struct {
	tex  *sdl.Texture
	x, y float64

	// used for shot cooldown
	lastShot time.Time
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
		&sdl.Rect{X: 0, Y: 0, W: playerSize, H: playerSize},
		&sdl.Rect{X: int32(x), Y: int32(y), W: playerSize, H: playerSize},
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
		if time.Since(p.lastShot) >= playerShotCooldown {
			if bul, ok := bulletFromPool(); ok {
				bul.active = true
				bul.x = p.x
				bul.y = p.y
				bul.angle = 270 * (math.Pi / 180)

				p.lastShot = time.Now()
			}
		}
	}
}
