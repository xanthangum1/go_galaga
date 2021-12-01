package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x float64, y float64) (be basicEnemy) {
	// create enemy texture
	be.tex = textureFromBMP(renderer, "sprites/basic_enemy.bmp")

	// set default location coordinates
	be.x = x
	be.y = y

	return be
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	// convert enemy coordinates to top left of sprite
	x := be.x - basicEnemySize/2.0
	y := be.y - basicEnemySize/2.0

	// Create enemy spaceship in game
	renderer.CopyEx(
		be.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
		180,
		&sdl.Point{X: basicEnemySize / 2, Y: basicEnemySize / 2},
		sdl.FLIP_NONE,
	)
}
