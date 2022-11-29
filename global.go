package main

import "github.com/veandco/go-sdl2/sdl"

// all used global const
const (
	windowSize       = 512
	blockSize        = 20
	applePath        = "./assets/apple.png"
	snakeStartLength = 6
	appleCount       = 5
	north            = 0
	east             = 1
	south            = 2
	west             = 3
)

// all used global vars
var (
	snake        = make([]*snakeChunk, 0, windowSize*windowSize)
	renderer     *sdl.Renderer
	window       *sdl.Window
	appleTexture *sdl.Texture
	running      = true
	apples       = make([]*apple, 0, appleCount)
	delay        = uint32(55)
)

// all used global types
type (
	snakeChunk struct {
		X, Y, direction int32
	}
	apple struct {
		X, Y int32
	}
)
