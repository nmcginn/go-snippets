package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func httpHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "%q", ctx.RequestURI())
}

func main() {
	fasthttp.ListenAndServe(":8080", httpHandler)
}

