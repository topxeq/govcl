// 由res2go自动生成。
package main

import (
	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/rtl"
)

const Lazarus = true

func main() {
	if rtl.LcLLoaded() {
		vcl.Application.SetFormScaled(true)
	}
	vcl.Application.Initialize()
    vcl.Application.CreateForm(&MainForm)
    vcl.Application.CreateForm(&About)
	vcl.Application.Run()
}
