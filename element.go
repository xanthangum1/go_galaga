package main

import "github.com/veandco/go-sdl2/sdl"

// Good use of interface here. Anything that uses the component interface
// needs to have at lease the properties in the component interface
type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type vector struct {
	x, y float64
}

type element struct {
	position   vector
	rotation   float64
	active     bool
	components []component
}

func (elem *element) addComponent(new component) {

}
