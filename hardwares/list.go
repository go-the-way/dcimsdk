package hardwares

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-the-way/dcimsdk"

	q "github.com/google/go-querystring/query"
)

type (
	ListReq struct {
		Page     uint `url:"page"`
		PageSize uint `url:"page_size"`
	}
	ListResp struct {
		Success  bool   `json:"success"`
		Error    string `json:"error"`
		Hardware []struct {
			Id          uint     `json:"id"`
			Name        string   `json:"name"`
			KvmType     string   `json:"kvm_type"`
			Size        int      `json:"size"`
			NodeNum     int      `json:"node_num"`
			Row         int      `json:"row"`
			MacOffset   int      `json:"mac_offset"`
			IpmiShFile  string   `json:"ipmi_sh_file"`
			Remarks     string   `json:"remarks"`
			Img1        string   `json:"img1"`
			Img2        string   `json:"img2"`
			Features    string   `json:"features"`
			IpmiSupport int      `json:"ipmi_support"`
			Card        int      `json:"card"`
			NetworkCard []string `json:"network_card"`
			CreatedAt   string   `json:"created_at"`
			UpdatedAt   string   `json:"updated_at"`
		} `json:"hardware"`
		Total int `json:"total"`
	}
)

func (r *ListReq) Url() (url string)           { return "/api/admin/hardware" }
func (r *ListReq) Method() (method string)     { return http.MethodGet }
func (r *ListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *ListReq) Body() (body any)            { return }

func (r ListResp) Ok() (ok bool)    { return r.Success }
func (r ListResp) Err() (err error) { return errors.New(r.Error) }

// List 获取硬件类型列表
// https://www.eolink.com/share/inside/XIPzIs/api/1392934/detail/5756496
func List(ctx *dcimsdk.Context, request *ListReq) (resp ListResp, err error) {
	return dcimsdk.Execute[*ListReq, ListResp](ctx, request)
}
