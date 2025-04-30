package switchs

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type PortUpdateMacReq struct {
	PortId uint   `json:"-"`
	Mac    string `json:"mac"`
}

func (r *PortUpdateMacReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/switchs/ports/%d/mac", r.PortId)
}
func (r *PortUpdateMacReq) Method() (method string)     { return http.MethodPut }
func (r *PortUpdateMacReq) Values() (values url.Values) { return }
func (r *PortUpdateMacReq) Body() (body any)            { return r }

func PortUpdateMac(ctx *dcimsdk.Context, request *PortUpdateMacReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*PortUpdateMacReq, dcimsdk.CreateUpdateResp](ctx, request)
}
