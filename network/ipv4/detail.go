package ipv4

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	DetailReq struct {
		Id int `json:"-"`
	}

	DetailResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}
)

func (r *DetailReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/network/ipv4/blocks/%d", r.Id)
}
func (r *DetailReq) Method() (method string)     { return http.MethodGet }
func (r *DetailReq) Values() (values url.Values) { return }
func (r *DetailReq) Body() (body any)            { return }

// Detail 获取IP段详情
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5758800
func Detail(ctx *dcimsdk.Context, request *DetailReq) (resp DetailResp, err error) {
	return dcimsdk.Execute[*DetailReq, DetailResp](ctx, request)
}
