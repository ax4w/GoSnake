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
	err := renderer.SetDrawColor(0, 0, 0, 255)
	if err != nil {
		panic(err.Error())
	}
	err = renderer.Clear()
	if err != nil {
		panic(err.Error())
	}
}

/*
drawBackground: draws the background image onto the screen
*/
func drawBackground() {
	err := renderer.SetDrawColor(48, 53, 48, 255)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = renderer.FillRect(&sdl.Rect{
		X: 0,
		Y: 0,
		W: windowSize,
		H: windowSize},
	)
	if err != nil {
		panic(err.Error())
	}
	/*err = renderer.Copy(
		backgroundTexture,
		&sdl.Rect{W: 512, H: 512},
		&sdl.Rect{W: 512, H: 512},
	)*/
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
		err := window.Destroy()
		if err != nil {
			//fmt.Println(err.Error())
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
		err := renderer.Destroy()
		if err != nil {
			//fmt.Println(err.Error())
		}
	}(renderer)

	initSnake()
	initApples()

	loadAppleImage()
	defer func(appleTexture *sdl.Texture) {
		err := appleTexture.Destroy()
		if err != nil {
			fmt.Println(err.Error())
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
		t := <-snakeLen
		window.SetTitle(
			fmt.Sprintf("Go,Snake! Score: %d", t-snakeStartLength),
		)
		sdl.Delay(delay)
	}

	return 0
}

func main() {
	os.Exit(run())
}
