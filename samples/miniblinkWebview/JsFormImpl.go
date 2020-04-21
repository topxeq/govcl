// 在这里写你的事件

package main

import "github.com/topxeq/govcl/vcl"

//::private::
type TJsFormFields struct {
}

func (f *TJsForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
}
