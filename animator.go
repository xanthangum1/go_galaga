package main

import "github.com/veandco/go-sdl2/sdl"

type animator struct {
	container *element
	sequences map[string]*sequence
	current   string
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator

	an.sequences = sequences
	an.current = defaultSequence
	an.container = container

	return &an
}

func (an *animator) onDraw(renderer *sdl.Renderer) error {
	tex := an.sequences[an.current].texture()

	return drawTexture(
		tex,
		an.container.position,
		an.container.rotation,
		renderer,
	)
}

type sequence struct {
	textures []*sdl.Texture
	frame    int
	// rate of frame change per second
	sampleRate float64
}

func (seq *sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}
