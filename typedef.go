package dcimsdk

import (
	"errors"
	"net/url"
	"time"
)

type (
	Request interface {
		Url() (url string)
		Method() (method string)
		Values() (values url.Values)
		Body() (body any)
	}
	CreateUpdateResp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Id      uint   `json:"id"`
	}
	Response interface {
		Ok() (ok bool)
		Err() (err error)
	}
	IdResponse interface {
		Response
		ID() (id uint)
	}
	Option struct {
		timeout     time.Duration
		transformer Transformer
	}
	Transformer func(inBody string) (outBody string)
	OptionFunc  func(opt *Option)
	OnlyIdType  struct {
		Id uint `json:"id"`
	}
)

func (c CreateUpdateResp) ID() uint         { return c.Id }
func (c CreateUpdateResp) Ok() (ok bool)    { return c.Success }
func (c CreateUpdateResp) Err() (err error) { return errors.New(c.Error) }

func (o *Option) Timeout(timeout time.Duration) *Option       { o.timeout = timeout; return o }
func (o *Option) Transformer(transformer Transformer) *Option { o.transformer = transformer; return o }

func NewOnlyIdType(id uint) OnlyIdType { return OnlyIdType{Id: id} }
