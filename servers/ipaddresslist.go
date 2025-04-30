package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	IpAddressListReq struct {
		BlockId  int  `url:"-"` // ip段id
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	IpAddressListResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Address []struct {
			Id      int    `json:"id"`       // ip地址id
			BlockId int    `json:"block_id"` // ip段id
			Address string `json:"address"`  // ip地址
		} `json:"address"` // ip地址列表
		Total int `json:"total"` // 总条数
	}
)

func (r *IpAddressListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/address/%d", r.BlockId)
}
func (r *IpAddressListReq) Method() (method string)     { return http.MethodGet }
func (r *IpAddressListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *IpAddressListReq) Body() (body any)            { return }

// IpAddressList 获取IP地址列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5782481
func IpAddressList(ctx *dcimsdk.Context, request *IpAddressListReq) (resp IpAddressListResp, err error) {
	return dcimsdk.Execute[*IpAddressListReq, IpAddressListResp](ctx, request)
}
