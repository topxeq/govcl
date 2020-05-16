//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
	_ "github.com/topxeq/govcl/vcl/locales/zh_CN"
)

func main() {
	vcl.RunApp(&MainForm, &NewConnection)
}
