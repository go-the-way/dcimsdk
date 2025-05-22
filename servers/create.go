package servers

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"
)

type (
	SwitchPort struct {
		Id     uint `json:"id"`
		PortId uint `json:"port_id"`
	}
	Ipv4 struct {
		BlockId   string `json:"block_id"`   // ip段id
		AddressId string `json:"address_id"` // ip地址id
	}
	IpMetadata struct {
		AddressId     uint   `json:"address_id,omitempty"`
		Address       string `json:"address,omitempty"`
		Username      string `json:"username,omitempty"`
		Password      string `json:"password,omitempty"`
		BlockId       uint   `json:"block_id,omitempty"`
		Cidr          string `json:"cidr,omitempty"`
		IpmiBlockId   uint   `json:"ipmi_block_id,omitempty"`
		Id            uint   `json:"id,omitempty"`
		IpmiAddressId uint   `json:"ipmi_address_id,omitempty"`
		IpmiAddress   string `json:"ipmi_address,omitempty"`
	}
	CreateReq struct {
		Name              string             `json:"name"`
		Hostname          string             `json:"hostname"`
		Datacenter        dcimsdk.OnlyIdType `json:"datacenter"`          // 数据中心信息(机房)
		Cabinet           dcimsdk.OnlyIdType `json:"cabinet"`             // 机柜信息
		Hardware          dcimsdk.OnlyIdType `json:"hardware"`            // 硬件信息
		Switch            []SwitchPort       `json:"switch,omitempty"`    // 交换机信息（必填）
		IpmiMetadata      IpMetadata         `json:"ipmi_metadata"`       // ipmi元信息
		Ipv4              []Ipv4             `json:"ipv4"`                // IP地址
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
		Unit              []int              `json:"unit"`                // 单元容量
		BladeId           string             `json:"blade_id"`            // 刀片服务器架构id
		BladeUnit         string             `json:"blade_unit"`          // 刀片单元容量
	}
	CreateResp struct {
		Server  Server `json:"-"`
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}
)

func (r *CreateReq) Url() (url string)           { return "/api/admin/servers" }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

func (r CreateResp) Ok() (ok bool)    { return r.Success }
func (r CreateResp) Err() (err error) { return errors.New(r.Error) }

// Create 添加服务器
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5779927
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp CreateResp, err error) {
	resp0, err0 := dcimsdk.Execute[*CreateReq, dcimsdk.CreateUpdateResp](ctx, request)
	if err0 != nil {
		return CreateResp{Success: resp0.Success, Error: resp0.Error}, err0
	}

	resp1, err1 := dcimsdk.Execute[*ListReq, ListResp](ctx, &ListReq{Name: request.Name, Page: 1, PageSize: 10000})
	if err0 != nil {
		return CreateResp{Success: resp1.Success, Error: resp1.Error}, err1
	}

	resp.Success = true
	for _, srv := range resp1.Server {
		if srv.Name == request.Name {
			resp.Server = srv
			break
		}
	}

	return
}
