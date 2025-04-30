package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type MacLockReq struct{ PortId uint }

func (r *MacLockReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/ports/%d/bind/mac", r.PortId)
}
func (r *MacLockReq) Method() (method string)     { return http.MethodPut }
func (r *MacLockReq) Values() (values url.Values) { return }
func (r *MacLockReq) Body() (body any)            { return r }

// PortMacLock 加锁交换机mac地址绑定
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5756321
func PortMacLock(ctx *dcimsdk.Context, request *MacLockReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*MacLockReq, dcimsdk.CreateUpdateResp](ctx, request)
}
