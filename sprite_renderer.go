package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	container *element
	tex       *sdl.Texture
}

func newSpriteRenderer(container *element, renderer *sdl.Renderer, filename string) *spriteRenderer {
	return spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	_, _, width, height, err := sr.tex.Query()
	if err != nil {
		return fmt.Errorf("Querying texture: %v", err)
	}
	// convert coordinates to top left of sprite element
	x := sr.container.position.x - float64(width/2.0)
	y := sr.container.position.y - float64(height/2.0)
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
	// use previously loaded spaceship to create texture
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		panic(fmt.Errorf("creating texture %v: %v", filename, err))
	}
	return tex
}
