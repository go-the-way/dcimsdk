package servers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"
)

type (
	PartListReq  struct{ ServerId int }
	PartListResp struct {
		Success     bool   `json:"success"`
		Error       string `json:"error"`
		ServerParts []struct {
			Id           int    `json:"id"`            // 硬件id
			TypeId       int    `json:"type_id"`       // 硬件类型id
			Brand        string `json:"brand"`         // 品牌
			Supplier     string `json:"supplier"`      // 供应商
			OfficalModel string `json:"officalmodel"`  // 型号
			Price        string `json:"price"`         // 价格
			DatacenterId int    `json:"datacenter_id"` // 数据中心id
			Qty          int    `json:"qty"`           // 数量
			Customfield0 string `json:"customfield0"`  // 自定义字段1
			Customfield1 string `json:"customfield1"`  // 自定义字段2
			Customfield2 string `json:"customfield2"`  // 自定义字段3
			Customfield3 string `json:"customfield3"`  // 自定义字段4
			PurchaseTime string `json:"purchase_time"` // 购买时间
			Remark       string `json:"remark"`        // 备注
			CreatedAt    int    `json:"created_at"`
			UpdatedAt    int    `json:"updated_at"`
			PartId       int    `json:"part_id"`   // 硬件组件id
			PartType     string `json:"part_type"` // 硬件组件类型
			Status       string `json:"status"`    // 硬件组件状态
		} `json:"server_parts"` // 服务器硬件列表
	}
)

func (r *PartListReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/parts", r.ServerId)
}
func (r *PartListReq) Method() (method string)     { return http.MethodGet }
func (r *PartListReq) Values() (values url.Values) { return }
func (r *PartListReq) Body() (body any)            { return }

// PartList 获取服务器硬件列表
// https://www.eolink.com/share/inside/XIPzIs/api/1389881/detail/5770400
func PartList(ctx *dcimsdk.Context, request *PartListReq) (resp PartListResp, err error) {
	return dcimsdk.Execute[*PartListReq, PartListResp](ctx, request)
}
