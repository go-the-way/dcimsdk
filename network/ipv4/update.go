package ipv4

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type UpdateReq struct {
	Id           uint   `json:"-"`
	DatacenterId []uint `json:"datacenter_id"`     // 机房Ids
	Gid          uint   `json:"gid,omitempty"`     // 分组Id
	Type         string `json:"type,omitempty"`    // ip段类别 IPMI:IPMI;Temporary:临时IP;Public:公网IP;Private:内网IP;UserPrivate:客户内网IP;Other:其他
	Dns1         string `json:"dns1,omitempty"`    // 首选DNS
	Dns2         string `json:"dns2,omitempty"`    // 备用DNS
	Gateway      string `json:"gateway,omitempty"` // 网关
	Label        string `json:"label,omitempty"`   // 标签（页面不填写，后端调用接口传固定值）
	Remark       string `json:"remark,omitempty"`  // 备注
	//vlan:0
	//auto:1
	// is_lock:0
	// tag:
}

func (r *UpdateReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/network/ipv4/blocks/%d", r.Id)
}
func (r *UpdateReq) Method() (method string)     { return http.MethodPut }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Body() (body any)            { return r }

// Update 修改IP段
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5748038
func Update(ctx *dcimsdk.Context, request *UpdateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*UpdateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
