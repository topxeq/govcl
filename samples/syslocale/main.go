package main

import (
	"fmt"

	_ "github.com/topxeq/govcl/pkgs/winappres"
	"github.com/topxeq/govcl/vcl/rtl"
)

func main() {
	fmt.Println(rtl.SysLocale)
}
