package dcimsdk

import (
	"strings"
	"time"
)

var (
	replaceAllFunc                    = func(s, old, new0 string) string { return strings.ReplaceAll(s, old, new0) }
	idTypeFixed           Transformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `_id":""`, `_id":0`) }
	deviceTypeFixed       Transformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"device":[]`, `"device":{}`) }
	powerStatusFixed      Transformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"power":false`, `"power":"off"`) }
	gidTypeFixed          Transformer = func(inBody string) (outBody string) { return replaceAllFunc(inBody, `"gid":""`, `"gid":0`) }
	IdTypeFixedOptFn                  = func(opt *Option) { opt.Transformer(idTypeFixed) }
	DeviceTypeFixedOptFn              = func(opt *Option) { opt.Transformer(deviceTypeFixed) }
	PowerStatusFixedOptFn             = func(opt *Option) { opt.Transformer(powerStatusFixed) }
	GidTypeFixedOptFn                 = func(opt *Option) { opt.Transformer(gidTypeFixed) }
)

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

func Timeout(timeout time.Duration) OptionFunc { return func(opt *Option) { opt.timeout = timeout } }

func JoinOptFunc(opt OptionFunc, opts ...OptionFunc) []OptionFunc {
	if opt != nil {
		opts = append(opts, opt)
	}
	return opts
}
