// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
	"github.com/topxeq/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Panel1       *vcl.TPanel
	Panel2       *vcl.TPanel
	Button1      *vcl.TButton
	Button2      *vcl.TButton
	Button3      *vcl.TButton
	PnlVideo     *vcl.TPanel
	StatusBar1   *vcl.TStatusBar
	ActionList1  *vcl.TActionList
	ActPlay      *vcl.TAction
	ActPause     *vcl.TAction
	ActStop      *vcl.TAction
	LblCurTime   *vcl.TLabel
	LblTotalTime *vcl.TLabel
	Timer1       *vcl.TTimer
	TrackBar1    *vcl.TTrackBar

	//::private::
	TMainFormFields
}

var MainForm *TMainForm

// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

var (
	mainFormBytes = []byte{
		0x54, 0x50, 0x46, 0x30, 0x05, 0x54, 0x46, 0x6F, 0x72, 0x6D, 0x08, 0x4D,
		0x61, 0x69, 0x6E, 0x46, 0x6F, 0x72, 0x6D, 0x04, 0x4C, 0x65, 0x66, 0x74,
		0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x00, 0x07, 0x43, 0x61, 0x70,
		0x74, 0x69, 0x6F, 0x6E, 0x14, 0x0C, 0x00, 0x00, 0x00, 0x6C, 0x69, 0x62,
		0x76, 0x6C, 0x63, 0xE6, 0xB5, 0x8B, 0xE8, 0xAF, 0x95, 0x0C, 0x43, 0x6C,
		0x69, 0x65, 0x6E, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x03, 0x41,
		0x02, 0x0B, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x57, 0x69, 0x64, 0x74,
		0x68, 0x03, 0xD4, 0x03, 0x05, 0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x09,
		0x63, 0x6C, 0x42, 0x74, 0x6E, 0x46, 0x61, 0x63, 0x65, 0x0E, 0x44, 0x6F,
		0x75, 0x62, 0x6C, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64,
		0x09, 0x0C, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x43, 0x68, 0x61, 0x72, 0x73,
		0x65, 0x74, 0x07, 0x0F, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4C, 0x54, 0x5F,
		0x43, 0x48, 0x41, 0x52, 0x53, 0x45, 0x54, 0x0A, 0x46, 0x6F, 0x6E, 0x74,
		0x2E, 0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x0C, 0x63, 0x6C, 0x57, 0x69,
		0x6E, 0x64, 0x6F, 0x77, 0x54, 0x65, 0x78, 0x74, 0x0B, 0x46, 0x6F, 0x6E,
		0x74, 0x2E, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0xF3, 0x09, 0x46,
		0x6F, 0x6E, 0x74, 0x2E, 0x4E, 0x61, 0x6D, 0x65, 0x06, 0x06, 0x54, 0x61,
		0x68, 0x6F, 0x6D, 0x61, 0x0A, 0x46, 0x6F, 0x6E, 0x74, 0x2E, 0x53, 0x74,
		0x79, 0x6C, 0x65, 0x0B, 0x00, 0x0E, 0x4F, 0x6C, 0x64, 0x43, 0x72, 0x65,
		0x61, 0x74, 0x65, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x08, 0x08, 0x50, 0x6F,
		0x73, 0x69, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x0A, 0x70, 0x6F, 0x44, 0x65,
		0x73, 0x69, 0x67, 0x6E, 0x65, 0x64, 0x0D, 0x50, 0x69, 0x78, 0x65, 0x6C,
		0x73, 0x50, 0x65, 0x72, 0x49, 0x6E, 0x63, 0x68, 0x02, 0x60, 0x0A, 0x54,
		0x65, 0x78, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x10, 0x00,
		0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E, 0x65,
		0x6C, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 0x03, 0x54, 0x6F,
		0x70, 0x02, 0x00, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xD4, 0x03,
		0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x32, 0x05, 0x41, 0x6C,
		0x69, 0x67, 0x6E, 0x07, 0x05, 0x61, 0x6C, 0x54, 0x6F, 0x70, 0x0A, 0x42,
		0x65, 0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62,
		0x76, 0x4E, 0x6F, 0x6E, 0x65, 0x0B, 0x53, 0x68, 0x6F, 0x77, 0x43, 0x61,
		0x70, 0x74, 0x69, 0x6F, 0x6E, 0x08, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72,
		0x64, 0x65, 0x72, 0x02, 0x00, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E,
		0x65, 0x6C, 0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x32, 0x04, 0x4C, 0x65,
		0x66, 0x74, 0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x03, 0xFF, 0x01, 0x05,
		0x57, 0x69, 0x64, 0x74, 0x68, 0x03, 0xD4, 0x03, 0x06, 0x48, 0x65, 0x69,
		0x67, 0x68, 0x74, 0x02, 0x2F, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07,
		0x08, 0x61, 0x6C, 0x42, 0x6F, 0x74, 0x74, 0x6F, 0x6D, 0x0A, 0x42, 0x65,
		0x76, 0x65, 0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62, 0x76,
		0x4E, 0x6F, 0x6E, 0x65, 0x0B, 0x53, 0x68, 0x6F, 0x77, 0x43, 0x61, 0x70,
		0x74, 0x69, 0x6F, 0x6E, 0x08, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64,
		0x65, 0x72, 0x02, 0x01, 0x00, 0x06, 0x54, 0x4C, 0x61, 0x62, 0x65, 0x6C,
		0x0A, 0x4C, 0x62, 0x6C, 0x43, 0x75, 0x72, 0x54, 0x69, 0x6D, 0x65, 0x04,
		0x4C, 0x65, 0x66, 0x74, 0x03, 0x51, 0x01, 0x03, 0x54, 0x6F, 0x70, 0x02,
		0x10, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x3E, 0x06, 0x48, 0x65,
		0x69, 0x67, 0x68, 0x74, 0x02, 0x10, 0x08, 0x41, 0x75, 0x74, 0x6F, 0x53,
		0x69, 0x7A, 0x65, 0x08, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E,
		0x06, 0x08, 0x30, 0x30, 0x3A, 0x30, 0x30, 0x3A, 0x30, 0x30, 0x00, 0x00,
		0x06, 0x54, 0x4C, 0x61, 0x62, 0x65, 0x6C, 0x0C, 0x4C, 0x62, 0x6C, 0x54,
		0x6F, 0x74, 0x61, 0x6C, 0x54, 0x69, 0x6D, 0x65, 0x04, 0x4C, 0x65, 0x66,
		0x74, 0x03, 0x72, 0x03, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x10, 0x05, 0x57,
		0x69, 0x64, 0x74, 0x68, 0x02, 0x3E, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68,
		0x74, 0x02, 0x10, 0x08, 0x41, 0x75, 0x74, 0x6F, 0x53, 0x69, 0x7A, 0x65,
		0x08, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69, 0x6F, 0x6E, 0x06, 0x08, 0x30,
		0x30, 0x3A, 0x30, 0x30, 0x3A, 0x30, 0x30, 0x00, 0x00, 0x07, 0x54, 0x42,
		0x75, 0x74, 0x74, 0x6F, 0x6E, 0x07, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E,
		0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x39, 0x03, 0x54, 0x6F, 0x70,
		0x02, 0x0E, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x4B, 0x06, 0x48,
		0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x06, 0x41, 0x63, 0x74, 0x69,
		0x6F, 0x6E, 0x07, 0x07, 0x41, 0x63, 0x74, 0x50, 0x6C, 0x61, 0x79, 0x08,
		0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x00, 0x00, 0x00,
		0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x07, 0x42, 0x75, 0x74,
		0x74, 0x6F, 0x6E, 0x32, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0x8A, 0x00,
		0x03, 0x54, 0x6F, 0x70, 0x02, 0x0D, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68,
		0x02, 0x4B, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x19, 0x06,
		0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x08, 0x41, 0x63, 0x74, 0x50,
		0x61, 0x75, 0x73, 0x65, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64, 0x65,
		0x72, 0x02, 0x01, 0x00, 0x00, 0x07, 0x54, 0x42, 0x75, 0x74, 0x74, 0x6F,
		0x6E, 0x07, 0x42, 0x75, 0x74, 0x74, 0x6F, 0x6E, 0x33, 0x04, 0x4C, 0x65,
		0x66, 0x74, 0x03, 0xE5, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x0D, 0x05,
		0x57, 0x69, 0x64, 0x74, 0x68, 0x02, 0x4B, 0x06, 0x48, 0x65, 0x69, 0x67,
		0x68, 0x74, 0x02, 0x19, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x07,
		0x07, 0x41, 0x63, 0x74, 0x53, 0x74, 0x6F, 0x70, 0x08, 0x54, 0x61, 0x62,
		0x4F, 0x72, 0x64, 0x65, 0x72, 0x02, 0x02, 0x00, 0x00, 0x09, 0x54, 0x54,
		0x72, 0x61, 0x63, 0x6B, 0x42, 0x61, 0x72, 0x09, 0x54, 0x72, 0x61, 0x63,
		0x6B, 0x42, 0x61, 0x72, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0x91,
		0x01, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x0E, 0x05, 0x57, 0x69, 0x64, 0x74,
		0x68, 0x03, 0xD4, 0x01, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02,
		0x1B, 0x03, 0x4D, 0x61, 0x78, 0x02, 0x64, 0x08, 0x50, 0x61, 0x67, 0x65,
		0x53, 0x69, 0x7A, 0x65, 0x02, 0x01, 0x0F, 0x50, 0x6F, 0x73, 0x69, 0x74,
		0x69, 0x6F, 0x6E, 0x54, 0x6F, 0x6F, 0x6C, 0x54, 0x69, 0x70, 0x07, 0x05,
		0x70, 0x74, 0x54, 0x6F, 0x70, 0x0C, 0x53, 0x68, 0x6F, 0x77, 0x53, 0x65,
		0x6C, 0x52, 0x61, 0x6E, 0x67, 0x65, 0x08, 0x08, 0x54, 0x61, 0x62, 0x4F,
		0x72, 0x64, 0x65, 0x72, 0x02, 0x03, 0x09, 0x54, 0x69, 0x63, 0x6B, 0x53,
		0x74, 0x79, 0x6C, 0x65, 0x07, 0x06, 0x74, 0x73, 0x4E, 0x6F, 0x6E, 0x65,
		0x00, 0x00, 0x00, 0x06, 0x54, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x08, 0x50,
		0x6E, 0x6C, 0x56, 0x69, 0x64, 0x65, 0x6F, 0x04, 0x4C, 0x65, 0x66, 0x74,
		0x02, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x02, 0x32, 0x05, 0x57, 0x69, 0x64,
		0x74, 0x68, 0x03, 0xD4, 0x03, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
		0x03, 0xCD, 0x01, 0x05, 0x41, 0x6C, 0x69, 0x67, 0x6E, 0x07, 0x08, 0x61,
		0x6C, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x0A, 0x42, 0x65, 0x76, 0x65,
		0x6C, 0x4F, 0x75, 0x74, 0x65, 0x72, 0x07, 0x06, 0x62, 0x76, 0x4E, 0x6F,
		0x6E, 0x65, 0x05, 0x43, 0x6F, 0x6C, 0x6F, 0x72, 0x07, 0x0C, 0x63, 0x6C,
		0x42, 0x61, 0x63, 0x6B, 0x67, 0x72, 0x6F, 0x75, 0x6E, 0x64, 0x10, 0x50,
		0x61, 0x72, 0x65, 0x6E, 0x74, 0x42, 0x61, 0x63, 0x6B, 0x67, 0x72, 0x6F,
		0x75, 0x6E, 0x64, 0x08, 0x0B, 0x53, 0x68, 0x6F, 0x77, 0x43, 0x61, 0x70,
		0x74, 0x69, 0x6F, 0x6E, 0x08, 0x08, 0x54, 0x61, 0x62, 0x4F, 0x72, 0x64,
		0x65, 0x72, 0x02, 0x02, 0x00, 0x00, 0x0A, 0x54, 0x53, 0x74, 0x61, 0x74,
		0x75, 0x73, 0x42, 0x61, 0x72, 0x0A, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
		0x42, 0x61, 0x72, 0x31, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x02, 0x00, 0x03,
		0x54, 0x6F, 0x70, 0x03, 0x2E, 0x02, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68,
		0x03, 0xD4, 0x03, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x02, 0x13,
		0x06, 0x50, 0x61, 0x6E, 0x65, 0x6C, 0x73, 0x0E, 0x00, 0x00, 0x00, 0x0B,
		0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x4C, 0x69, 0x73, 0x74, 0x0B,
		0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x4C, 0x69, 0x73, 0x74, 0x31, 0x04,
		0x4C, 0x65, 0x66, 0x74, 0x03, 0xD0, 0x00, 0x03, 0x54, 0x6F, 0x70, 0x03,
		0xA2, 0x00, 0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x07,
		0x41, 0x63, 0x74, 0x50, 0x6C, 0x61, 0x79, 0x07, 0x43, 0x61, 0x70, 0x74,
		0x69, 0x6F, 0x6E, 0x12, 0x02, 0x00, 0x00, 0x00, 0xAD, 0x64, 0x3E, 0x65,
		0x00, 0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x08, 0x41,
		0x63, 0x74, 0x50, 0x61, 0x75, 0x73, 0x65, 0x07, 0x43, 0x61, 0x70, 0x74,
		0x69, 0x6F, 0x6E, 0x12, 0x02, 0x00, 0x00, 0x00, 0x82, 0x66, 0x5C, 0x50,
		0x00, 0x00, 0x07, 0x54, 0x41, 0x63, 0x74, 0x69, 0x6F, 0x6E, 0x07, 0x41,
		0x63, 0x74, 0x53, 0x74, 0x6F, 0x70, 0x07, 0x43, 0x61, 0x70, 0x74, 0x69,
		0x6F, 0x6E, 0x12, 0x02, 0x00, 0x00, 0x00, 0x5C, 0x50, 0x62, 0x6B, 0x00,
		0x00, 0x00, 0x06, 0x54, 0x54, 0x69, 0x6D, 0x65, 0x72, 0x06, 0x54, 0x69,
		0x6D, 0x65, 0x72, 0x31, 0x07, 0x45, 0x6E, 0x61, 0x62, 0x6C, 0x65, 0x64,
		0x08, 0x04, 0x4C, 0x65, 0x66, 0x74, 0x03, 0xB1, 0x01, 0x03, 0x54, 0x6F,
		0x70, 0x03, 0x69, 0x01, 0x00, 0x00, 0x00}
)
var _ = vcl.RegisterFormResource(MainForm, &mainFormBytes)
