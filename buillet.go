package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const bulletSize = 32

type bullet struct {
	tex  *sdl.Texture
	x, y float64

	active bool
}

// create a new bullet
func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/player_bullet.bmp")

	return bul
}

func (bul *bullet) draw(renderer *sdl.Renderer) {
	// if bullet is not active, bullet is not drawn
	if !bul.active {
		return
	}
	// set bullet location reference to center of bullet object
	x := bul.x - bulletSize/2.0
	y := bul.y - bulletSize/2.0

	// render bullet
	renderer.Copy(
		bul.tex,
		&sdl.Rect{X: 0, Y: 0, W: 105, H: 105},
		&sdl.Rect{X: int32(x), Y: int32(y), W: 105, H: 105},
	)
}

var bulletPool []*bullet

func initBulletPool(renderer *sdl.Renderer) {
	//fill up bulletPool with bullets
	for i := 0; i > 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, &bul)
	}
}

func bulletFromPool() (*bullet, bool) {
	// comb through bulletPool to find bullet not in use
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
