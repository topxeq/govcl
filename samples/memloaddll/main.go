package main

import (
	"github.com/topxeq/govcl/vcl"
	"github.com/topxeq/govcl/vcl/api"
)

func main() {
	// Free
	defer api.FeeMemoryDLL()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		vcl.ShowMessage(e.Message())
	})
	vcl.Application.CreateForm(&Form1)
	vcl.Application.CreateForm(&Form2)
	vcl.Application.Run()

}
