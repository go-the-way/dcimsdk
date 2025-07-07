package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type IpFreeReq struct {
	ServerId  uint // 服务器id
	AddressId uint // ip地址id
}

func (r *IpFreeReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/ipv4/address/%d", r.ServerId, r.AddressId)
}
func (r *IpFreeReq) Method() (method string)     { return http.MethodPut }
func (r *IpFreeReq) Values() (values url.Values) { return }
func (r *IpFreeReq) Body() (body any)            { return map[string]any{"force": 1} }

// IpFree 空闲ip
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770222
func IpFree(ctx *dcimsdk.Context, request *IpFreeReq, opts ...dcimsdk.OptionFunc) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*IpFreeReq, dcimsdk.CreateUpdateResp](ctx, request, opts...)
}
