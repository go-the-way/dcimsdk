package dcimsdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"sync"
	"time"
)

var (
	mu         = &sync.Mutex{}
	cacheCtx   = map[string]*Context{}
	cacheCtxID = map[uint]*Context{}
	logger     = log.New(os.Stdout, "[dcimsdk executor] ", log.LstdFlags)
	httpClient = &http.Client{
		Timeout: time.Minute * 2,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
)

func Execute[REQ Request, RESP any](ctx *Context, request REQ, transformer ...BodyTransformer) (resp RESP, err error) {
	debug0 := os.Getenv("DCIM_SDK_DEBUG") == "T"
	var (
		body    io.Reader
		req     *http.Request
		rawResp *http.Response
	)
	if b := request.Body(); b != nil {
		typeOf := reflect.TypeOf(b)
		kind := typeOf.Kind()
		switch kind {
		case reflect.String:
			body = bytes.NewBufferString(b.(string))
			if debug0 {
				logger.Println("execute api body:", b.(string))
			}
		case reflect.Map, reflect.Ptr, reflect.Struct:
			bs, _ := json.Marshal(b)
			body = bytes.NewReader(bs)
			if debug0 {
				logger.Println("execute api body:", string(bs))
			}
		case reflect.Slice, reflect.Array:
			if typeOf.Elem().Kind() == reflect.Uint {
				body = bytes.NewReader(b.([]byte))
				if debug0 {
					logger.Println("execute api body:", string(b.([]byte)))
				}
			}
		}
	}
	if debug0 {
		logger.Println("execute api url:", request.Url())
		logger.Println("execute req method:", request.Method())
	}
	reqUrl := ctx.BaseUrl + request.Url()
	if debug0 {
		logger.Println("execute req url:", reqUrl)
	}
	v := request.Values()
	if v == nil {
		v = make(url.Values)
	}
	v.Set("get_source", "true")
	if vv := v.Encode(); len(vv) > 0 {
		if strings.LastIndexByte(reqUrl, '?') == -1 {
			reqUrl += "?" + vv
		} else {
			reqUrl += "&" + vv
		}
	}
	if debug0 {
		logger.Println("execute req url with values:", reqUrl)
	}
	req, err = http.NewRequest(request.Method(), reqUrl, body)
	if err != nil {
		return
	}
	if len(ctx.token) > 0 {
		req.Header.Set("Access-Token", ctx.token)
	}
	if req.Method != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Admin-Access", "yes")
	if debug0 {
		logger.Printf("execute req headers: %v\n", req.Header)
	}
	if rawResp, err = httpClient.Do(req); err != nil {
		logger.Println("execute api err:", err)
		return
	}
	readAll, err := io.ReadAll(rawResp.Body)
	if err != nil && debug0 {
		fmt.Println(err)
	}
	var bodyTransformer BodyTransformer
	if len(transformer) > 0 {
		bodyTransformer = transformer[0]
	}
	if bodyTransformer != nil {
		readAll = []byte(bodyTransformer(string(readAll)))
	}
	if debug0 {
		readAllStr := string(readAll)
		limit := 1024 * 100
		if len(readAllStr) > limit {
			outputFile := url.PathEscape(reqUrl) + ".json"
			_ = os.WriteFile(outputFile, readAll, 0700)
			fmt.Printf("The dcim sdk response data length over than %d, saved as %s\n", limit, outputFile)
		} else {
			fmt.Println(readAllStr)
		}
	}
	if len(readAll) > 0 {
		err = json.Unmarshal(readAll, &resp)
	}
	return
}
