//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import "github.com/topxeq/govcl/vcl/api"

// Shortcut

/*
   CN: 创建一个url的快捷方式
   EN: Create a shortcut to a URL

   rtl.CreateURLShortCut("C:\\aaa\\bbb\\", "govcl", "https://github.com/ying32/govcl")
*/
func CreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	api.DCreateURLShortCut(aDestPath, aShortCutName, aURL)
}

/*
   CN: 创建一个快捷方式
   EN: Create a shortcut

   1. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "", "")
   2. rtl.CreateShortCut("C:\\Users\\administrator\\Desktop\\", "govcl", os.Args[0], "", "Description text", "-a -b")
*/
func CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return api.DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs)
}
