package main

import (
	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
	_ "github.com/topxeq/govcl/vcl/locales/zh_CN"
	"github.com/topxeq/govcl/vcl/types"
	"github.com/topxeq/govcl/vcl/win"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	vcl.ShowMessage("消息")
	if vcl.MessageDlg("消息", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes {
		vcl.ShowMessage("你点击了“是")
	}
	if vcl.Application.MessageBox("消息", "标题", win.MB_OKCANCEL+win.MB_ICONINFORMATION) == types.MrOk {
		vcl.ShowMessage("你点击了“是")
	}

	vcl.Application.Run()
}
