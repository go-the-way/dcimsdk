package datacenters

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	ListResp struct {
		Success     bool   `json:"success"`
		Error       string `json:"error"`
		Datacenters []struct {
			Id          uint   `json:"id"`          // id
			Name        string `json:"name"`        // 名字
			Geographicg string `json:"geographicg"` // 地理位置
			Bandwidth   int    `json:"bandwidth"`   // 带宽
			Power       int    `json:"power"`       // 电力
			Contact     struct {
				Name  string `json:"name"`  // 联系人名字
				Tel   string `json:"tel"`   // 联系电话
				Email string `json:"email"` // 联系邮箱
			} `json:"contact"` // 联系信息
			InternalCidr string `json:"internal_cidr"`
			VlanOs       int    `json:"vlan_os"`
			SpeedOs      int    `json:"speed_os"`
			ArpOs        int    `json:"arp_os"`
			CreatedAt    int    `json:"created_at"`
			UpdatedAt    int    `json:"updated_at"`
			ServerStock  struct {
				Total   int `json:"total"`    // 总库存
				Rent    int `json:"rent"`     // 租赁
				Hosting int `json:"hosting"`  // 托管
				SelfUse int `json:"self-use"` // 自用
			} `json:"server_stock"` // 服务器库存
		} `json:"datacenters"` // 数据中心列表
		Total int `json:"total"`
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/datacenters" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取机房列表
// https://www.eolink.com/share/inside/XIPzIs/api/1399365/detail/5783883
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request)
}
