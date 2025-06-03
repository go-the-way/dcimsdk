package dcimsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func init() { go startRefreshToken() }

func startRefreshToken() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			refreshToken()
		}
	}
}

func refreshToken() {
	if !mu.TryLock() {
		return
	}
	defer mu.Unlock()
	for _, v := range cacheCtx {
		v.refreshToken()
	}
	return
}

func (ctx *Context) refreshToken() {
	logger.Println("start refresh token")
	req, _ := http.NewRequest(http.MethodPost, ctx.BaseUrl+"/api/admin/login", bytes.NewBufferString(fmt.Sprintf("{\"key\":\"%s\",\"secret\":\"%s\"}", ctx.Key, ctx.Secret)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Admin-Access", "yes")
	if resp, err := newHttpClient(time.Second * 5).Do(req); err != nil {
		logger.Println("get token err:", err)
	} else if resp.StatusCode != http.StatusOK {
		logger.Println("get token not 200")
	} else {
		if readAll, err0 := io.ReadAll(resp.Body); err0 != nil {
			logger.Println("get token read err:", err0)
		} else {
			var resp0 loginResp
			if err0 = json.Unmarshal(readAll, &resp0); err0 != nil {
				logger.Println("get token json unmarshal err:", err0)
			} else {
				ctx.token = resp0.Token
				logger.Println("get token:", ctx.BaseUrl, ctx.token)
			}
		}
	}
	logger.Println("end refresh token")
	return
}

type loginResp struct {
	Success     bool   `json:"success"`
	Token       string `json:"token"`
	Ttl         string `json:"ttl"`
	Permissions []struct {
		Id     string `json:"id"`
		Method string `json:"method"`
		Uri    string `json:"uri"`
	} `json:"permissions"`
}
