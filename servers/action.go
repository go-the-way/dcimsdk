package servers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"
)

type ActionReq struct {
	ServerId uint   `json:"-"`      // 服务器id（必填）
	Action   string `json:"action"` // on(默认值) off reset
}

func (r *ActionReq) Url() (url string)           { return fmt.Sprintf("/api/admin/servers/%d/power", r.ServerId) }
func (r *ActionReq) Method() (method string)     { return http.MethodPost }
func (r *ActionReq) Values() (values url.Values) { return }
func (r *ActionReq) Body() (body any)            { return r }

// action 操作
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5773835
func action(ctx *dcimsdk.Context, request *ActionReq, opts ...dcimsdk.OptionFunc) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*ActionReq, dcimsdk.CreateUpdateResp](ctx, request, opts...)
}

// ActionPowerOn 开机
func ActionPowerOn(ctx *dcimsdk.Context, serverId uint) (resp dcimsdk.CreateUpdateResp, err error) {
	return action(ctx, &ActionReq{ServerId: serverId, Action: "on"})
}

// ActionPowerOff 关机
func ActionPowerOff(ctx *dcimsdk.Context, serverId uint) (resp dcimsdk.CreateUpdateResp, err error) {
	return action(ctx, &ActionReq{ServerId: serverId, Action: "off"})
}

// ActionPowerReset 重启
func ActionPowerReset(ctx *dcimsdk.Context, serverId uint) (resp dcimsdk.CreateUpdateResp, err error) {
	return action(ctx, &ActionReq{ServerId: serverId, Action: "reset"})
}
