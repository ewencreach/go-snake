package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
)

const (
	screenWidth     = 1920
	screenHeight    = 1080
	snakeSquareSize = 20
	tps             = 30
)

var (
	snakeImage     *ebiten.Image
	snakePosX      float64
	snakePosY      float64
	snakeDirection ebiten.Key
	foodPosX       float64
	foodPosY       float64
	score          int
	positions      map[int]struct{ X, Y float64 }
)

type Game struct{}

func (g *Game) Update() error {
	if score > 0 {
		for i := score; i > 0; i-- {
			positions[i] = positions[i-1]
		}
	}

	movement()

	if score > 0 {
		for i := score; i > 0; i-- {
			if positions[i].X == positions[0].X && positions[i].Y == positions[0].Y {
				start()
			}
		}
	}

	if snakePosY == foodPosY && snakePosX == foodPosX {
		score++

		foodPosX = float64(rand.Intn(screenWidth/snakeSquareSize) * snakeSquareSize)
		foodPosY = float64(rand.Intn(screenHeight/snakeSquareSize) * snakeSquareSize)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", score))

	for i := 0; i <= score; i++ {
		snakeOptions := &ebiten.DrawImageOptions{}
		snakeOptions.GeoM.Translate(positions[i].X, positions[i].Y)
		screen.DrawImage(snakeImage, snakeOptions)
	}

	foodImage := ebiten.NewImage(snakeSquareSize, snakeSquareSize)
	foodImage.Fill(color.RGBA{R: 255, A: 255})
	foodOptions := &ebiten.DrawImageOptions{}
	foodOptions.GeoM.Translate(foodPosX, foodPosY)
	screen.DrawImage(foodImage, foodOptions)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	start()

	ebiten.SetTPS(tps)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}

func start() {
	snakeImage = ebiten.NewImage(snakeSquareSize, snakeSquareSize)
	snakeImage.Fill(color.White)

	snakePosX = float64(rand.Intn(screenWidth/snakeSquareSize) * snakeSquareSize)
	snakePosY = float64(rand.Intn(screenHeight/snakeSquareSize) * snakeSquareSize)

	snakeDirection = ebiten.KeyArrowUp

	foodPosX = float64(rand.Intn(screenWidth/snakeSquareSize) * snakeSquareSize)
	foodPosY = float64(rand.Intn(screenHeight/snakeSquareSize) * snakeSquareSize)

	score = 0
	positions = make(map[int]struct{ X, Y float64 })
}

func movement() {
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || snakeDirection == ebiten.KeyArrowUp) && snakeDirection != ebiten.KeyArrowDown {
		snakePosY -= snakeSquareSize
		if snakePosY < 0 {
			snakePosY = screenHeight
		}
		snakeDirection = ebiten.KeyArrowUp
	}

	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || snakeDirection == ebiten.KeyArrowDown) && snakeDirection != ebiten.KeyArrowUp {
		snakePosY += snakeSquareSize
		if snakePosY > screenHeight {
			snakePosY = 0
		}
		snakeDirection = ebiten.KeyArrowDown
	}

	if (ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || snakeDirection == ebiten.KeyArrowLeft) && snakeDirection != ebiten.KeyArrowRight {
		snakePosX -= snakeSquareSize
		if snakePosX < 0 {
			snakePosX = screenWidth
		}
		snakeDirection = ebiten.KeyArrowLeft
	}

	if (ebiten.IsKeyPressed(ebiten.KeyArrowRight) || snakeDirection == ebiten.KeyArrowRight) && snakeDirection != ebiten.KeyArrowLeft {
		snakePosX += snakeSquareSize
		if snakePosX > screenWidth {
			snakePosX = 0
		}
		snakeDirection = ebiten.KeyArrowRight
	}

	positions[0] = struct{ X, Y float64 }{X: snakePosX, Y: snakePosY}
}
