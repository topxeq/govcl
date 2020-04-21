// 由res2go自动生成。
package main

import (
	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&MainForm)
    vcl.Application.CreateForm(&About)
	vcl.Application.Run()
}
