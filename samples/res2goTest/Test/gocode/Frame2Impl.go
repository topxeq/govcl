// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"github.com/topxeq/govcl/vcl"
)

//::private::
type TFrame2Fields struct {
}

func (f *TFrame2) OnButton1Click(sender vcl.IObject) {
	vcl.ShowMessage("Frame2")
}