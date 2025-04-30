package cabinets

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type UpdateReq struct {
	Id           uint   `json:"-"`
	Name         string `json:"name"`          // 名称
	DatacenterId uint   `json:"datacenter_id"` // 机房id
	Electric     int    `json:"electric"`      // 电量
	Remark       string `json:"remark"`        // 备注
	IsLock       int    `json:"is_lock"`
	Contact      struct {
		Name  string `json:"name"`  // 联系人
		Tel   string `json:"tel"`   // 联系电话
		Email string `json:"email"` // 联系人邮箱
	} `json:"contact"`
}

func (r *UpdateReq) Url() (url string)           { return fmt.Sprintf("/api/admin/cabinets/%d", r.Id) }
func (r *UpdateReq) Method() (method string)     { return http.MethodPut }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Body() (body any)            { return r }

// Update 修改机柜
// https://space-9cdcdq.w.eolink.com/share/inside/XIPzIs/api/1392933/detail/5748522
func Update(ctx *dcimsdk.Context, request *UpdateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*UpdateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
