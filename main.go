package main

import (
	"github.com/dracher/raspisrvs/routes"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
	"github.com/spf13/viper"

	"github.com/dracher/raspisrvs/services/airindex"
	"github.com/dracher/raspisrvs/services/pistatus"
)

var airIndexCache = airindex.NewAqiData()
var piStatus = pistatus.NewPiStatus()

func init() {
	viper.SetConfigFile("./conf.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		iris.Logger.Panicln(err)
	}
}

func main() {

	dev := viper.GetString("dev")
	addr := viper.GetString("addr")

	if dev == "true" {
		iris.Logger.Println("Current in dev mode")
		devConfig()
	} else {
		iris.Logger.Println("Current in prod mode")
		prodConfig()
	}

	iris.UseFunc(allCacheMiddleware)

	registeRouter()
	registeAPI()

	iris.Listen(addr)
}

func allCacheMiddleware(ctx *iris.Context) {
	ctx.Set("airIndexCache", airIndexCache)
	ctx.Set("piStatus", piStatus)
	ctx.Next()
}

func prodConfig() {

	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).
		Directory("./templates", "html").
		Binary(Asset, AssetNames)

	iris.StaticEmbedded("/static", ".assets", Asset, AssetNames)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("404.html",
			iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})
}

func devConfig() {
	iris.Config.IsDevelopment = true
	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).
		Directory("./templates", "html")

	iris.Static("/static", "./assets", 1)

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("404.html",
			iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})
}

func registeRouter() {
	iris.Get("/", routes.IndexPage)
	iris.Get("/aqi", routes.AirIndexPage)
	iris.Get("/pistatus", routes.PiStatusPage)
}

func registeAPI() {
	iris.Get("/airindex/:city", routes.AirIndex)
	// iris.Get("/pistatus", routes.PiStatus)
}
