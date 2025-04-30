package datacenters

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type UpdateReq struct {
	Id   uint   `json:"-"`
	Name string `json:"name"`
}

func (r *UpdateReq) Url() (url string)           { return fmt.Sprintf("/api/admin/datacenters/%d", r.Id) }
func (r *UpdateReq) Method() (method string)     { return http.MethodPut }
func (r *UpdateReq) Values() (values url.Values) { return }
func (r *UpdateReq) Body() (body any)            { return r }

// Update 编辑数据中心
// https://www.eolink.com/share/inside/XIPzIs/api/1399365/detail/5784204
func Update(ctx *dcimsdk.Context, request *UpdateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*UpdateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
