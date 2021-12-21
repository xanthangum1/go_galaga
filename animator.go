package main

import "github.com/veandco/go-sdl2/sdl"

type animator struct {
	containter *element
	sequences  map[string]*sequence
}

type sequence struct {
	textures []*sdl.Texture
	frame    int
	// rate of frame change per second
	sampleRate float64
}
