//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/ying32/dylib"
)

var (
	application_Instance   = (*dylib.LazyProc)(nil)
	application_CreateForm = (*dylib.LazyProc)(nil)
	application_Run        = (*dylib.LazyProc)(nil)
	application_Initialize = (*dylib.LazyProc)(nil)

	form_Create2                = (*dylib.LazyProc)(nil)
	form_EnabledMaximize        = (*dylib.LazyProc)(nil)
	form_EnabledMinimize        = (*dylib.LazyProc)(nil)
	form_EnabledSystemMenu      = (*dylib.LazyProc)(nil)
	form_SetAllowDropFiles      = (*dylib.LazyProc)(nil)
	form_GetAllowDropFiles      = (*dylib.LazyProc)(nil)
	form_SetOnDropFiles         = (*dylib.LazyProc)(nil)
	form_SetOnDestroy           = (*dylib.LazyProc)(nil)
	form_SetOnConstrainedResize = (*dylib.LazyProc)(nil)
	form_SetOnDeactivate        = (*dylib.LazyProc)(nil)
	form_SetOnActivate          = (*dylib.LazyProc)(nil)
	form_SetOnStyleChanged      = (*dylib.LazyProc)(nil)
	form_SetOnWndProc           = (*dylib.LazyProc)(nil)
	form_SetShowInTaskBar       = (*dylib.LazyProc)(nil)
	form_ShowInTaskBar          = (*dylib.LazyProc)(nil)
	form_ScaleForCurrentDpi     = (*dylib.LazyProc)(nil)
	form_InheritedWndProc       = (*dylib.LazyProc)(nil)

	setEventCallback      = (*dylib.LazyProc)(nil)
	setMessageCallback    = (*dylib.LazyProc)(nil)
	setThreadSyncCallback = (*dylib.LazyProc)(nil)

	dGetStringArrOf = (*dylib.LazyProc)(nil)
	dStrLen         = (*dylib.LazyProc)(nil)
	dMove           = (*dylib.LazyProc)(nil)

	dShowMessage = (*dylib.LazyProc)(nil)
	dMessageDlg  = (*dylib.LazyProc)(nil)

	mouse_Instance  = (*dylib.LazyProc)(nil)
	screen_Instance = (*dylib.LazyProc)(nil)

	dSynchronize = (*dylib.LazyProc)(nil)

	// TMenuItem
	dTextToShortCut = (*dylib.LazyProc)(nil)
	dShortCutToText = (*dylib.LazyProc)(nil)

	// TClipboard
	clipboard_Instance     = (*dylib.LazyProc)(nil)
	clipboard_SetClipboard = (*dylib.LazyProc)(nil)

	// DSysOpen
	dSysOpen = (*dylib.LazyProc)(nil)

	// TMemoryStream
	memoryStream_Read  = (*dylib.LazyProc)(nil)
	memoryStream_Write = (*dylib.LazyProc)(nil)

	// TCanvas
	canvas_BrushCopy     = (*dylib.LazyProc)(nil)
	canvas_CopyRect      = (*dylib.LazyProc)(nil)
	canvas_Draw1         = (*dylib.LazyProc)(nil)
	canvas_Draw2         = (*dylib.LazyProc)(nil)
	canvas_DrawFocusRect = (*dylib.LazyProc)(nil)
	canvas_FillRect      = (*dylib.LazyProc)(nil)
	canvas_FrameRect     = (*dylib.LazyProc)(nil)
	canvas_StretchDraw   = (*dylib.LazyProc)(nil)
	canvas_TextRect1     = (*dylib.LazyProc)(nil)
	canvas_TextRect2     = (*dylib.LazyProc)(nil)
	canvas_Polygon       = (*dylib.LazyProc)(nil)
	canvas_Polyline      = (*dylib.LazyProc)(nil)
	canvas_PolyBezier    = (*dylib.LazyProc)(nil)
	canvas_PolyBezierTo  = (*dylib.LazyProc)(nil)
	canvas_Pixels        = (*dylib.LazyProc)(nil)
	canvas_SetPixels     = (*dylib.LazyProc)(nil)

	// TImageList
	imageList_Draw1        = (*dylib.LazyProc)(nil)
	imageList_Draw2        = (*dylib.LazyProc)(nil)
	imageList_DrawOverlay1 = (*dylib.LazyProc)(nil)
	imageList_DrawOverlay2 = (*dylib.LazyProc)(nil)
	imageList_GetIcon1     = (*dylib.LazyProc)(nil)
	imageList_GetIcon2     = (*dylib.LazyProc)(nil)

	// TBitmap
	bitmap_Clear          = (*dylib.LazyProc)(nil)
	bitmap_BeginUpdate    = (*dylib.LazyProc)(nil)
	bitmap_EndUpdate      = (*dylib.LazyProc)(nil)
	bitmap_LoadFromDevice = (*dylib.LazyProc)(nil)

	dExtractFilePath = (*dylib.LazyProc)(nil)
	dFileExists      = (*dylib.LazyProc)(nil)

	dSelectDirectory1 = (*dylib.LazyProc)(nil)
	dSelectDirectory2 = (*dylib.LazyProc)(nil)
	dInputBox         = (*dylib.LazyProc)(nil)
	dInputQuery       = (*dylib.LazyProc)(nil)

	// TForm相关设置
	setGlobalFormScaled      = (*dylib.LazyProc)(nil)
	form_ScaleForPPI         = (*dylib.LazyProc)(nil)
	form_ScaleControlsForDpi = (*dylib.LazyProc)(nil)

	// TSysLocaled
	dSysLocale = (*dylib.LazyProc)(nil)

	// Shortcut
	dCreateURLShortCut = (*dylib.LazyProc)(nil)
	dCreateShortCut    = (*dylib.LazyProc)(nil)

	// SetProperty
	dSetPropertyValue    = (*dylib.LazyProc)(nil)
	dSetPropertySecValue = (*dylib.LazyProc)(nil)

	// Printer
	printer_Instance = (*dylib.LazyProc)(nil)

	// guid
	dGUIDToString = (*dylib.LazyProc)(nil)
	dStringToGUID = (*dylib.LazyProc)(nil)
	dCreateGUID   = (*dylib.LazyProc)(nil)

	// libResouces
	dGetLibResouceCount = (*dylib.LazyProc)(nil)
	dGetLibResouceItem  = (*dylib.LazyProc)(nil)
	dModifyLibResouce   = (*dylib.LazyProc)(nil)
	dLibAbout           = (*dylib.LazyProc)(nil)

	// 库的信息
	dLibStringEncoding = (*dylib.LazyProc)(nil)
	dLibVersion        = (*dylib.LazyProc)(nil)

	dMainThreadId    = (*dylib.LazyProc)(nil)
	dCurrentThreadId = (*dylib.LazyProc)(nil)
)

func DoDefInit() {
	application_Instance = libvcl.NewProc("Application_Instance")
	application_CreateForm = libvcl.NewProc("Application_CreateForm")
	application_Run = libvcl.NewProc("Application_Run")
	application_Initialize = libvcl.NewProc("Application_Initialize")

	form_Create2 = libvcl.NewProc("Form_Create2")
	form_EnabledMaximize = libvcl.NewProc("Form_EnabledMaximize")
	form_EnabledMinimize = libvcl.NewProc("Form_EnabledMinimize")
	form_EnabledSystemMenu = libvcl.NewProc("Form_EnabledSystemMenu")
	form_SetAllowDropFiles = libvcl.NewProc("Form_SetAllowDropFiles")
	form_GetAllowDropFiles = libvcl.NewProc("Form_GetAllowDropFiles")
	form_SetOnDropFiles = libvcl.NewProc("Form_SetOnDropFiles")
	form_SetOnDestroy = libvcl.NewProc("Form_SetOnDestroy")
	form_SetOnConstrainedResize = libvcl.NewProc("Form_SetOnConstrainedResize")
	form_SetOnDeactivate = libvcl.NewProc("Form_SetOnDeactivate")
	form_SetOnActivate = libvcl.NewProc("Form_SetOnActivate")
	form_SetOnStyleChanged = libvcl.NewProc("Form_SetOnStyleChanged")
	form_SetOnWndProc = libvcl.NewProc("Form_SetOnWndProc")
	form_SetShowInTaskBar = libvcl.NewProc("Form_SetShowInTaskBar")
	form_ShowInTaskBar = libvcl.NewProc("Form_ShowInTaskBar")
	form_ScaleForCurrentDpi = libvcl.NewProc("Form_ScaleForCurrentDpi")
	form_InheritedWndProc = libvcl.NewProc("Form_InheritedWndProc")

	setEventCallback = libvcl.NewProc("SetEventCallback")
	setMessageCallback = libvcl.NewProc("SetMessageCallback")
	setThreadSyncCallback = libvcl.NewProc("SetThreadSyncCallback")

	dGetStringArrOf = libvcl.NewProc("DGetStringArrOf")
	dStrLen = libvcl.NewProc("DStrLen")
	dMove = libvcl.NewProc("DMove")

	dShowMessage = libvcl.NewProc("DShowMessage")
	dMessageDlg = libvcl.NewProc("DMessageDlg")

	mouse_Instance = libvcl.NewProc("Mouse_Instance")
	screen_Instance = libvcl.NewProc("Screen_Instance")

	dSynchronize = libvcl.NewProc("DSynchronize")

	// TMenuItem
	dTextToShortCut = libvcl.NewProc("DTextToShortCut")
	dShortCutToText = libvcl.NewProc("DShortCutToText")

	// TClipboard
	clipboard_Instance = libvcl.NewProc("Clipboard_Instance")
	clipboard_SetClipboard = libvcl.NewProc("Clipboard_SetClipboard")

	// DSysOpen
	dSysOpen = libvcl.NewProc("DSysOpen")

	// TMemoryStream
	memoryStream_Read = libvcl.NewProc("MemoryStream_Read")
	memoryStream_Write = libvcl.NewProc("MemoryStream_Write")

	// TCanvas
	canvas_BrushCopy = libvcl.NewProc("Canvas_BrushCopy")
	canvas_CopyRect = libvcl.NewProc("Canvas_CopyRect")
	canvas_Draw1 = libvcl.NewProc("Canvas_Draw1")
	canvas_Draw2 = libvcl.NewProc("Canvas_Draw2")
	canvas_DrawFocusRect = libvcl.NewProc("Canvas_DrawFocusRect")
	canvas_FillRect = libvcl.NewProc("Canvas_FillRect")
	canvas_FrameRect = libvcl.NewProc("Canvas_FrameRect")
	canvas_StretchDraw = libvcl.NewProc("Canvas_StretchDraw")
	canvas_TextRect1 = libvcl.NewProc("Canvas_TextRect1")
	canvas_TextRect2 = libvcl.NewProc("Canvas_TextRect2")
	canvas_Polygon = libvcl.NewProc("Canvas_Polygon")
	canvas_Polyline = libvcl.NewProc("Canvas_Polyline")
	canvas_PolyBezier = libvcl.NewProc("canvas_PolyBezier")
	canvas_PolyBezierTo = libvcl.NewProc("Canvas_PolyBezierTo")
	canvas_Pixels = libvcl.NewProc("Canvas_Pixels")
	canvas_SetPixels = libvcl.NewProc("Canvas_SetPixels")

	// TImageList
	imageList_Draw1 = libvcl.NewProc("ImageList_Draw1")
	imageList_Draw2 = libvcl.NewProc("ImageList_Draw2")
	imageList_DrawOverlay1 = libvcl.NewProc("ImageList_DrawOverlay1")
	imageList_DrawOverlay2 = libvcl.NewProc("ImageList_DrawOverlay2")
	imageList_GetIcon1 = libvcl.NewProc("ImageList_GetIcon1")
	imageList_GetIcon2 = libvcl.NewProc("ImageList_GetIcon2")

	// TBitmap
	bitmap_Clear = libvcl.NewProc("Bitmap_Clear")
	bitmap_BeginUpdate = libvcl.NewProc("Bitmap_BeginUpdate")
	bitmap_EndUpdate = libvcl.NewProc("Bitmap_EndUpdate")
	bitmap_LoadFromDevice = libvcl.NewProc("Bitmap_LoadFromDevice")

	dExtractFilePath = libvcl.NewProc("DExtractFilePath")
	dFileExists = libvcl.NewProc("DFileExists")

	dSelectDirectory1 = libvcl.NewProc("DSelectDirectory1")
	dSelectDirectory2 = libvcl.NewProc("DSelectDirectory2")
	dInputBox = libvcl.NewProc("DInputBox")
	dInputQuery = libvcl.NewProc("DInputQuery")

	// TForm相关设置
	setGlobalFormScaled = libvcl.NewProc("SetGlobalFormScaled")
	form_ScaleForPPI = libvcl.NewProc("Form_ScaleForPPI")
	form_ScaleControlsForDpi = libvcl.NewProc("Form_ScaleControlsForDpi")

	// TSysLocaled
	dSysLocale = libvcl.NewProc("DSysLocale")

	// Shortcut
	dCreateURLShortCut = libvcl.NewProc("DCreateURLShortCut")
	dCreateShortCut = libvcl.NewProc("DCreateShortCut")

	// SetProperty
	dSetPropertyValue = libvcl.NewProc("DSetPropertyValue")
	dSetPropertySecValue = libvcl.NewProc("DSetPropertySecValue")

	// Printer
	printer_Instance = libvcl.NewProc("Printer_Instance")

	// guid
	dGUIDToString = libvcl.NewProc("DGUIDToString")
	dStringToGUID = libvcl.NewProc("DStringToGUID")
	dCreateGUID = libvcl.NewProc("DCreateGUID")

	// libResouces
	dGetLibResouceCount = libvcl.NewProc("DGetLibResouceCount")
	dGetLibResouceItem = libvcl.NewProc("DGetLibResouceItem")
	dModifyLibResouce = libvcl.NewProc("DModifyLibResouce")
	dLibAbout = libvcl.NewProc("DLibAbout")

	// 库的信息
	dLibStringEncoding = libvcl.NewProc("DLibStringEncoding")
	dLibVersion = libvcl.NewProc("DLibVersion")

	dMainThreadId = libvcl.NewProc("DMainThreadId")
	dCurrentThreadId = libvcl.NewProc("DCurrentThreadId")
}
