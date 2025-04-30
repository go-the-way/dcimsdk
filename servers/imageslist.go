package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	ImagesListReq struct {
		ServerId uint // 服务器id
	}
	ImagesListResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
	}
)

func (r *ImagesListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/rebuild/images", r.ServerId)
}
func (r *ImagesListReq) Method() (method string)     { return http.MethodGet }
func (r *ImagesListReq) Values() (values url.Values) { return }
func (r *ImagesListReq) Body() (body any)            { return }

// ImagesList 获取重装系统镜像列表
// https://www.eolink.com/share/inside/XIPzIs/api/1398114/detail/5775256
func ImagesList(ctx *dcimsdk.Context, request *ImagesListReq) (resp ImagesListResp, err error) {
	return dcimsdk.Execute[*ImagesListReq, ImagesListResp](ctx, request)
}
