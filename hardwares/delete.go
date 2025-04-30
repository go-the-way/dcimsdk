package hardwares

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type DeleteReq struct{ Id uint }

func (r *DeleteReq) Url() (url string)           { return fmt.Sprintf("/api/admin/hardware/%d", r.Id) }
func (r *DeleteReq) Method() (method string)     { return http.MethodDelete }
func (r *DeleteReq) Values() (values url.Values) { return }
func (r *DeleteReq) Body() (body any)            { return }

// Delete 删除数据中心
// https://www.eolink.com/share/inside/XIPzIs/api/1392934/detail/5756619
func Delete(ctx *dcimsdk.Context, request *DeleteReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*DeleteReq, dcimsdk.CreateUpdateResp](ctx, request)
}
