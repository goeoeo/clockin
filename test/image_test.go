package test

import (
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/phpdi/ant/image"
	"testing"
)

func TestImage(t *testing.T) {
	imageControl := new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png", "/home/yu/code/clockin/data/bluetooth.png", 400, 1400, 350, 750)
}

func TestImage_end(t *testing.T) {
	imageControl := new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png", "/home/yu/code/clockin/data/bluetooth.png", 400, 1400, 350, 750)
}

func TestImage_start(t *testing.T) {
	imageControl := new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png", "/home/yu/code/clockin/data/bluetooth_start.png", 400, 1300, 350, 350)
}

func TestCosineSimilarity(t *testing.T) {
	cos, err := imgo.CosineSimilarity("/home/yu/code/clockin/data/bluetooth.png", "/home/yu/code/clockin/data/bluetooth_tmp.png")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cos)
}
