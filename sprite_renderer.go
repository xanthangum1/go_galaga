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
		panic(fmt.Errorf("Querying texture: %v", err))
	}

	return &spriteRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error {
	// convert coordinates to top left of sprite element
	x := sr.container.position.x - float64(sr.width/2.0)
	y := sr.container.position.y - float64(sr.height/2.0)

	// render
	renderer.CopyEx(
		sr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(sr.width), H: int32(sr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(sr.width), H: int32(sr.height)},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.width) / 2, Y: int32(sr.height) / 2},
		sdl.FLIP_NONE,
	)
	return nil
}

func (sr *spriteRenderer) onUpdate() error {
	return nil
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
