package hardwares

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type CreateReq struct {
	Name        string `json:"name"`
	Size        uint   `json:"size"`         // 节点大小（U）
	NodeNum     uint   `json:"node_num"`     // 节点数量
	IpmiSupport uint   `json:"ipmi_support"` // ipmi支持
	Card        uint   `json:"card"`         // 网卡数量
	KvmType     string `json:"kvm_type"`     // KVM类型
	Row         uint   `json:"row"`          // 行数
}

func (r *CreateReq) Url() (url string)           { return fmt.Sprintf("/api/admin/hardware") }
func (r *CreateReq) Method() (method string)     { return http.MethodPost }
func (r *CreateReq) Values() (values url.Values) { return }
func (r *CreateReq) Body() (body any)            { return r }

// Create 添加硬件类型
// https://www.eolink.com/share/inside/XIPzIs/api/1392934/detail/5756529
func Create(ctx *dcimsdk.Context, request *CreateReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*CreateReq, dcimsdk.CreateUpdateResp](ctx, request)
}
