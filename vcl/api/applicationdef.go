//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
// "fmt"
// "github.com/topxeq/tk"
)

func Application_Instance() uintptr {
	ret, _, _ := application_Instance.Call()
	return ret
}

func Application_CreateForm(app uintptr, initScale bool) uintptr {
	ret, _, _ := application_CreateForm.Call(app, GoBoolToDBool(initScale))
	return ret
}

func Application_Run(app uintptr) {
	// tk.Pl("tid: %#v", tk.GetCurrentThreadID())
	// fmt.Println(application_Run)
	// fmt.Println(app)
	// a, b, c := application_Run.Call(app)
	// fmt.Printf("run: %#v, %#v, %#v\n", a, b, c)
	// 运行完后free下
	application_Run.Call(app)
	closeLib()
}

func Application_Initialize(obj uintptr) {
	application_Initialize.Call(obj)
}
