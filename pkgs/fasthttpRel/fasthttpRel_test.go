package fasthttpRel

import (
	"github.com/buaazp/fasthttprouter"
	json "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
	"log"
	"testing"
)

func defaultHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNoContent)
}

func onReportV1(ctx *fasthttp.RequestCtx) {
	content := ctx.PostBody()
	fields := make(map[string]interface{})
	if err := json.Unmarshal(content, &fields); err != nil {
		panic("invalid content")
	}
	log.Printf("%v", fields)
	ctx.SetStatusCode(fasthttp.StatusOK)

	// 处理 HTTP 响应数据
	ctx.Response.Header.SetStatusCode(200)
	ctx.Response.Header.SetConnectionClose() // 关闭本次连接, 这就是短连接 HTTP
	ctx.Response.Header.SetBytesKV([]byte("Content-Type"), []byte("text/plain; charset=utf8"))
	ctx.Response.SetBody(content)
	ctx.Response.SetStatusCode(fasthttp.StatusOK)
}

func Test01(t *testing.T) {
	listener, err := reuseport.Listen("tcp4", "0.0.0.0:1234")
	if err != nil {
		log.Fatal(err)
	}
	router := fasthttprouter.New()
	router.POST("/v1/report", onReportV1)
	router.NotFound = defaultHandler
	fasthttp.Serve(listener, router.Handler)
}
/*
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"username":"xyz","password":"xyz"}' \
  http://localhost:1234//v1/report
*/