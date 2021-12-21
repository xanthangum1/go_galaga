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
	tex := textureFromBMP(renderer, filename)

	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
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

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	// load object from bmp file
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		//panic out if error
		panic(fmt.Errorf("loading %v: %v", filename, err))
	}
	// prevent memory leak
	defer img.Free()
	// use previously loaded bmp to create texture
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		panic(fmt.Errorf("creating texture from %v: %v", filename, err))
	}
	return tex
}

// apply collisions
func (sr *spriteRenderer) onCollision(other *element) error {
	return nil
}
