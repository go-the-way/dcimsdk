package servers

import (
	"errors"
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	PowerStatusReq  struct{ ServerId uint }
	PowerStatusResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Power   string `json:"power"`
	}
)

func (r *PowerStatusReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/power", r.ServerId)
}
func (r *PowerStatusReq) Method() (method string)     { return http.MethodGet }
func (r *PowerStatusReq) Values() (values url.Values) { return }
func (r *PowerStatusReq) Body() (body any)            { return }

func (r PowerStatusResp) Ok() (ok bool)    { return r.Success }
func (r PowerStatusResp) Err() (err error) { return errors.New(r.Error) }

// PowerStatus 获取服务器电源状态
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5774112
func PowerStatus(ctx *dcimsdk.Context, request *PowerStatusReq) (resp PowerStatusResp, err error) {
	return dcimsdk.Execute[*PowerStatusReq, PowerStatusResp](ctx, request, dcimsdk.PowerStatusFixed)
}
