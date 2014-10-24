/*
	A simple program encapsulated go/image library
	Build drawing algortihms for it.
*/
package sketchgo

import (
	"image"
	"image/color"
	"image/png"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"fmt"
)

//in order to define new method for non-local struct, extend the struct too
//can also be used for image.NRGBA functions because those functions only care about its elements
//if it is defined as struct {ig image.NRGBA}. this is not a inherence, then it cannot be used
//in those functions
type imgLocal struct {
	image.NRGBA
	toFill bool
	fillColor color.RGBA
}

type fillColor struct {
	toFill bool
	colorFill color.RGBA
}

type pixel struct {
	X int
	Y int
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
	if img.toFill {
		floodFill(x0,y0,clr,img.fillColor,img)
	}
}

func (img *imgLocal) Fill(clr color.RGBA) {
	img.toFill = true
	img.fillColor = clr
}

func (img *imgLocal) FillClear() {
	img.toFill = false
	img.fillColor = color.RGBA{255,255,255,255}
}

func (img *imgLocal) Rectangle(x0, y0, x1, y1 int, clr color.RGBA) {
	img.HorizentalLine(y0,x0,x1,clr)
	img.HorizentalLine(y1,x0,x1,clr)
	img.VerticalLine(x0,y0,y1,clr)
	img.VerticalLine(x1,y0,y1,clr)
	if img.toFill {
		midX := (x0 + x1) >> 1
		midY := (y0 + y1) >> 1
		floodFill(midX, midY, clr, img.fillColor, img)
	}
}

//Initiate a img data structure
func NewImageToDraw(eX, eY int) (img *imgLocal){
	ig := image.NewNRGBA(image.Rect(0,0,eX,eY))
	return &imgLocal{
		*ig,
		true,
		color.RGBA{255,255,255,255},
	}
}

//Create a File
func (img *imgLocal) GenerateImgFile(name string, imgFormat string) {
	imgFormat = strings.ToLower(imgFormat)

	imgFile, _ := os.Create(name + "." + imgFormat)
	defer imgFile.Close()

	//local stuct inherents *image.NRGBA's method
	switch {
	case imgFormat == "jpeg":
		if err := jpeg.Encode(imgFile,img,nil); err != nil {
			log.Fatal(err)
		}
	case imgFormat == "png":
		if err := png.Encode(imgFile,img); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Encode Format Error, Please use jpeg or png")
		return
	}
}

//If two types of color equal to each other
func colorEqual(cx color.NRGBA, cy color.RGBA) bool {
	if cx.R == cy.R && cx.G == cy.G && cx.B == cy.B && cx.A == cy.A {
		return true
	} else {
		return false
	}
}


//FloodFill ALgorithm fill in color
//Bug is two shapes are overlapped and share same color edge, only part of the shape will be dyed
func floodFill(x, y int, edgeColor color.RGBA, fColor color.RGBA, img *imgLocal) {
	dx := []int{1,0,-1,0}
	dy := []int{0,1,0,-1}
	var deque *Deque = NewDeque()
	deque.Append(pixel{x,y})
	img.Set(x,y,fColor)
	for !deque.Empty() {
		p := deque.First().(pixel)
		deque.Shift()
		for i := 0; i < 4; i++ {
			nx := p.X + dx[i]
			ny := p.Y + dy[i]
			pColor := img.At(nx,ny).(color.NRGBA)
			if !(colorEqual(pColor, edgeColor) || colorEqual(pColor,fColor)){
				deque.Append(pixel{nx,ny})
				img.Set(nx,ny,fColor)
			}
		}
	}
}

