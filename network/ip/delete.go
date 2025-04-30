package ip

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type DeleteReq struct {
	BlockId   uint // ip分段id
	AddressId uint // ip地址id
}

func (r *DeleteReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/network/ipv4/blocks/%d/address/delete/%d", r.BlockId, r.AddressId)
}
func (r *DeleteReq) Method() (method string)     { return http.MethodDelete }
func (r *DeleteReq) Values() (values url.Values) { return }
func (r *DeleteReq) Body() (body any)            { return }

// Delete 单个IP地址删除
// https://www.eolink.com/share/inside/XIPzIs/api/1394996/detail/5758437
func Delete(ctx *dcimsdk.Context, request *DeleteReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*DeleteReq, dcimsdk.CreateUpdateResp](ctx, request)
}
