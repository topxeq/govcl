// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/topxeq/govcl/vcl"
)

type TFrame1 struct {
    *vcl.TFrame
    Button1 *vcl.TButton

    //::private::
    TFrame1Fields
}


func NewFrame1(owner vcl.IComponent) (root *TFrame1)  {
    vcl.CreateResFrame(owner, &root)
    return
}

var frame1Bytes = []byte("\x54\x50\x46\x30\x07\x54\x46\x72\x61\x6D\x65\x31\x06\x46\x72\x61\x6D\x65\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x23\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xD7\x01\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x23\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xD7\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x0A\x44\x65\x73\x69\x67\x6E\x4C\x65\x66\x74\x03\xA2\x01\x09\x44\x65\x73\x69\x67\x6E\x54\x6F\x70\x02\x1C\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x03\xD0\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x79\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x46\x72\x61\x6D\x65\x31\x07\x4F\x6E\x43\x6C\x69\x63\x6B\x07\x0C\x42\x75\x74\x74\x6F\x6E\x31\x43\x6C\x69\x63\x6B\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(TFrame1{}, &frame1Bytes)
