package dcimsdk

import "os"

type Context struct {
	ID      uint
	BaseUrl string
	Key     string
	Secret  string
	token   string
}

func AddCtx(ctx ...*Context) {
	for _, c := range ctx {
		if debugToken := os.Getenv("DCIM_DEBUG_TOKEN"); len(debugToken) > 0 {
			c.token = debugToken
		} else {
			c.refreshToken()
		}
		cacheCtx[c.BaseUrl] = c
		cacheCtxID[c.ID] = c
	}
}

func GetCtxByUrl(baseUrl string) *Context {
	mu.Lock()
	defer mu.Unlock()
	return cacheCtx[baseUrl]
}

func GetCtxByID(id uint) *Context {
	mu.Lock()
	defer mu.Unlock()
	return cacheCtxID[id]
}

func UpCtxByID(ctx *Context) {
	mu.Lock()
	defer mu.Unlock()
	if ctx.ID == 0 {
		return
	}
	cCtx := cacheCtxID[ctx.ID]
	if cCtx == nil {
		return
	}
	if ctx.BaseUrl != cCtx.BaseUrl {
		delete(cacheCtx, cCtx.BaseUrl)
	}
	cacheCtx[ctx.BaseUrl] = ctx
	cacheCtxID[ctx.ID] = ctx
}

func DelCtxByID(id uint) {
	mu.Lock()
	defer mu.Unlock()
	ctx := cacheCtxID[id]
	if ctx == nil {
		return
	}
	delete(cacheCtx, ctx.BaseUrl)
	delete(cacheCtxID, ctx.ID)
}
