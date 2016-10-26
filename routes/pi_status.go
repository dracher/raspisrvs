package routes

import (
	"github.com/dracher/raspisrvs/services/pistatus"
	"github.com/kataras/iris"
)

// PiStatus is
func PiStatus(ctx *iris.Context) {
	piStatus := ctx.Get("piStatus").(*pistatus.PiStatus)

	res := piStatus.Get()

	ctx.JSON(iris.StatusOK, iris.Map{"resp": res})
}

// PiStatusPage is
func PiStatusPage(ctx *iris.Context) {
	piStatus := ctx.Get("piStatus").(*pistatus.PiStatus)

	res := piStatus.Get()
	ctx.MustRender("pistatus.html", struct{ Res map[string]string }{Res: res})
}
