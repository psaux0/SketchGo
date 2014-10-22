package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

//in order to define new method for non-local struct
type imgLocal struct {
	image.NRGBA
}

func (img *imgLocal) SetBackgroundColor(clr color.RGBA) {
	for x := img.Rect.Min.X; x <= img.Rect.Max.X; x++ {
		for y := img.Rect.Min.Y; y <= img.Rect.Max.Y; y++ {
			img.Set(x,y,clr)
		}
	}
}


//Draw Line by Bresenham's Algorithm, Cannot Draw Vertical Line, Better to use HorizentalLine to draw Horizental lines
func (img *imgLocal) Line(x0, y0, x1, y1 int, clr color.RGBA) {
	dx := x1 - x0
	dy := y1 - y0

	D := 2*dy - dx
	img.Set(x0,y0,clr)
	y := y0

	for x := x0 + 1; x <= x1; x++ {
		if D > 0 {
			y++
			img.Set(x,y,clr)
			D += 2*(dy - dx)
		} else {
			img.Set(x,y,clr)
			D += 2*dy
		}
	}
}

//Draw VerticalLine
func (img *imgLocal) VerticalLine(x, y0, y1 int, clr color.RGBA) {
	for y := y0; y <= y1; y++ {
		img.Set(x,y,clr)
	}
}

//Draw HorizentalLine
func (img *imgLocal) HorizentalLine(y, x0, x1 int, clr color.RGBA) {
	for x := x0; x <= x1; x++ {
		img.Set(x,y,clr)
	}
}


//Draw Circle by Bresenham's Algorithm
func (img *imgLocal) Circle(x0, y0, radius int, clr color.RGBA) {
	x := radius
	y := 0
	radiusError := 1 - x
	for x >= y {
		img.Set(x + x0, y + y0, clr)
		img.Set(y + x0, x + y0, clr)
		img.Set(-x + x0, y + y0, clr)
		img.Set(-y + x0, x + y0, clr)	
		img.Set(-x + x0, -y + y0, clr)
		img.Set(-y + x0, -x + y0, clr)
		img.Set(x + x0, -y + y0, clr)
		img.Set(y + x0, -x + y0, clr)
		y++;
		if (radiusError < 0) {
			radiusError += 2 * y + 1
		} else {
			x--
			radiusError += 2 * (y - x + 1)
		}
	}
}

func (img *imgLocal) Rectangel(x0, y0, x1, y1 int, clr color.RGBA) {
	img.HorizentalLine(y0,x0,x1,clr)
	img.HorizentalLine(y1,x0,x1,clr)
	img.VerticalLine(x0,y0,y1,clr)
	img.VerticalLine(x1,y0,y1,clr)
}

//Initiate a img data structure
func NewImageToDraw(eX, eY int) (img *imgLocal){
	ig := image.NewNRGBA(image.Rect(0,0,eX,eY))
	return &imgLocal{*ig}
}

func main() {
	//Edge length of the image
	//Circle center, radius
	const (
		edgeX = 255
		edgeY = 255
		cX = 10
		cY = 10
		cRadius = 40
		)
	//define local struct to implement method
	img := NewImageToDraw(edgeX,edgeY)
	img.SetBackgroundColor(color.RGBA{0,255,0,255})
	img.Line(0,0,edgeX,edgeY,color.RGBA{0,0,0,255})
	img.Circle(cX,cY,cRadius,color.RGBA{255,0,0,255})
	img.Rectangel(30,30,60,90,color.RGBA{255,0,0,255})

	imgFile, _ := os.Create("line.png")
	defer imgFile.Close()

	//local stuct inherents *image.NRGBA's method
	err := png.Encode(imgFile,img)
	if err != nil {
		log.Fatal(err)
	}
}