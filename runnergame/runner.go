package runnergame

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

const (
	screenWidth  = 640
	screenHeight = 420

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameCount  = 8

	jumpHeight = 40
)

var (
	runnerImage     *ebiten.Image
	runnerOffset    float64 = 10
	jumpFlag        bool    = false
	centerPosition          = float64(screenHeight / 2)
	runnerPosition          = centerPosition
	runnerYMovement         = -float64(1.5)
)

type Game struct {
	count int
	// pressedKeys []ebiten.Key
}

func Init() *Game {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal("Error loading runner image", err)
	}
	runnerImage = ebiten.NewImageFromImage(img)
	return &Game{}
}

func (g *Game) GetHeight() int {
	return screenHeight
}

func (g *Game) GetWidth() int {
	return screenWidth
}

func (g *Game) Update() error {
	g.count++

	if ebiten.IsKeyPressed(ebiten.KeySpace) ||
		ebiten.IsKeyPressed(ebiten.KeyArrowUp) ||
		ebiten.IsKeyPressed(ebiten.KeyW) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) ||
		ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) &&
			!jumpFlag {
		jumpFlag = true
	}

	if jumpFlag {
		g.runnerJump()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawRunner(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) runnerJump() {

	runnerPosition += runnerYMovement

	fmt.Println(runnerPosition, runnerYMovement, centerPosition)

	if runnerPosition > centerPosition {
		runnerYMovement *= -1
		runnerPosition = centerPosition
		jumpFlag = false
	}

	if runnerPosition <= (centerPosition - jumpHeight) {
		runnerYMovement *= -1
	}
}

func (g *Game) drawRunner(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	// op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(runnerOffset, runnerPosition)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
}
