package ipv4

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		Page     uint   `url:"page"`
		PageSize uint   `url:"page_size"`
		Type     string `url:"type"`
	}
	ListResp struct {
		Success    bool    `json:"success"`
		Error      string  `json:"error"`
		Ipv4Blocks []Block `json:"ipv4blocks"` // IP段列表
		Total      int     `json:"total"`
	}
	Block struct {
		Id              uint     `json:"id"`              // ip地址id
		Path            string   `json:"path"`            // 子网
		Gid             uint     `json:"gid"`             // 分组id
		Type            string   `json:"type"`            // 类型
		Label           string   `json:"label"`           // 标签
		Disabled        int      `json:"disabled"`        // 状态
		Netmask         string   `json:"netmask"`         // 掩码
		Gateway         string   `json:"gateway"`         // 网关
		Vlan            string   `json:"vlan"`            // VLAN
		AssignTime      any      `json:"assign_time"`     // 分配时间
		ServerId        int      `json:"server_id"`       // 服务器id
		UserId          any      `json:"user_id"`         // 用户id
		Auto            int      `json:"auto"`            // 是否自动分配
		UsableipNum     int      `json:"usableip_num"`    // 可使用的ip数量
		UsingipNum      int      `json:"usingip_num"`     // 使用中的ip数量
		TotalipNum      int      `json:"totalip_num"`     // 使用中的ip数量
		SpecialipNum    int      `json:"specialip_num"`   // 特殊ip数量
		Remark          string   `json:"remark"`          // 备注
		IsLock          int      `json:"is_lock"`         // 是否锁定
		CreatedAt       string   `json:"created_at"`      // 创建时间
		UpdatedAt       string   `json:"updated_at"`      // 修改时间
		ServerName      string   `json:"server_name"`     // 服务器名称
		Dns1            string   `json:"dns_1"`           // DNS1
		Dns2            string   `json:"dns_2"`           // DNS2
		GroupName       string   `json:"group_name"`      // 分组名称
		Cidr            string   `json:"cidr"`            // CIDR
		IsPiecewise     bool     `json:"is_piecewise"`    // 是否分段
		FillIp          bool     `json:"fill_ip"`         // 是否自动填充ip
		CanAllocate     bool     `json:"can_allocate"`    // 是否可分配
		Username        string   `json:"username"`        // 用户名
		DatacenterId    []uint   `json:"datacenter_id"`   // 机房id
		DatacenterName  []string `json:"datacenter_name"` // 机房名称
		Child           []any    `json:"child"`
		AllIp           bool     `json:"all_ip"` // 是否全ip
		Ipv4Datacenters []struct {
			Id           int    `json:"id"`
			BlockId      int    `json:"block_id"`      // ip段id
			DatacenterId int    `json:"datacenter_id"` // 机房id
			CreatedAt    string `json:"created_at"`
			UpdatedAt    string `json:"updated_at"`
			Datacenter   struct {
				Id          int    `json:"id"`
				Name        string `json:"name"`        // 机房名称
				Geographicg string `json:"geographicg"` // 位置
				Bandwidth   int    `json:"bandwidth"`   // 带宽
				Voltage     int    `json:"voltage"`     // 电力
				Contact     struct {
					Name  string `json:"name"`  // 姓名
					Tel   string `json:"tel"`   // 电话
					Email string `json:"email"` // 邮箱
				} `json:"contact"` // 联系信息
				InternalCidr string `json:"internal_cidr"` // 重装系统使用网段
				VlanOs       int    `json:"vlan_os"`       // 重装系统临时VLAN
				SpeedOs      int    `json:"speed_os"`      // 重装系统端口限速
				ArpOs        int    `json:"arp_os"`        // 重装系统Arp绑定
				CreatedAt    int    `json:"created_at"`
				UpdatedAt    int    `json:"updated_at"`
			} `json:"datacenter"` // 机房
		} `json:"ipv4datacenters"` // IP段绑定机房列表
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/network/ipv4/blocks" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取IP段列表
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5745855
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request, dcimsdk.GidTypeFixed)
}
