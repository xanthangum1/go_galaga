package main

import "github.com/veandco/go-sdl2/sdl"

type keyboardMover struct {
	container *element
	speed     float64
	sr        *spriteRenderer
}

func newKeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&spriteRenderer{}).(*spriteRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// User cant move off screen left
		if cont.position.x-(playerSize/2.0) > 0 {
			// Move player left
			cont.position.x -= playerSpeed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// User cant move off screen right
		if cont.position.x+(playerSize/2.0) < screenWidth {
			// move player right
			cont.position.x += playerSpeed
		}
	}
	return nil
}
