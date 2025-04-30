package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	RebuildReq struct {
		ServerId         int    `url:"server_id" json:"server_id"`                   // 服务器id（必填）
		ImageId          string `url:"image_id" json:"image_id"`                     // 镜像id（必填）
		FormatSystemOnly bool   `url:"format_system_only" json:"format_system_only"` // 是否只格式化系统
		Password         string `url:"password" json:"password"`                     // 密码（必填）
		RemotePort       int    `url:"remote_port" json:"remote_port"`               // 远程端口
		IsPreinstall     int    `url:"is_preinstall" json:"is_preinstall"`           // 是否预装
		NewHostname      string `url:"new_hostname" json:"new_hostname"`             // 新的主机名
	}

	RebuildResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		TaskUid string `json:"task_uid"` // 任务uid
		Log     string `json:"log"`      // 日志
	}
)

func (r *RebuildReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/rebuild", r.ServerId)
}
func (r *RebuildReq) Method() (method string)     { return http.MethodPost }
func (r *RebuildReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *RebuildReq) Body() (body any)            { return r }

// Rebuild 重装系统
// https://www.eolink.com/share/inside/XIPzIs/api/1398114/detail/5775299
func Rebuild(ctx *dcimsdk.Context, request *RebuildReq) (resp RebuildResp, err error) {
	return dcimsdk.Execute[*RebuildReq, RebuildResp](ctx, request)
}
