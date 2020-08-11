package main

import (
	"bytes"
	"ebitenAwesome/source"
	"github.com/hajimehoshi/ebiten"
	"image"
	_ "image/png"
	"log"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32 // 每个人仔的大小（宽）
	frameHeight = 32 // 每个人仔的大小（长）
	frameNum    = 8
)

var (
	count       = 0
	runnerImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	count++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	// 设置图片的位置
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	// 5是速度控制
	i := (count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY
	// 图片切割
	screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)
	return nil
}

func main() {
	// 读取图片
	img, _, err := image.Decode(bytes.NewReader(source.Runner_png))
	if err != nil {
		log.Fatal(err)
	}

	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Animation (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}

}
