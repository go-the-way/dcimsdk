package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type IpAllotReq struct {
	ServerId uint   `json:"server_id"` // 服务器id
	Address  []uint `json:"address"`   // ip地址数组
}

func (r *IpAllotReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/ipv4/address", r.ServerId)
}
func (r *IpAllotReq) Method() (method string)     { return http.MethodPut }
func (r *IpAllotReq) Values() (values url.Values) { return }
func (r *IpAllotReq) Body() (body any)            { return map[string]any{"address": r.Address, "force": 1} }

// IpAllot 服务器分配ip
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770392
func IpAllot(ctx *dcimsdk.Context, request *IpAllotReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*IpAllotReq, dcimsdk.CreateUpdateResp](ctx, request)
}
