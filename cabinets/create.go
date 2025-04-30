package cabinets

import (
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"
)

type (
	CreateReq struct {
		CreatedAt      int64  `json:"created_at"`
		Name           string `json:"name"`
		DatacenterId   uint   `json:"datacenter_id"`
		DatacenterName string `json:"datacenter_name"`
		Id             int    `json:"id"`
		IsLock         int    `json:"is_lock"`
		Capacity       uint   `json:"capacity"`
		Contact        struct {
			Name  string `json:"name"`
			Tel   string `json:"tel"`
			Email string `json:"email"`
		} `json:"contact"`
		Electric     string `json:"electric"`
		ElectricUnit string `json:"electric_unit"`
		Remark       string `json:"remark"`
		Edited       bool   `json:"edited"`
	}
)

func (r *CreateReq) Url() (url string)           { return "/api/admin/cabinets" }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

// Create 添加机柜
// https://space-9cdcdq.w.eolink.com/share/inside/XIPzIs/api/1392933/detail/5748524
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*CreateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
