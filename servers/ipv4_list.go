package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	Ipv4ListReq  struct{ ServerId uint }
	Ipv4ListResp struct {
		Success bool `json:"success"`
		Ipv4    []struct {
			Id           int    `json:"id"`
			Cidr         string `json:"cidr"`
			Gateway      string `json:"gateway"`
			Netmask      string `json:"netmask"`
			Group        string `json:"group"`
			AssignTime   any    `json:"assign_time"`
			Vlan         string `json:"vlan"`
			IsWholeBlock bool   `json:"is_whole_block"`
			CardName     string `json:"card_name"`
			SwitchId     int    `json:"switch_id"`
			RouteType    any    `json:"route_type"`
			NextHop      any    `json:"next_hop"`
			Ranges       []struct {
				Id            int    `json:"id"`
				BlockId       int    `json:"block_id"`
				AddressString string `json:"address_string"`
			} `json:"ranges"`
		} `json:"ipv4"`
		Total int `json:"total"`
	}
)

func (r *Ipv4ListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/ipv4", r.ServerId)
}
func (r *Ipv4ListReq) Method() (method string)     { return http.MethodGet }
func (r *Ipv4ListReq) Values() (values url.Values) { return }
func (r *Ipv4ListReq) Body() (body any)            { return }

func Ipv4List(ctx *dcimsdk.Context, request *Ipv4ListReq) (resp Ipv4ListResp, err error) {
	return dcimsdk.Execute[*Ipv4ListReq, Ipv4ListResp](ctx, request)
}
