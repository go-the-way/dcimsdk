package ipv4

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	CalcReq struct {
		Block string `json:"block" url:"block"` // IP段
	}

	CalcResp struct {
		Success  bool     `json:"success"`
		Error    string   `json:"error"`
		CidrInfo CidrInfo `json:"cidr_info"`
	}
	CidrInfo struct {
		FirstIp        string `json:"first_ip"`         // 首位ip
		LastIp         string `json:"last_ip"`          // 末尾ip
		NetworkIp      string `json:"network_ip"`       // 网络ip
		BroadcastIp    string `json:"broadcast_ip"`     // 广播
		Cidr           string `json:"cidr"`             // CIDR(ipv4段)
		Netmask        string `json:"netmask"`          // 子网掩码
		Total          int    `json:"total"`            // 总条数
		GatewayFirstIp string `json:"gateway_first_ip"` // 网关首位ip
		GatewayLastIp  string `json:"gateway_last_ip"`  // 网关末尾ip
	}
)

func (r *CalcReq) Url() (url string)           { return "/api/admin/network/ipv4/blocks/calc" }
func (r *CalcReq) Method() (method string)     { return http.MethodGet }
func (r *CalcReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *CalcReq) Body() (body any)            { return }

func (r CalcResp) Ok() (ok bool)    { return r.Success }
func (r CalcResp) Err() (err error) { return errors.New(r.Error) }

// Calc IP段计算器
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5747435
func Calc(ctx *dcimsdk.Context, request *CalcReq) (resp CalcResp, err error) {
	return dcimsdk.Execute[*CalcReq, CalcResp](ctx, request)
}
