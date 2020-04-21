package main

import (
	"fmt"

	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	vcl.Application.CreateForm(formListViewDrawBytes, &FormListViewDraw)

	// 加载信息
	trainData, err := parseFromFile("testtraindata.json")
	if err == nil {
		// 设计器设计的ListView
		fullResFormListViewInstance(trainData)
		// 代码创建的ListView
		codeBuildistViewInstance(trainData)
	} else {
		fmt.Println("err:", err)
	}

	vcl.Application.Run()
}
