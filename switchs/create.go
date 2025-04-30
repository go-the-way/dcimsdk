package switchs

import (
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
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
	MainIpv4Address struct {
		Id uint `json:"id"`
	}
)

func (r *CreateReq) Url() (url string)           { return "/api/admin/switchs" }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

// Create 添加交换机
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5750695
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*CreateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
