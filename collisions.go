package main

import "math"

type circle struct {
	center vector
	radius float64
}

func collides(c1, c2 circle) bool {
	// Use euclidian distance to detect collision
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) + math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= math.Max(c1.radius+c2.radius)
}

func checkCollisions() error {
	// checking for combination for every pair in elements
	for i := 0; i < len(elements)-1; i++ {
		for j := i + 1; j < len(elements); j++ {
			for _, c1 := range elements[i].collisions {
				for _, c2 := range elements[j].collisions {
					if collides(c1, c2) && elements[i].active && elements[j].active {
						err := elements[i].collision(elements[j])
						if err != nil {
							return err
						}
						err = elements[j].collision(elements[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}
	return nil
}
