//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"fmt"

	. "github.com/topxeq/govcl/vcl/win"
)

func showError(err interface{}) {
	MessageBox(0, fmt.Sprint(err), "Error", MB_ICONERROR)
}

// CN: 尝试加载默认Application icon
// EN: Try to load the default Application icon.
func tryLoadAppIcon() {
	hIcon := LoadIcon2(GetSelfModuleHandle(), "MAINICON")
	if hIcon != 0 {
		Application.Icon().SetHandle(hIcon)
	}
}
