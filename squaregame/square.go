package squaregame

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 420
	squareSize   = 100
)

var (
	squarePositionX = float64(screenWidth) / 2
	squarePositionY = float64(screenHeight) / 2
	squareMovementX = float64(5)
	squareMovementY = float64(5)

	rgbaColor = color.RGBA{255, 0, 255, 255}
)

type Game struct{}

func (g *Game) GetHeight() int {
	return screenHeight
}

func (g *Game) GetWidth() int {
	return screenWidth
}

func (g *Game) Update() error {
	squarePositionX += squareMovementX
	squarePositionY += squareMovementY
	if squarePositionX >= screenWidth-squareSize || squarePositionX <= 0 {
		squareMovementX *= -1
		rgbaColor = g.getRandomColor()
	}
	if squarePositionY >= screenHeight-squareSize || squarePositionY <= 0 {
		squareMovementY *= -1
		rgbaColor = g.getRandomColor()
	}
	return nil
}

func (g *Game) getRandomColor() color.RGBA {

	colorList := []color.RGBA{
		{255, 0, 255, 255},   // purple
		{255, 0, 0, 255},     //red
		{255, 165, 0, 255},   //orange
		{255, 255, 0, 255},   //yellow
		{0, 255, 0, 255},     //green
		{0, 0, 255, 255},     //blue
		{75, 0, 130, 255},    //indigo
		{238, 130, 238, 255}, // violet
	}

	// pick random

	randomIndex := rand.Intn(len(colorList))

	return colorList[randomIndex]
}

func (g *Game) drawSquare(screen *ebiten.Image, rgbaColor color.RGBA, squarPosX int, squarePosY int) {

	fmt.Println("Drawing square with pos", squarPosX, squarePosY)
	for y := squarePosY; y < squarePosY+squareSize; y++ {
		// fmt.Println("First or last line", y)
		if y == squarePosY || y == (squarePosY+squareSize-1) {
			for x := squarPosX; x < squarPosX+squareSize; x++ {
				screen.Set(x, y, rgbaColor)
			}
		} else {
			// fmt.Println("Middle lines", squarePositionX, y)
			screen.Set(squarPosX, y, rgbaColor)
			screen.Set(squarPosX+squareSize-1, y, rgbaColor)
		}

	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrintAt(screen, "Hello, World!", 30, 150)

	x1 := int(math.Round(squarePositionX))
	y1 := int(math.Round(squarePositionY))
	fmt.Println("Drawing with val", x1, y1)
	g.drawSquare(screen, rgbaColor, x1, y1)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
