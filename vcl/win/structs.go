//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package win

import . "github.com/topxeq/govcl/vcl/types"

type TRGBQuad struct {
	RgbBlue     uint8
	RgbGreen    uint8
	RgbRed      uint8
	RgbReserved uint8
}

type TBitmapInfoHeader struct {
	BiSize          uint32
	BiWidth         uint32
	BiHeight        uint32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter uint32
	BiYPelsPerMeter uint32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type TBitmapInfo struct {
	BmiHeader TBitmapInfoHeader
	BmiColors [1]TRGBQuad
}

type TBlendFunction struct {
	BlendOp             uint8
	BlendFlags          uint8
	SourceConstantAlpha uint8
	AlphaFormat         uint8
}

type TSystemInfo struct {

	//0: (
	//dwOemId: DWORD);
	//1: (
	ProcessorArchitecture uint16
	Reserved              uint16

	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type TSecurityAttributes struct {
	nLength              uint32
	lpSecurityDescriptor uintptr
	bInheritHandle       bool // BOOL
}

type TSHItemID struct {
	Cb   uint16  // Size of the ID (including cb itself)
	AbID [1]byte // The item ID (variable length)
}

type TItemIDList struct {
	Mkid TSHItemID
}

// ShellExecuteEx
type TShellExecuteInfo struct {
	CbSize       uint32
	FMask        uint32
	Wnd          HWND
	LpVerb       LPCWSTR
	LpFile       LPCWSTR
	LpParameters LPCWSTR
	LpDirectory  LPCWSTR
	NShow        int32
	HInstApp     HINST
	// Optional fields
	LpIDList  uintptr
	LpClass   LPCWSTR
	HkeyClass HKEY
	DwHotKey  uint32

	//	0: (
	// HICON
	//	1: (
	HIcon_hMonitor uintptr

	HProcess uintptr
}
