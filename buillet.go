package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const bulletSize = 32
const bulletSpeed = 20

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	sr := newSpriteRenderer(bullet, renderer, "sprites/bullet.bmp")
	bullet.addComponent(sr)
	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	//fill up bulletPool with bullets
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	// comb through bulletPool to find bullet not in use
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
