package ip

import (
	"errors"
	"fmt"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	ListReq struct {
		BlockId  uint `url:"-"`
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	ListResp struct {
		Success   bool      `json:"success"`
		Error     string    `json:"error"`
		Addresses []Address `json:"addresses"`
		Total     int       `json:"total"` // 总条数
	}
	Address struct {
		Id            uint   `json:"id"`             // ip地址id
		BlockId       uint   `json:"block_id"`       // ip段id
		Address       string `json:"address"`        // ip地址
		AddressString string `json:"address_string"` // ip地址
		Status        string `json:"status"`         // 状态
		Type          string `json:"type"`           // 类型
		Relid         int    `json:"relid"`          // 类型关联id
		Mac           string `json:"mac"`            // mac地址
		Remark        string `json:"remark"`         // 备注
		IsLock        int    `json:"is_lock"`        // 锁定状态
		AssignTime    any    `json:"assign_time"`    // 分配时间
		CreatedAt     string `json:"created_at"`     // 分配时间
		UpdatedAt     string `json:"updated_at"`     // 更新时间
		Device        string `json:"device"`         // 驱动
		Vlan          string `json:"vlan"`           // vlan
		Netmask       string `json:"netmask"`        // 子网掩码
		Gateway       string `json:"gateway"`        // 网关
	}
)

func (r *ListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/network/ipv4/blocks/%d/address", r.BlockId)
}
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取ip地址
// https://www.eolink.com/share/inside/XIPzIs/api/1392256/detail/5759697
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request)
}
