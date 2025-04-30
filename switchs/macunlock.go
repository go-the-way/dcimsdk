package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type MacUnLockReq struct{ PortId uint }

func (r *MacUnLockReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/ports/%d/bind/mac", r.PortId)
}
func (r *MacUnLockReq) Method() (method string)     { return http.MethodDelete }
func (r *MacUnLockReq) Values() (values url.Values) { return }
func (r *MacUnLockReq) Body() (body any)            { return }

// PortMacUnLock 解锁交换机mac地址绑定
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5756312
func PortMacUnLock(ctx *dcimsdk.Context, request *MacUnLockReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*MacUnLockReq, dcimsdk.CreateUpdateResp](ctx, request)
}
