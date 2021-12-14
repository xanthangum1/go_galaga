package main

import (
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

	// render from bmp file
	sr := newSpriteRenderer(basicEnemy, renderer, "sprites/basic_enemy.bmp")
	// Add basic enemy components to element
	basicEnemy.addComponent(sr)

	col := circle{
		center: basicEnemy.position,
		radius: 38,
	}
	basicEnemy.collisions = append(basicEnemy.collisions, col)

	basicEnemy.active = true

	return basicEnemy
}
