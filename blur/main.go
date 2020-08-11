package main

import (
	"bytes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"image"
	_ "image/jpeg"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	gophersImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(gophersImage, op)

	// Box blur (7x7)
	// https://en.wikipedia.org/wiki/Box_blur
	for j := -3; j <= 3; j++ {
		for i := -3; i <= 3; i++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i), 244+float64(j))
			// Alpha scale should be 1.0/49.0, but accumulating 1/49 49 times doesn't reach to 1, because
			// the final color is affected by the destination alpha when CompositeModeSourceOver is used.
			// This composite mode is the default mode. See how this is calculated at the doc:
			// https://pkg.go.dev/github.com/hajimehoshi/ebiten#CompositeMode
			//
			// Use a higher value than 1.0/49.0. Here, 1.0/25.0 here to get a reasonable result.
			op.ColorM.Scale(1, 1, 1, 1.0/25.0)
			screen.DrawImage(gophersImage, op)
		}
	}

	return nil
}

func main() {
	// 读取图片
	img, _, err := image.Decode(bytes.NewReader(images.FiveYears_jpg))
	if err != nil {
		log.Fatal(err)
	}

	gophersImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Blur (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}

}
