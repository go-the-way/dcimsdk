package datacenters

import (
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type CreateReq struct {
	Name string `json:"name"`
}

func (r *CreateReq) Url() (url string)           { return "/api/admin/datacenters" }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

// Create 添加数据中心
// https://www.eolink.com/share/inside/XIPzIs/api/1399365/detail/5784144
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*CreateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
