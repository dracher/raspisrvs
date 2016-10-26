package routes

import (
	"github.com/kataras/iris"
)

// IndexPage is
func IndexPage(ctx *iris.Context) {
	ctx.MustRender("index.html", struct{}{})
}
