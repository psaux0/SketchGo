package sketchgo

import (
	"testing"
	"os"
)

func TestNewImageToDraw(t *testing.T) {
	ig := NewImageToDraw(40,40)
	if ig == nil {
		t.Log("Cannot Create New Image")
		t.Fail()
	}
}

func TestGenerateImgFile(t *testing.T) {
	img := NewImageToDraw(50,50)
	img.GenerateImgFile("hello", "png")
	defer os.Remove("hello.png")

	if _, err := os.Stat("hello.png"); os.IsNotExist(err) {
		t.Log("Encoding Data Error")
		t.Fail()
	}

	img.GenerateImgFile("hello", "jpeg")
	defer os.Remove("hello.jpeg")

	if _, err := os.Stat("hello.jpeg"); os.IsNotExist(err) {
		t.Log("Encoding Data Error")
		t.Fail()
	}

}