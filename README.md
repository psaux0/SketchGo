SketchGo
========

An Encapsulation of Go Image Library with various sketch algorithms.

###Example
```
package main

import (
	sg "sketchgo"
	"image/color"
)

func main() {
	//Edge length of the image
	//Circle center, radius
	const (
		edgeX = 255
		edgeY = 255
		cX = 80
		cY = 80
		cRadius = 40
		)
	
	img := sg.NewImageToDraw(edgeX,edgeY)
	img.SetBackgroundColor(color.RGBA{0,255,0,255})
	img.Line(0,0,edgeX,edgeY,color.RGBA{0,0,0,255})

	//Filling color is blue
	img.Fill(color.RGBA{0,0,255,255})
	img.Circle(cX,cY,cRadius,color.RGBA{255,0,0,255})
	//Clear
	img.FillClear()

	img.Fill(color.RGBA{255,125,255,255})
	img.Rectangle(150,150,200,230,color.RGBA{255,0,0,255})
	img.FillClear()

	//Generate an image file
	//jpeg is not a good choice
	img.GenerateImgFile("testpic","png")
}
```

###Output
![](https://lh6.googleusercontent.com/-DU-m3-3TqVw/VEnXKLzcmOI/AAAAAAAAAK4/h8Ez2JlwjNo/s255-no/dgo.png)