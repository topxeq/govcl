package main

import "github.com/topxeq/govcl/vcl"
import _ "github.com/topxeq/govcl/pkgs/winappres"

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm, true)
	vcl.Application.Run()
}
