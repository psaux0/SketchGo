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
	img.Circle(cX,cY,cRadius,color.RGBA{255,0,0,255})
	img.Rectangle(30,30,60,90,color.RGBA{255,0,0,255})

	//jpeg is not a good choice
	img.GenerateImgFile("testpic","png")
}
```

###Output
![](https://lh3.googleusercontent.com/-CFBgTcr8Hjs/VEh5BLF-ExI/AAAAAAAAAKM/2rqCXsgHX28/s255-no/testpic.png)