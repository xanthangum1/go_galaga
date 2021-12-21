package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func drawTexture(
	tex *sdl.Texture,
	position vector,
	rotation float64,
	renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}
	// convert coordinates to top left of sprite element
	position.x -= float64(width) - 2.0
	position.y -= float64(height) - 2.0

	// render
	return renderer.CopyEx(
		tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)},
		&sdl.Rect{X: int32(position.x), Y: int32(position.y), W: int32(width), H: int32(height)},
		rotation,
		&sdl.Point{X: int32(width) / 2, Y: int32(height) / 2},
		sdl.FLIP_NONE)
}

func loadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	// load object from bmp file
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		//panic out if error
		return nil, fmt.Errorf("loading %v: %v", filename, err)
	}
	// prevent memory leak
	defer img.Free()
	// use previously loaded bmp to create texture
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		//return empyty player if error to bubble error up the call stack
		return nil, fmt.Errorf("creating texture from %v: %v", filename, err)
	}
	return tex, nil
}
