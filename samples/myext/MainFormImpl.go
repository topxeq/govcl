// 在这里写你的事件

package main

import (
	"github.com/topxeq/govcl/samples/myext/msrdp"
	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/types"
)

//::private::
type TMainFormFields struct {
	rdp *msrdp.TMsRdpClient9NotSafeForScripting
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()

	f.rdp = msrdp.NewMsRdpClient9NotSafeForScripting(f)
	f.rdp.SetParent(f)
	f.rdp.SetAlign(types.AlClient)
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *TMainForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {

}
