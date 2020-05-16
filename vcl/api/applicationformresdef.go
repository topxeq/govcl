//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------
package api

import (
	"github.com/ying32/dylib"
)

var (
	resFormLoadFromStream       = (*dylib.LazyProc)(nil)
	resFormLoadFromFile         = (*dylib.LazyProc)(nil)
	resFormLoadFromResourceName = (*dylib.LazyProc)(nil)
)

func DoResInit() {
	resFormLoadFromStream = libvcl.NewProc("ResFormLoadFromStream")
	resFormLoadFromFile = libvcl.NewProc("ResFormLoadFromFile")
	resFormLoadFromResourceName = libvcl.NewProc("ResFormLoadFromResourceName")
}

// ResFormLoadFromStream
func ResFormLoadFromStream(obj, root uintptr) {
	resFormLoadFromStream.Call(obj, root)
}

func ResFormLoadFromFile(filename string, root uintptr) {
	resFormLoadFromFile.Call(GoStrToDStr(filename), root)
}

func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	resFormLoadFromResourceName.Call(instance, GoStrToDStr(resName), root)
}
