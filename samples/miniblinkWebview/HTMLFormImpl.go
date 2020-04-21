// 在这里写你的事件

package main

import "github.com/topxeq/govcl/vcl"

//::private::
type THTMLFormFields struct {
}

func (f *THTMLForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
}
