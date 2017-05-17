package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func httpHandler(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Hello, World!")
}

func main() {
	fmt.Println("HTTP server starting on port 8080.")
	fasthttp.ListenAndServe(":8080", httpHandler)
}

