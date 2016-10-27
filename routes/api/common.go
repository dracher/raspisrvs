package api

import (
	"github.com/dracher/raspisrvs/services/airindex"
	"github.com/dracher/raspisrvs/services/pistatus"
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

// PiStatus is
func PiStatus(ctx *iris.Context) {
	piStatus := ctx.Get("piStatus").(*pistatus.PiStatus)

	res := piStatus.Get()

	ctx.JSON(iris.StatusOK, iris.Map{"resp": res})
}
