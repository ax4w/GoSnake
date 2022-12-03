package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"os"
	"time"
)

func loadAppleImage() {
	image, _ := img.Load(applePath)
	defer image.Free()
	appleTexture, _ = renderer.CreateTextureFromSurface(image)
}
func clearRenderer() {
	if renderer.SetDrawColor(0, 0, 0, 255) != nil {
		panic("Error in main, in clearRenderer while calling setDrawColor")
	}
	if renderer.Clear() != nil {
		panic("Error in main, in clearRenderer while calling Clear")
	}
}

/*
drawBackground: draws the background image onto the screen
*/
func drawBackground() {
	if renderer.SetDrawColor(48, 53, 48, 255) != nil {
		panic("Error in main, in drawBackground while calling setDrawColor")
	}
	if renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: windowSize,
		H: windowSize},
	) != nil {
		panic("Error in main, in drawBackground while calling FillRect")
	}
}

func draw(snakeLen chan int) {
	for running {
		if hasSnakeHitBounds() || hasSnakeHitItSelf() {
			close(snakeLen)
			running = false
			break
		}
		handleAppleHit()
		clearRenderer()
		drawBackground()
		renderApples(renderer)
		renderSnake(renderer)
		moveSnake()
		renderer.Present()
		snakeLen <- len(snake)
		sdl.Delay(delay)
	}

}

func run() int {
	rand.Seed(time.Now().UnixNano())
	/*
		Create Window
	*/
	window, err := sdl.CreateWindow(
		"Go, Snake!",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		windowSize,
		windowSize,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err.Error())
	}

	defer func(window *sdl.Window) {
		if window.Destroy() != nil {
			//panic("Error while destroying window")
		}
	}(window)

	/*
		Create Renderer
	*/
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err.Error())
	}

	defer func(renderer *sdl.Renderer) {
		if renderer.Destroy() != nil {
			panic("Error while destroying renderer")
		}
	}(renderer)

	initSnake()
	initApples()

	loadAppleImage()
	defer func(appleTexture *sdl.Texture) {
		if appleTexture.Destroy() != nil {
			panic("Error while destroying apple texture")
		}
	}(appleTexture)

	snakeLen := make(chan int)
	go draw(snakeLen)

	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				switch keyCode {
				case 'a', 's', 'd', 'w':
					changeSnakeDirection(keyCode)
				}

			case *sdl.QuitEvent:
				running = false
			}
		}
		select {
		case t := <-snakeLen:
			window.SetTitle(
				fmt.Sprintf("Go,Snake! Score: %d", t-snakeStartLength),
			)
		default:
			continue
		}
	}

	return 0
}

func main() {
	os.Exit(run())
}
