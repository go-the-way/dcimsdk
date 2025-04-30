package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type UpdateReq struct {
	Id uint `json:"-"`

	Name            string          `json:"name"`             // 名称
	SnmpSupport     uint            `json:"snmp_support"`     // 控制snmp
	ControlSupport  uint            `json:"control_support"`  // 手动控制
	ControlProtocol string          `json:"control_protocol"` // 手动控制协议
	Device          string          `json:"device"`           // 驱动
	Type            string          `json:"type"`             // 类型
	CabinetId       uint            `json:"cabinet_id"`       // 机柜id
	Remark          string          `json:"remark"`           // 备注
	MainIpv4Address MainIpv4Address `json:"main_ipv4address"`
}

func (r *UpdateReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/%d", r.Id)
}
func (r *UpdateReq) Method() (method string)     { return http.MethodPut }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Body() (body any)            { return r }

// Update 修改交换机
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5754604
func Update(ctx *dcimsdk.Context, request *UpdateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*UpdateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
