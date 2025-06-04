package servers

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-the-way/dcimsdk"
)

type UpdateReq struct {
	Id                uint               `json:"id"` // 服务器id
	Name              string             `json:"name"`
	Hostname          string             `json:"hostname"`
	Datacenter        dcimsdk.OnlyIdType `json:"datacenter"`          // 数据中心信息(机房)
	Cabinet           dcimsdk.OnlyIdType `json:"cabinet"`             // 机柜信息
	Hardware          dcimsdk.OnlyIdType `json:"hardware"`            // 硬件信息
	Switch            []SwitchPort       `json:"switch"`              // 交换机信息（必填）
	IpmiMetadata      IpMetadata         `json:"ipmi_metadata"`       // ipmi元信息
	Type              string             `json:"type"`                // 类型
	Card              int                `json:"card"`                // 网卡数量
	Part              struct{}           `json:"part"`                // 组成部分
	ShelfTime         string             `json:"shelf_time"`          // 上架时间
	Gid               uint               `json:"gid,omitempty"`       // 分组id
	DisabledNetwork   int                `json:"disabled_network"`    // 是否禁用网络
	NetworkMetadata   int                `json:"network_metadata"`    // 网络元数据
	Status            string             `json:"status"`              // 状态
	Osid              string             `json:"osid"`                // 镜像名称
	Disabled          bool               `json:"disabled"`            // 禁用状态
	Username          string             `json:"username"`            // 用户名
	Password          string             `json:"password"`            // 密码
	MainMac           string             `json:"main_mac"`            // 主mac地址
	Bandwidth         int                `json:"bandwidth"`           // 带宽
	AllowRemoteSwitch bool               `json:"allow_remote_switch"` // 是否允许远程交换机
	Force             bool               `json:"force"`               // 是否强制
	DhcpEnable        int                `json:"dhcp_enable"`         // dhcp可用状态
	Unit              []int              `json:"unit,omitempty"`      // 单元容量
	BladeId           string             `json:"blade_id"`            // 刀片服务器架构id
	BladeUnit         string             `json:"blade_unit"`          // 刀片单元容量
}

func (r *UpdateReq) Url() (url string)           { return fmt.Sprintf("/api/admin/servers/%d", r.Id) }
func (r *UpdateReq) Method() (method string)     { return http.MethodPut }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Body() (body any)            { return r }

// Update 修改服务器
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770188
func Update(ctx *dcimsdk.Context, request *UpdateReq, timeout ...time.Duration) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*UpdateReq, dcimsdk.CreateUpdateResp](ctx, request, dcimsdk.OptionFn(timeout...))
}
