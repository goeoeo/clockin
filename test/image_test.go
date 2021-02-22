package test

import (
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/phpdi/ant/image"
	"github.com/phpdi/clockin/core"
	"log"
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

func TestImage_srceen(t *testing.T) {
	cmd := core.AdbCommand("shell screencap -p | sed 's/\r$//' > /home/yu/code/clockin/data/screen.png")

	if _, err := cmd.CombinedOutput(); err != nil {
		log.Println("err;", err.Error())
	}
}

//进入主页
func TestImage_1(t *testing.T) {
	imageControl := new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png", "/home/yu/code/clockin/data/1.png", 423, 974, 200, 70)
}
