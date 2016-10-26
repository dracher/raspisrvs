package routes

import (
	"github.com/dracher/raspisrvs/services/airindex"
	"github.com/kataras/iris"
)

// AirIndex is
func AirIndex(ctx *iris.Context) {
	aqiCache := ctx.Get("airIndexCache").(*airindex.AqiData)

	city := ctx.Param("city")

	ctx.Log(city)

	res := aqiCache.CurrentData(city)

	ctx.JSON(iris.StatusOK, iris.Map{"resp": res})
}

// AirIndexPage is
func AirIndexPage(ctx *iris.Context) {
	ctx.MustRender("airindex.html", struct{}{})
}
