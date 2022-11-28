package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func head() *snakeChunk {
	return snake[0]
}

func hasSnakeHitBounds() bool {
	return head().X < 0 ||
		head().X >= windowSize ||
		head().Y < -5 || //IDK why -5 tbh
		head().Y >= windowSize
}

func hasSnakeHitItSelf() bool {
	for i, v := range snake {
		if head().X == v.X && head().Y == v.Y && i != 0 {
			return true
		}
	}
	return false
}

func handleAppleHit() {
	for _, v := range apples {
		if head().X >= v.X-blockSize && head().X <= v.X+blockSize &&
			head().Y >= v.Y-blockSize/2 && head().Y <= v.Y+blockSize {
			snake = append(snake, &snakeChunk{
				X: snake[len(snake)-1].X,
				Y: snake[len(snake)-1].Y,
			})
			x := randomPosition()
			y := randomPosition()
			for areApplesOverlapping(x, y) || areApplesOverlappingSnakeBody(x, y) {
				x = randomPosition()
				y = randomPosition()
			}
			for _, k := range apples {
				if k.X == v.X && k.Y == v.Y {
					k.X = x
					k.Y = y
				}
			}
		}
	}
}

func renderSnake(renderer *sdl.Renderer) {

	for i, v := range snake {
		red := uint8(67)
		green := uint8(198)
		blue := uint8(68)
		if i%3 == 0 {
			red = uint8(50)
			green = uint8(149)
			blue = uint8(50)
		}
		r := new(sdl.Rect)
		r.X = v.X
		r.Y = v.Y
		r.W = blockSize
		r.H = blockSize

		err := renderer.SetDrawColor(red, green, blue, 255)
		if err != nil {
			return
		}
		err = renderer.FillRect(r)
		if err != nil {
			return
		}
	}
}

func moveSnake() {
	var tmpX, tmpY int32
	prevX := head().X
	prevY := head().Y
	for i, v := range snake {
		//head is extra
		if i == 0 {
			switch head().direction {
			case north:
				head().Y -= blockSize / 2
			case south:
				head().Y += blockSize / 2
			case east:
				head().X += blockSize / 2
			case west:
				head().X -= blockSize / 2
			}
			continue
		}
		tmpX = v.X
		tmpY = v.Y
		v.X = prevX
		v.Y = prevY
		prevX = tmpX
		prevY = tmpY
	}
}

func changeSnakeDirection(key sdl.Keycode) {
	switch key {
	case 'a':
		if head().direction == north || head().direction == south {
			head().direction = west
		}
	case 'd':
		if head().direction == north || head().direction == south {
			head().direction = east
		}
	case 'w':
		if head().direction == east || head().direction == west {
			head().direction = north
		}
	case 's':
		if head().direction == east || head().direction == west {
			head().direction = south
		}
	}
}

func initSnake() {
	for i := 0; i < snakeStartLength; i++ {
		/*
			Build the head
		*/
		if i == 0 {
			snake = append(snake, &snakeChunk{
				X:         snakeStartLength * blockSize,
				Y:         windowSize / 2,
				direction: east,
			})
			continue
		}
		snake = append(snake, &snakeChunk{
			X: head().X - (int32(i) * blockSize),
			Y: head().Y,
		})
	}

}
