package switchs

import (
	"errors"
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	PortListReq  struct{ SwitchId uint }
	PortListResp struct {
		Success string `json:"success"`
		Error   string `json:"error"` //  错误信息（错误信息，当success===false时才会存在）
		Ports   []Port `json:"ports"` // 端口列表
	}
	Port struct {
		Id                uint   `json:"id"`
		SwitchId          uint   `json:"switch_id"`        // 交换机id
		PortName          string `json:"port_name"`        // 端口名称
		PortType          string `json:"port_type"`        // 端口类型
		PortRemark        string `json:"port_remark"`      // 端口备注
		Type              string `json:"type"`             // 类型
		Disconnected      int    `json:"disconnected"`     // 断线状态
		Disabled          bool   `json:"disabled"`         // 禁用状态：false启用;true禁用
		VlanType          string `json:"vlan_type"`        // VLAN类型
		Speed             any    `json:"speed"`            // 链路速率
		Mtu               any    `json:"mtu"`              // MTU
		BindMac           int    `json:"bind_mac"`         // 绑定mac状态
		BindArp           int    `json:"bind_arp"`         // 绑定arp状态
		UpPort            uint   `json:"up_port"`          // 上行端口
		DownPort          uint   `json:"down_port"`        // 下载端口
		MacAddress        string `json:"mac_address"`      // mac地址
		PhysMacAddress    any    `json:"phys_mac_address"` // 物理mac地址
		MacReachTime      any    `json:"mac_reach_time"`   // mac延伸时间
		Remarks           string `json:"remarks"`          // 备注
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
		MacAddressesCount int    `json:"mac_addresses_count"` // mac地址数量
		Device            struct {
			Id     uint   `json:"id"`      // 交换机id
			Name   string `json:"name"`    // 交换机名称
			Type   string `json:"type"`    // 交换机端口类型
			MainIp string `json:"main_ip"` // 交换机ip地址
		} `json:"device"` // 驱动
	}
)

func (r *PortListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/%d/ports", r.SwitchId)
}
func (r *PortListReq) Method() (method string)     { return http.MethodGet }
func (r *PortListReq) Values() (values url.Values) { return }
func (r *PortListReq) Body() (body any)            { return }

func (r PortListResp) Ok() (ok bool)    { return r.Success == "true" }
func (r PortListResp) Err() (err error) { return errors.New(r.Error) }

// PortList 获取交换机端口列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389879/detail/5754666
func PortList(ctx *dcimsdk.Context, request *PortListReq) (resp PortListResp, err error) {
	return dcimsdk.Execute[*PortListReq, PortListResp](ctx, request, dcimsdk.DeviceTypeFixed)
}
