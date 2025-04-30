package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	VncReq  struct{ ServerId int }
	VncResp struct {
		Success  bool   `json:"success"`
		Error    string `json:"error"`
		Endpoint string `json:"endpoint"` // 被控主机
	}
)

func (r *VncReq) Url() (url string)           { return fmt.Sprintf("/api/admin/servers/%d/console", r.ServerId) }
func (r *VncReq) Method() (method string)     { return http.MethodGet }
func (r *VncReq) Values() (values url.Values) { return }
func (r *VncReq) Body() (body any)            { return }

// Vnc 获取vnc(测试)
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5773774
func Vnc(ctx *dcimsdk.Context, request *VncReq) (resp VncResp, err error) {
	return dcimsdk.Execute[*VncReq, VncResp](ctx, request)
}
