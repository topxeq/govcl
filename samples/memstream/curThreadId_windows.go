package main

import "github.com/topxeq/govcl/vcl/win"

func GetCurrentThreadId() uintptr {
	return win.GetCurrentThreadId()
}
