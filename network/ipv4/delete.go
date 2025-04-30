package ipv4

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type DeleteReq struct {
	Id uint `json:"-"`
}

func (r *DeleteReq) Url() (url string)           { return fmt.Sprintf("/api/admin/network/ipv4/blocks/%d", r.Id) }
func (r *DeleteReq) Method() (method string)     { return http.MethodDelete }
func (r *DeleteReq) Values() (values url.Values) { return }
func (r *DeleteReq) Body() (body any)            { return }

// Delete 删除数据中心
// https://www.eolink.com/share/inside/XIPzIs/api/1399365/detail/5784268
func Delete(ctx *dcimsdk.Context, request *DeleteReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*DeleteReq, dcimsdk.CreateUpdateResp](ctx, request)
}
