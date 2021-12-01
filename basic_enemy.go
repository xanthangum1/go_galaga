package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

type basicEnemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newBasicEnemy(renderer *sdl.Renderer, x float64, y float64) (be basicEnemy, err error) {
	// load player spaceship for use
	img, err := sdl.LoadBMP("sprites/basic_enemy.bmp")
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		return basicEnemy{}, fmt.Errorf("loading enemy sprite: %v", err)
	}
	// prevent memory leak
	defer img.Free()
	// use previously loaded spaceship to create texture
	be.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		return basicEnemy{}, fmt.Errorf("creating enemy texture: %v", err)
	}

	be.x = x
	be.y = y

	return be, nil
}

func (be *basicEnemy) draw(renderer *sdl.Renderer) {
	// convert enemy coordinates to top left of sprite
	x := be.x - basicEnemySize/2.0
	y := be.y - basicEnemySize/2.0

	// Create enemy spaceship in game
	renderer.Copy(
		be.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
	)
}
