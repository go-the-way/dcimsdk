package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type MainIpReq struct {
	ServerId  uint `json:"-"` // 服务器id
	AddressId uint `json:"-"` // ip地址id
}

func (r *MainIpReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/ipv4/address/%d", r.ServerId, r.AddressId)
}
func (r *MainIpReq) Method() (method string)     { return http.MethodPost }
func (r *MainIpReq) Values() (values url.Values) { return }
func (r *MainIpReq) Body() (body any)            { return }

// MainIp 设置主ip
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770224
func MainIp(ctx *dcimsdk.Context, request *MainIpReq, opts ...dcimsdk.OptionFunc) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*MainIpReq, dcimsdk.CreateUpdateResp](ctx, request, opts...)
}
