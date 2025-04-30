package cabinets

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	ListResp struct {
		Success  bool   `json:"success"`
		Error    string `json:"error"`
		Cabinets []struct {
			Id           uint   `json:"id"`
			Name         string `json:"name"`
			DatacenterId uint   `json:"datacenter_id"`
			Capacity     int    `json:"capacity"`
			Electric     string `json:"electric"`
			Remark       string `json:"remark"`
			IsLock       int    `json:"is_lock"`
			Contact      struct {
				Name  string `json:"name"`
				Tel   string `json:"tel"`
				Email string `json:"email"`
			} `json:"contact"`
			CreatedAt  int `json:"created_at"`
			UpdatedAt  int `json:"updated_at"`
			SpaceCount struct {
				Total int `json:"total"`
				Free  int `json:"free"`
				Using int `json:"using"`
				Real  int `json:"real"`
			} `json:"space_count"`
		} `json:"cabinets"`
		Total int `json:"total"`
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/cabinets" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取机柜列表
// https://space-9cdcdq.w.eolink.com/share/inside/XIPzIs/api/1392933/detail/5748523
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request)
}
