package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	tex  *sdl.Texture
	x, y float64
}

// create a new bullet
func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = textureFromBMP(renderer, "sprites/player_bullet.bmp")

	return bul
}
