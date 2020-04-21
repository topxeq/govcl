// +build windows

package main

import (
	"os"

	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/rtl"
	"github.com/topxeq/govcl/vcl/win"
)

func main() {
	rtl.CreateURLShortCut(win.GetDesktopPath(), "govcl", "https://github.com/ying32/govcl")

	rtl.CreateShortCut(win.GetDesktopPath(), "shortcut", os.Args[0], "", "描述", "-b -c")

	vcl.ShowMessage("Hello!")

}
