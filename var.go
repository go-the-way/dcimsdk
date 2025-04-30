package dcimsdk

import "strings"

var replaceAllFunc = func(s, old, new0 string) string { return strings.ReplaceAll(s, old, new0) }
var IdTypeFixed BodyTransformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `_id":""`, `_id":0`) }
var DeviceTypeFixed BodyTransformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"device":[]`, `"device":{}`) }
var PowerStatusFixed BodyTransformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"power":false`, `"power":"off"`) }
var GidTypeFixed BodyTransformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"gid":""`, `"gid":0`) }

type defaultResponse struct{}

func (d *defaultResponse) Ok() (ok bool)    { return }
func (d *defaultResponse) Err() (err error) { return }
func (d *defaultResponse) ID() (id uint)    { return }

func OptionalId(resp IdResponse) (id uint) {
	if resp != nil {
		return resp.ID()
	}
	return
}
