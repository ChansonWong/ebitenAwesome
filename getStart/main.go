package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"log"
	"strconv"
)

type Game struct{}

var num = 0

var img *ebiten.Image

func (g *Game) Update(screen *ebiten.Image) error {
	num++
	return nil
}

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("source/gopher.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(1.5, 1)
	screen.DrawImage(img, op)
	ebitenutil.DebugPrint(screen, "Hello, World! "+strconv.Itoa(num))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
