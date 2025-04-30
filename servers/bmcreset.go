package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type BmcResetReq struct {
	ServerId int // 服务器id
}

func (r *BmcResetReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/bmc/reset", r.ServerId)
}
func (r *BmcResetReq) Method() (method string)     { return http.MethodPost }
func (r *BmcResetReq) Values() (values url.Values) { return }
func (r *BmcResetReq) Body() (body any)            { return }

// BmcReset 重置BMC
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5774059
func BmcReset(ctx *dcimsdk.Context, request *BmcResetReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*BmcResetReq, dcimsdk.CreateUpdateResp](ctx, request)
}
