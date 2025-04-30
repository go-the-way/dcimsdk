package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type TaskStatusReq struct {
	ServerId int // 服务器id
	TaskId   int // 任务id
}

func (r *TaskStatusReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/task/%d", r.ServerId, r.TaskId)
}
func (r *TaskStatusReq) Method() (method string)     { return http.MethodGet }
func (r *TaskStatusReq) Values() (values url.Values) { return }
func (r *TaskStatusReq) Body() (body any)            { return }

// TaskStatus 服务器重装进度
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5779893
func TaskStatus(ctx *dcimsdk.Context, request *TaskStatusReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*TaskStatusReq, dcimsdk.CreateUpdateResp](ctx, request)
}
