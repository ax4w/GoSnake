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
		err := renderer.SetDrawColor(48, 53, 48, 255)
		if err != nil {
			panic(err.Error())
		}
		err = renderer.FillRect(&sdl.Rect{
			X: v.X,
			Y: v.Y,
			W: blockSize,
			H: blockSize,
		},
		)
		if err != nil {
			return
		}
		err = renderer.Copy(
			appleTexture,
			&sdl.Rect{W: blockSize, H: blockSize},
			&sdl.Rect{X: v.X, Y: v.Y, W: blockSize + 5, H: blockSize + 5},
		)
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
