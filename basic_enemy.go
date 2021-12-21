package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const basicEnemySize = 105

// creates a basic enemy
func newBasicEnemy(renderer *sdl.Renderer, position vector) *element {
	basicEnemy := &element{}

	// establishes position
	basicEnemy.position = position
	// establishes roation
	basicEnemy.rotation = 180

	idleSequence, err := newSequence("sprites/basic_enemy/idle", 10, true, renderer)
	if err != nil {
		panic(fmt.Errorf("creating idle sequence: %v %v", error))
	}
	destroySequence, err := newSequence("sprites/basic_enemy/destroy", 10, false, renderer)
	if err != nil {
		panic(fmt.Errorf("creating destroy sequence: %v %v", error))
	}

	sequences := map[string]*sequence{
		"idle":    idleSequence,
		"destroy": destroySequence,
	}

	animator := newAnimator(basicEnemy)
	basicEnemy.addComponent(animator)

	// render from bmp file
	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	// Add basic enemy components to element
	basicEnemy.addComponent(sr)

	vtb := newVulnerableToBullets(basicEnemy)
	basicEnemy.addComponent(vtb)

	col := circle{
		center: basicEnemy.position,
		radius: 38,
	}
	basicEnemy.collisions = append(basicEnemy.collisions, col)

	basicEnemy.active = true

	return basicEnemy
}
