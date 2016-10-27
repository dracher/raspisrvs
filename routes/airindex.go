package routes

import (
	"github.com/kataras/iris"
)

// AirIndexPage is
func AirIndexPage(ctx *iris.Context) {
	ctx.MustRender("airindex.html", struct{}{})
}
