// +build windows

package miniblink

import "github.com/topxeq/govcl/vcl/rtl"

var (
	isLcl bool
)

func init() {
	isLcl = rtl.LcLLoaded()
	//Initialize()

}
