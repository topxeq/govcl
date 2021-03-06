// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
	"github.com/topxeq/govcl/vcl"
)

type TLoginForm struct {
	*vcl.TForm
	ButtonLogin  *vcl.TButton
	Label1       *vcl.TLabel
	Label2       *vcl.TLabel
	EditUserName *vcl.TEdit
	EditPassword *vcl.TEdit
}

var LoginForm *TLoginForm

// 以字节形式加载
// vcl.Application.CreateForm(&LoginForm)

var (
	loginFormBytes = []byte{
		0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00,
		0x93, 0x48, 0xF6, 0xDA, 0x58, 0x85, 0xEB, 0x1B, 0x9D, 0x7F, 0x88, 0x1B,
		0x07, 0x78, 0x36, 0x50, 0x51, 0x2C, 0x90, 0x1C, 0x87, 0x84, 0x1F, 0xAA,
		0x29, 0x13, 0xA4, 0xDA, 0xBA, 0x86, 0xC0, 0x0A, 0x8A, 0x11, 0x2B, 0x53,
		0x5B, 0x9A, 0x05, 0x23, 0xD1, 0x1F, 0xE6, 0x61, 0x28, 0x59, 0xCA, 0x56,
		0x8F, 0x6E, 0xFE, 0xBD, 0x33, 0xAE, 0xB0, 0x7B, 0x81, 0xB2, 0x71, 0x6F,
		0x39, 0xCF, 0x5B, 0x94, 0xB5, 0xE8, 0x69, 0xC7, 0xDF, 0x53, 0xD7, 0xA5,
		0x88, 0x22, 0x63, 0x48, 0x9C, 0x74, 0x82, 0x2F, 0x72, 0x68, 0xEA, 0xD7,
		0x78, 0x75, 0xE7, 0x6C, 0x51, 0xCF, 0xF8, 0x78, 0xE9, 0xF4, 0x0F, 0xE3,
		0x1C, 0xAC, 0x1C, 0xFA, 0xEB, 0xB9, 0xDC, 0x9A, 0xEA, 0x11, 0x0E, 0xCF,
		0xC2, 0x90, 0x2B, 0x10, 0xFD, 0xC3, 0x82, 0xC0, 0x2C, 0x2F, 0x5F, 0xF2,
		0xEC, 0xCF, 0x46, 0xE2, 0x78, 0x34, 0x59, 0x53, 0xE8, 0xCF, 0x05, 0xE9,
		0xD7, 0x40, 0x2A, 0x6B, 0xE2, 0x65, 0x1B, 0x72, 0x9D, 0x86, 0x40, 0xA0,
		0xD9, 0x44, 0x07, 0xFB, 0x44, 0x45, 0x30, 0xFA, 0xEE, 0xBF, 0xB2, 0x2F,
		0xBE, 0x5E, 0xD2, 0xBA, 0x5D, 0x47, 0xC3, 0x28, 0xE5, 0x15, 0x50, 0x74,
		0x1F, 0x42, 0x18, 0xB6, 0x24, 0xF6, 0x90, 0xA6, 0x22, 0x8C, 0x1C, 0x63,
		0xC2, 0x89, 0x1F, 0x07, 0x11, 0x4D, 0x1F, 0xBB, 0x78, 0xF7, 0x3F, 0xAC,
		0x87, 0x43, 0xA0, 0x31, 0x5A, 0x87, 0xB1, 0x13, 0xC8, 0x22, 0x88, 0x64,
		0xEB, 0xFF, 0xCB, 0x7B, 0x28, 0x2F, 0x5D, 0xBE, 0x71, 0x8D, 0x98, 0xD6,
		0x92, 0xFD, 0x7B, 0x37, 0xE9, 0x70, 0x1F, 0xFB, 0xB2, 0x10, 0xD4, 0x50,
		0x51, 0xDC, 0x6A, 0xDA, 0xCB, 0xFD, 0xD4, 0x81, 0x66, 0xB6, 0xB0, 0x6E,
		0xEC, 0x17, 0x53, 0x6C, 0x71, 0x73, 0x6A, 0x7F, 0x23, 0x3A, 0x02, 0x06,
		0x89, 0x16, 0xBE, 0xE3, 0x2B, 0xD5, 0x16, 0x19, 0xEA, 0x6B, 0x0F, 0xD2,
		0x95, 0xB9, 0x6B, 0x10, 0x5F, 0xCF, 0xFF, 0x11, 0x0D, 0x27, 0x07, 0xAE,
		0xCA, 0x96, 0x13, 0x58, 0x97, 0x3C, 0x2D, 0x52, 0x08, 0x48, 0x10, 0x9C,
		0x7D, 0xA9, 0xA9, 0xCE, 0xB3, 0xFB, 0x0B, 0x1C, 0xD1, 0xA0, 0x44, 0x5D,
		0x2B, 0xA8, 0x41, 0xF0, 0x9B, 0xC1, 0x8A, 0x90, 0xB4, 0x53, 0x70, 0x57,
		0x3C, 0x66, 0x07, 0xAE, 0x8D, 0xBF, 0xD1, 0x00, 0xB7, 0x81, 0x2E, 0x22,
		0xE0, 0x82, 0xCA, 0xA9, 0x0B, 0xF7, 0x78, 0x54, 0x86, 0xAA, 0x66, 0x44,
		0x09, 0x8B, 0x27, 0xC0, 0x2F, 0x57, 0x18, 0x8D, 0x9C, 0xDA, 0x48, 0x54,
		0x89, 0x64, 0xBF, 0xED, 0x26, 0xC4, 0xAC, 0x56, 0xA1, 0x5F, 0x2B, 0x0E,
		0x4F, 0x81, 0xEB, 0x16, 0x9C, 0xB0, 0x39, 0x20, 0xB2, 0x11, 0xD4, 0x05,
		0xC1, 0xAF, 0xCC, 0x6C, 0xE3, 0x14, 0xD5, 0xD9, 0x05, 0x92, 0xC0, 0x29,
		0x5B, 0xDF, 0x54, 0x78, 0xFA, 0x38, 0x8F, 0x41, 0xAF, 0xE5, 0xCA, 0x94,
		0xCF, 0xB2, 0x27, 0x67, 0xFA, 0x41, 0xE5, 0xF6, 0xE3, 0x5C, 0x28, 0xE2,
		0x20, 0xFD, 0x42, 0x9D, 0x65, 0x7C, 0x3A, 0xB6, 0x77, 0xAA, 0x1B, 0x02,
		0x49, 0xF7, 0x41, 0x21, 0x2C, 0x1E, 0x63, 0x70, 0x7D, 0x8F, 0x90, 0x2B,
		0x4A, 0xAA, 0xEC, 0x37, 0xA8, 0x73, 0xC1, 0xA2, 0x61, 0x7B, 0xB8, 0x8D,
		0x9E, 0xAE, 0xE3, 0x63, 0x56, 0x09, 0xB3, 0x19, 0x08, 0xBC, 0x3D, 0x5E,
		0x16, 0x1C, 0xC6, 0xFA, 0x4A, 0x4A, 0x52, 0x97, 0x93, 0xF3, 0x86, 0xA4,
		0x65, 0x50, 0xC6, 0x15, 0x05, 0x72, 0xFA, 0x6B, 0x93, 0x27, 0xFA, 0x80,
		0x98, 0x78, 0xAB, 0x75, 0xEB, 0x7B, 0xCC, 0xB2, 0x70, 0xFE, 0x5C, 0x74,
		0xB6, 0xEF, 0xAC, 0x04, 0x93, 0x26, 0x93}
)
var _ = vcl.RegisterFormResource(LoginForm, &loginFormBytes)
