package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

func randomPosition() int32 {
	return int32(rand.Intn(((windowSize - blockSize) - blockSize) + blockSize))
}

func areApplesOverlappingSnakeBody(x, y int32) bool {
	for _, v := range snake {
		if x >= v.X-blockSize && x <= v.X+blockSize &&
			y >= v.Y-blockSize && y <= v.Y+blockSize {
			return true
		}
	}
	return false
}

func areApplesOverlapping(x, y int32) bool {
	for _, v := range apples {
		if x >= v.X-blockSize && x <= v.X+blockSize &&
			y >= v.Y-blockSize && y <= v.Y+blockSize {
			return true
		}
	}
	return false
}

func renderApples(renderer *sdl.Renderer) {
	for _, v := range apples {
		if renderer.SetDrawColor(48, 53, 48, 255) != nil {
			panic("Error in apple, in renderApples while calling setDrawColor")
		}
		if renderer.FillRect(&sdl.Rect{
			X: v.X,
			Y: v.Y,
			W: blockSize,
			H: blockSize},
		) != nil {
			panic("Error in apple, in renderApples while calling FillRect")
		}
		if renderer.Copy(
			appleTexture,
			&sdl.Rect{W: blockSize, H: blockSize},
			&sdl.Rect{X: v.X, Y: v.Y, W: blockSize + 5, H: blockSize + 5},
		) != nil {
			panic("Error in apple, in renderApples while calling Copy")
		}
	}
}

func initApples() {
	for i := 0; i < appleCount; i++ {
		x := randomPosition()
		y := randomPosition()
		for areApplesOverlapping(x, y) || areApplesOverlappingSnakeBody(x, y) {
			x = randomPosition()
			y = randomPosition()
		}
		apples = append(apples, &apple{
			X: x,
			Y: y,
		})
	}
}
