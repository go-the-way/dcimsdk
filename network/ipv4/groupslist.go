package ipv4

import (
	"errors"
	"github.com/go-the-way/dcimsdk"
	q "github.com/google/go-querystring/query"
	"net/http"
	"net/url"
)

type (
	GroupsListReq  struct{}
	GroupsListResp struct {
		Success bool          `json:"success"`
		Error   string        `json:"error"`
		Groups  []SimpleGroup `json:"groups"`
	}
	SimpleGroup struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
)

func (r *GroupsListReq) Url() (url string)           { return "/api/admin/network/ipv4/groups/list" }
func (r *GroupsListReq) Method() (method string)     { return http.MethodGet }
func (r *GroupsListReq) Values() (values url.Values) { values, _ = q.Values(r); return }
func (r *GroupsListReq) Body() (body any)            { return }

func (r GroupsListResp) Ok() (ok bool)    { return r.Success }
func (r GroupsListResp) Err() (err error) { return errors.New(r.Error) }

// GroupsList 获取IP分组列表(简易)
// https://www.eolink.com/share/inside/XIPzIs/api/1392254/detail/5745551
func GroupsList(ctx *dcimsdk.Context, request *GroupsListReq) (resp GroupsListResp, err error) {
	return dcimsdk.Execute[*GroupsListReq, GroupsListResp](ctx, request)
}
