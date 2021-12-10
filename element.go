package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

// Good use of interface here. Anything that uses the component interface
// needs to have at lease the properties in the component interface
type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type vector struct {
	x, y float64
}

// takes shared information that every element in game will need
type element struct {
	position   vector
	rotation   float64
	active     bool
	components []component
}

// draws component inside element.component
func (elem *element) draw(renderer *sdl.Renderer) error {
	for _, comp := range elem.components {
		err := comp.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

// updates componenet inside element.component
func (elem *element) update() error {
	for _, comp := range elem.components {
		err := comp.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

// add a new component to element
func (elem *element) addComponent(new component) {
	// at compile, loop through every existing component and make sure it's not of the same type as the new component
	for _, existing := range elem.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existing) {
			panic(fmt.Sprintf("attempt to add new component with existing trype %v",
				reflect.TypeOf(new)))
		}
	}
	// add component after check
	elem.components = append(elem.components, new)
}

// easy access component type check for element
func (elem *element) getComponent(withType component) component {
	typ := reflect.TypeOf(withType)
	for _, comp := range elem.components {
		if reflect.TypeOf(comp) == typ {
			return comp
		}
	}
	panic(fmt.Sprintf("no component with type %v", reflect.TypeOf(withType)))
}
