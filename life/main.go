package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"math/rand"
	"time"
)

type World struct {
	area   []bool
	width  int
	height int
}

func NewWorld(width, height int, maxInitLiveCells int) *World {
	w := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	w.init(maxInitLiveCells)
	return w
}

func (w *World) init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
		// 随便生成亮点
		x := rand.Intn(w.width)
		y := rand.Intn(w.height)
		w.area[y*w.width+x] = true
	}
}

// Update game state by one tick.
func (w *World) Update() {
	width := w.width
	height := w.height
	next := make([]bool, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			pop := neighbourCount(w.area, width, height, x, y)
			switch {
			case pop < 2:
				// rule 1. Any live cell with fewer than two live neighbours
				// dies, as if caused by under-population.
				next[y*width+x] = false

			case (pop == 2 || pop == 3) && w.area[y*width+x]:
				// rule 2. Any live cell with two or three live neighbours
				// lives on to the next generation.
				next[y*width+x] = true

				/*case pop > 3:
					// rule 3. Any live cell with more than three live neighbours
					// dies, as if by over-population.
					next[y*width+x] = false

				case pop == 3:
					// rule 4. Any dead cell with exactly three live neighbours
					// becomes a live cell, as if by reproduction.
					next[y*width+x] = true*/
			}
		}
	}
	w.area = next
}

// Draw paints current game state.
func (w *World) Draw(pix []byte) {
	for i, v := range w.area {
		if v {
			// 一个像素点用四个字节表示RGBA
			pix[4*i] = 0xff
			pix[4*i+1] = 0xff
			pix[4*i+2] = 0xff
			pix[4*i+3] = 0xff
		} else {
			pix[4*i] = 0
			pix[4*i+1] = 0
			pix[4*i+2] = 0
			pix[4*i+3] = 0
		}
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// neighbourCount calculates the Moore neighborhood of (x, y).
func neighbourCount(a []bool, width, height, x, y int) int {
	c := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || width <= x2 || height <= y2 {
				continue
			}
			if a[y2*width+x2] {
				c++
			}
		}
	}
	return c
}

func init() {
	rand.Seed(time.Now().UnixNano())
	world = NewWorld(screenWidth, screenHeight, int((screenWidth*screenHeight)/10))
}

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	world  *World
	pixels = make([]byte, screenWidth*screenHeight*4)
)

func update(screen *ebiten.Image) error {
	world.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	world.Draw(pixels)
	// 替换像素点的显示效果
	screen.ReplacePixels(pixels)
	return nil
}

// 替换图片的某个像素点
func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Game of Life (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
