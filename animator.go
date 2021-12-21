package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type animator struct {
	container       *element
	sequences       map[string]*sequence
	current         string
	lastFrameChange time.Time
}

func newAnimator(container *element, sequences map[string]*sequence, defaultSequence string) *animator {
	var an animator

	an.sequences = sequences
	an.current = defaultSequence
	an.container = container
	an.lastFrameChange = time.Now()

	return &an
}

func (an *animator) onUpdate() error {
	sequence := an.sequences[an.current]
	frameInterval := float64(time.Second) / sequence.sampleRate
	if time.Since(an.lastFrameChange) >= time.Duration(frameInterval) {
		sequence.nextFrame()
		an.lastFrameChange = time.Now()
	}

	return nil

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
	loop       bool
}

func newSequence(filepath string, sampleRate float64, loop bool, renderer *sdl.Renderer) (*sequence, error) {
	var seq sequence

	files, err := ioutil.ReadDir(filepath)
	if err != nil {
		return nil, fmt.Errorf("reading directory %v: %v", filepath, err)
	}

	for _, file := range files {
		filename := path.Join(filepath, file.Name())
		tex, err := loadTextureFromBMP(filename, renderer)
		if err != nil {
			return nil, fmt.Errorf("loading sequence frame: %v", err)
		}
		seq.textures = append(seq.textures, tex)
	}

	seq.sampleRate = sampleRate
	seq.loop = loop

	return &seq, nil
}

func (seq *sequence) texture() *sdl.Texture {
	return seq.textures[seq.frame]
}

func (seq *sequence) nextFrame() {
	if seq.frame == len(seq.textures)-1 {
		if seq.loop {
			seq.frame = 0
		}
	} else {
		seq.frame++
	}
}
