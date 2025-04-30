package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type PortUpdateStatusReq struct {
	PortId uint `json:"-"`      // 交换机端口Id
	Enable bool `json:"enable"` // false:禁用;true:启用
}

func (r *PortUpdateStatusReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/ports/%d", r.PortId)
}
func (r *PortUpdateStatusReq) Method() (method string)     { return http.MethodPost }
func (r *PortUpdateStatusReq) Values() (values url.Values) { return }
func (r *PortUpdateStatusReq) Body() (body any)            { return r }

// action 修改交换机端口启用禁用状态 TODO 此方法再文档中未找到，从页面搬过来的
func updateStatus(ctx *dcimsdk.Context, request *PortUpdateStatusReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*PortUpdateStatusReq, dcimsdk.CreateUpdateResp](ctx, request)
}

// PortEnable 启用
func PortEnable(ctx *dcimsdk.Context, portId uint) (resp dcimsdk.CreateUpdateResp, err error) {
	return updateStatus(ctx, &PortUpdateStatusReq{PortId: portId, Enable: true})
}

// PortDisable 禁用
func PortDisable(ctx *dcimsdk.Context, portId uint) (resp dcimsdk.CreateUpdateResp, err error) {
	return updateStatus(ctx, &PortUpdateStatusReq{PortId: portId, Enable: false})
}
