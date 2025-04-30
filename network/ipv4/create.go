package ipv4

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	CreateReq struct {
		DatacenterId []uint `json:"datacenter_id"` // 机房Ids
		Gid          uint   `json:"gid"`           // 分组Id
		Cidr         string `json:"cidr"`          // IPV4段
		FirstIp      string `json:"first_ip"`      // 第一个IP
		LastIp       string `json:"last_ip"`       // 最后一个IP
		Type         string `json:"type"`          // ip段类别 IPMI:IPMI;Temporary:临时IP;Public:公网IP;Private:内网IP;UserPrivate:客户内网IP;Other:其他
		Netmask      string `json:"netmask"`       // 子网掩码
		Dns1         string `json:"dns1"`          // 首选DNS
		Dns2         string `json:"dns2"`          // 备用DNS
		Gateway      string `json:"gateway"`       // 网关
		Label        string `json:"label"`         // 标签（页面不填写，后端调用接口传固定值）
		Remark       string `json:"remark"`        // 备注
		/*	auto:自动分配
			vlan:VLAN
			auto_split:自动拆分
			is_lock:是否锁定
			tag:自定义标签*/
	}
	CreateResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		BlockId uint   `json:"block_id"`
		Pending bool   `json:"pending"` // 是否等待
	}
)

func (r *CreateReq) Url() (url string)           { return "/api/admin/network/ipv4/blocks" }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

func (c CreateResp) ID() uint         { return c.BlockId }
func (c CreateResp) Ok() (ok bool)    { return c.Success }
func (c CreateResp) Err() (err error) { return errors.New(c.Error) }

// Create 创建IP段
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5745856
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp CreateResp, err error) {
	return dcimsdk.Execute[*CreateReq, CreateResp](ctx, request)
}
