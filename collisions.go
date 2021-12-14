package main

import "math"

type circle struct {
	center vector
	radius float64
}

func collides(c1, c2 circle) bool {
	// Use euclidian distance to detect collision
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) + math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= c1.radius+c2.radius
}
