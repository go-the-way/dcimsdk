package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	Ipv4BlockListReq struct {
		ServerId uint `uri:"-"`
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	Ipv4BlockListResp struct {
		Success    bool   `json:"success"`
		Error      string `json:"error"`
		Ipv4Blocks []struct {
			Id          int    `json:"id"`           // ip段id
			Ip          string `json:"ip"`           // id段
			GroupName   string `json:"group_name"`   // ip分组名称
			CanAllocate bool   `json:"can_allocate"` // 是否可分配
			UsableipNum int    `json:"usableip_num"` // 可使用的ip数量
			LockNum     int    `json:"lock_num"`     // 可使用的ip数量
			Percent     int    `json:"percent"`      // 百分比
		} `json:"ipv4blocks"` // ip段列表
		Total int `json:"total"`
	}
)

func (r *Ipv4BlockListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/ipv4lists/%d", r.ServerId)
}
func (r *Ipv4BlockListReq) Method() (method string)     { return http.MethodGet }
func (r *Ipv4BlockListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *Ipv4BlockListReq) Body() (body any)            { return }

// Ipv4BlockList 获取IP段列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5782476
func Ipv4BlockList(ctx *dcimsdk.Context, request *Ipv4BlockListReq) (resp Ipv4ListResp, err error) {
	return dcimsdk.Execute[*Ipv4BlockListReq, Ipv4ListResp](ctx, request)
}
