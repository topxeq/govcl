//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build linux

package api

import "github.com/ying32/dylib"

var (
	GdkWindow_GetXId      = (*dylib.LazyProc)(nil)
	GdkWindow_FromForm    = (*dylib.LazyProc)(nil)
	GtkWidget_GetGtkFixed = (*dylib.LazyProc)(nil)
	GtkWidget_Window      = (*dylib.LazyProc)(nil)
)

func InitDefLinux() {
	GdkWindow_GetXId = libvcl.NewProc("GdkWindow_GetXId")
	GdkWindow_FromForm = libvcl.NewProc("GdkWindow_FromForm")
	GtkWidget_GetGtkFixed = libvcl.NewProc("GtkWidget_GetGtkFixed")
	GtkWidget_Window = libvcl.NewProc("GtkWidget_Window")

}
