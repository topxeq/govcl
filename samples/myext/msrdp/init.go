package msrdp

import "github.com/topxeq/govcl/vcl"

// 初始
func init() {
    vcl.RegisterExtEventCallback(eventCallback)
}
