package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

// {"up_port":"201","down_port":"201"}

type SpeedReq struct {
	ServerId     uint   `json:"-"`
	SwitchPortId uint   `json:"-"`
	UpPort       string `json:"up_port"`
	DownPort     string `json:"down_port"`
}

func (r *SpeedReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/ports/%d/speed", r.ServerId, r.SwitchPortId)
}
func (r *SpeedReq) Method() (method string)     { return http.MethodPut }
func (r *SpeedReq) Values() (values url.Values) { return }
func (r *SpeedReq) Body() (body any)            { return r }

func Speed(ctx *dcimsdk.Context, request *SpeedReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*SpeedReq, dcimsdk.CreateUpdateResp](ctx, request)
}
