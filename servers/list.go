package servers

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		Name     string `url:"name"`
		Page     uint   `url:"page"`
		PageSize uint   `url:"page_size"`
	}
	ListResp struct {
		Success bool     `json:"success"`
		Error   string   `json:"error"`
		Server  []Server `json:"server"` // 服务器列表
		Total   int      `json:"total"`  // 总条数
	}
	Server struct {
		Id              uint   `json:"id"`            // 服务器id
		Gid             int    `json:"gid"`           // 服务器分组id
		DatacenterId    uint   `json:"datacenter_id"` // 机房id
		CabinetId       uint   `json:"cabinet_id"`    // 机柜id
		Name            string `json:"name"`          // 服务器名称
		Hostname        string `json:"hostname"`      // 主机名
		PowerStatus     string `json:"power_status"`  // 电源状态
		Type            string `json:"type"`          // 类型
		MainMac         string `json:"main_mac"`      // 主mac地址
		MainIpv4Address struct {
			Address   string `json:"address"`    // ip地址
			Cidr      string `json:"cidr"`       // cidr
			Netmask   string `json:"netmask"`    // 子网掩码
			Gateway   string `json:"gateway"`    // 网关
			IsBlock   bool   `json:"is_block"`   // 是否分段
			Vlan      string `json:"vlan"`       // vlan
			GroupName string `json:"group_name"` // 分组名称
		} `json:"main_ipv4address"` // 主ip地址信息
		IpmiMetadata struct {
			AddressId       uint   `json:"address_id"`        // ip地址id
			Address         string `json:"address"`           // ip地址
			Username        string `json:"username"`          // 用户名
			Password        string `json:"password"`          // 密码
			BlockId         uint   `json:"block_id"`          // ip分段id
			PublicBlockId   uint   `json:"public_block_id"`   // 公网ip段id
			PublicAddressId uint   `json:"public_address_id"` // 公网ip地址id
			PublicAddress   string `json:"public_address"`    // 公网ip地址
			IpmiBlockId     uint   `json:"ipmi_block_id"`     // 公网ip地址id
			IpmiAddressId   uint   `json:"ipmi_address_id"`   // ipmi ip地址id
			IpmiAddress     string `json:"ipmi_address"`      // ipmi ip地址
			Cidr            string `json:"cidr"`              // cidr
		} `json:"ipmi_metadata"` // ipmi元数据
		Bandwidth    int    `json:"bandwidth"`     // 带宽
		Status       string `json:"status"`        // 状态
		DhcpEnable   int    `json:"dhcp_enable"`   // dhcp可用状态
		Osid         string `json:"osid"`          // 系统id
		Remark       string `json:"remark"`        // 备注
		RemarkUser   string `json:"remark_user"`   // 备注用户
		Disabled     int    `json:"disabled"`      // 禁用状态
		Username     string `json:"username"`      // 用户名
		Password     string `json:"password"`      // 密码
		RemotePort   int    `json:"remote_port"`   // 远程端口
		Card         int    `json:"card"`          // 网卡数量
		IsLock       bool   `json:"is_lock"`       // 锁定状态
		IsPreinstall int    `json:"is_preinstall"` // 是否预装系统
		CreatedAt    string `json:"created_at"`    // 创建时间
		UpdatedAt    string `json:"updated_at"`    // 更新时间
		Ipv4Count    int    `json:"ipv4_count"`    // ipv4数量
		User         []struct {
			Id   int    `json:"id"`   // 用户id
			Name string `json:"name"` // 用户名称
		} `json:"user"` // 用户信息
		Switch  []Switch `json:"switch"`   // 交换机信息
		IsIssue bool     `json:"is_issue"` // 是否故障
		Part    struct {
			Field1 any `json:"1"`
			Field2 any `json:"2"`
			Field3 any `json:"3"`
			Field4 any `json:"4"`
		} `json:"part"` // 硬件信息(TODO 里面属性暂时不写，有需要再补上)
		TotalFlow              int    `json:"total_flow"`                // 总流量
		Used                   int    `json:"used"`                      // 已使用
		Flow                   int    `json:"flow"`                      // 流量
		OtherFlow              int    `json:"other_flow"`                // 其他流量
		RealHardwareHistoryOne any    `json:"real_hardware_history_one"` // 真实硬件历史
		Hardware               string `json:"hardware"`                  // 硬件
		Cabinet                struct {
			Id     int    `json:"id"`      // 机柜id
			Name   string `json:"name"`    // 机柜名称
			IsLock int    `json:"is_lock"` // 锁定状态
		} `json:"cabinet"` // 机柜信息
		Datacenter struct {
			Id   int    `json:"id"`   // 数据中心id
			Name string `json:"name"` // 数据中心名称
		} `json:"datacenter"` // 数据中心信息(机房)
		Packages []struct {
			Id   int    `json:"id"`   // 套餐包id
			Name string `json:"name"` // 套餐包名称
		} `json:"packages"` // 套餐包信息
		Lock []any `json:"lock"` // 锁定状态
	}
	Switch struct {
		Id       uint     `json:"id"`        // 交换机id
		PortName string   `json:"port_name"` // 端口名
		UpPort   int      `json:"up_port"`   // 上行端口
		DownPort int      `json:"down_port"` // 下载端口
		VlanType string   `json:"vlan_type"` // vlan类型
		Vlans    []string `json:"vlans"`     // vlan类表
		PortId   uint     `json:"port_id"`   // 端口id
		Switch   string   `json:"switch"`    // 交换机ip地址
		IsLock   int      `json:"is_lock"`   // 锁定状态
		Disabled bool     `json:"-"`
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/servers" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取服务器列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5760736
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request, dcimsdk.IdTypeFixed)
}
