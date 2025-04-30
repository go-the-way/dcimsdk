package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type DeleteReq struct {
	Id uint // 服务器id
}

func (r *DeleteReq) Url() (url string)           { return fmt.Sprintf("/api/admin/servers/%d", r.Id) }
func (r *DeleteReq) Method() (method string)     { return http.MethodDelete }
func (r *DeleteReq) Values() (values url.Values) { return }
func (r *DeleteReq) Body() (body any)            { return }

// Delete 删除服务器
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770196
func Delete(ctx *dcimsdk.Context, request *DeleteReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*DeleteReq, dcimsdk.CreateUpdateResp](ctx, request)
}
