package main

import (
	"testing"

	"fmt"

	"github.com/topxeq/govcl/vcl"
)

func TestPSDReader(t *testing.T) {
	bmp := vcl.NewBitmap()
	defer bmp.Free()
	if err := PsdToBitmap("J:\\Delphi\\PsSources\\spotify.psd", bmp); err == nil {
		bmp.SaveToFile("test.bmp")
	} else {
		fmt.Println("err:", err)
	}
}
