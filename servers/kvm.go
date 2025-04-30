package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type KvmReq struct {
	ServerId int // 服务器id
}

func (r *KvmReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/kvm", r.ServerId)
}
func (r *KvmReq) Method() (method string)     { return http.MethodGet }
func (r *KvmReq) Values() (values url.Values) { return }
func (r *KvmReq) Body() (body any)            { return }

// Kvm 获取服务器 KVM
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5774909
func Kvm(ctx *dcimsdk.Context, request *KvmReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*KvmReq, dcimsdk.CreateUpdateResp](ctx, request)
}
