package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/chrsan/shadow"
)

func main() {
	ctx := shadow.NewContext()
	defer func() {
		if err := ctx.Dispose(); err != nil {
			log.Fatal(err)
		}
	}()
	s := ctx.NewRGBASurface(640, 360)
	defer s.Dispose()
	c := s.Canvas
	c.Clear(255, 255, 255, 255)
	p, d := ctx.StackAllocPaint()
	defer d()
	p.SetStyle(shadow.PaintStyleFill)
	p.SetColor(255, 0, 0, 255)
	p.SetAntiAlias(true)
	c.DrawRect(100, 100, 75, 50, p)
	p.SetStyle(shadow.PaintStyleStroke)
	p.SetStrokeCap(shadow.StrokeCapSquare)
	p.SetStrokeJoin(shadow.StrokeJoinRound)
	p.SetStrokeWidth(3)
	p.SetColor(0, 0, 255, 255)
	c.DrawRect(100, 200, 75, 50, p)
	c.Flush()
	im := image.RGBA{
		Pix:    s.Data(true),
		Stride: 4 * 640,
		Rect:   image.Rect(0, 0, 640, 360),
	}
	f, err := os.Create("sample.png")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error creating file:", err)
		return
	}
	defer f.Close()
	err = png.Encode(f, &im)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error encoding PNG:", err)
	}
}
