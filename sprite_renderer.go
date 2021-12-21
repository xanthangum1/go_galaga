package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container *element
	tex       *sdl.Texture
	width     float64
	height    float64
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	sr := spriteRenderer{}
	var err error

	sr.tex, err = loadTextureFromBMP(filename, renderer)
	if err != nil {
		panic(err)
	}

	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       sr.tex,
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	return drawTexture(sr.tex, sr.container.position, sr.container.rotation, renderer)
}

// apply collisions
func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
