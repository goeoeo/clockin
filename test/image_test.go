package test

import (
	"fmt"
	"github.com/Comdex/imgo"
	"github.com/phpdi/ant/image"
	"testing"
)

func TestImage(t *testing.T)  {
	imageControl:=new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png","/home/yu/code/clockin/b.png",0,1400,750,150)
}

func TestImage1(t *testing.T)  {
	imageControl:=new(image.ImageControl)
	imageControl.Trimming("/home/yu/code/clockin/data/screen.png","/home/yu/code/clockin/data/login.png",0,900,1500,150)
}

func TestCosineSimilarity(t *testing.T)  {
	cos,err:=imgo.CosineSimilarity("/home/yu/code/clockin/a.png","/home/yu/code/clockin/b.png")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cos)
}