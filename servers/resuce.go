package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ResuceReq struct {
		ServerId int    `url:"server_id" json:"server_id"` // 服务器id
		Type     string `url:"type" json:"type"`           // 类型win(默认);linux
	}

	ResuceResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		TaskUid string `json:"task_uid"` // 任务uid
		Log     string `json:"log"`      // 日志
	}
)

func (r *ResuceReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/resuce", r.ServerId)
}
func (r *ResuceReq) Method() (method string)     { return http.MethodPost }
func (r *ResuceReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ResuceReq) Body() (body any)            { return r }

// Resuce 救援模式
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5774095
func Resuce(ctx *dcimsdk.Context, request *ResuceReq) (resp RebuildResp, err error) {
	return dcimsdk.Execute[*ResuceReq, RebuildResp](ctx, request)
}
