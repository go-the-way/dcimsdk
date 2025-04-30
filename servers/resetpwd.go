package servers

import (
	"fmt"
	"github.com/go-the-way/dcimsdk"
	"net/http"
	"net/url"
)

type (
	ResetPwdReq struct {
		ServerId int    `json:"-"`        // 服务器id
		Password string `json:"password"` // 密码
	}
)

func (r *ResetPwdReq) Url() (url string) {
	return fmt.Sprintf("/api/admin/servers/%d/password/reset", r.ServerId)
}
func (r *ResetPwdReq) Method() (method string)     { return http.MethodPost }
func (r *ResetPwdReq) Values() (values url.Values) { return }
func (r *ResetPwdReq) Body() (body any)            { return r }

// ResetPwd 服务器重置系统管理员密码
// https://www.eolink.com/share/inside/XIPzIs/api/1398114/detail/5775337
func ResetPwd(ctx *dcimsdk.Context, request *ResetPwdReq) (resp dcimsdk.CreateUpdateResp, err error) {
	return dcimsdk.Execute[*ResetPwdReq, dcimsdk.CreateUpdateResp](ctx, request)
}
