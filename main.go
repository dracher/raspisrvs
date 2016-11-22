package main

import (
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
	"github.com/spf13/viper"

	"github.com/dracher/raspisrvs/routes"
	"github.com/dracher/raspisrvs/routes/api"
	"github.com/dracher/raspisrvs/services/airindex"
	"github.com/dracher/raspisrvs/services/pistatus"

	"github.com/dracher/raspisrvs/services/ws"
)

var airIndexCache *airindex.AqiData
var piStatus *pistatus.PiStatus

func init() {
	viper.SetConfigFile("./conf.yml")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		iris.Logger.Panicln(err)
	}
}

func main() {

	dev := viper.GetBool("dev")
	addr := viper.GetString("addr")

	if dev {
		iris.Logger.Println("Current in dev mode")
		devConfig()
	} else {
		iris.Logger.Println("Current in prod mode")
		prodConfig()
	}

	wsConfig()

	airIndexCache = airindex.NewAqiData(dev)
	piStatus = pistatus.NewPiStatus()

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

func wsConfig() {
	iris.Config.Websocket.Endpoint = "/ws"
	ws.PiStatusInit()
}

func prodConfig() {

	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).
		Directory("./templates", "html").
		Binary(Asset, AssetNames)

	iris.StaticEmbedded("/static", ".assets", Asset, AssetNames)
	// iris.Favicon("/static/favicon.ico")

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
	// iris.Favicon("./assets/favicon.ico")

	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.Render("404.html",
			iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
	})
}

func registeRouter() {
	iris.Get("/", routes.IndexPage)
	iris.Get("/srvs", routes.ServicesPage)
	iris.Get("/aqi", routes.AirIndexPage)
	iris.Get("/pistatus", routes.PiStatusPage)
}

func registeAPI() {
	apiv1 := iris.Party("/api/v1")

	apiv1.Get("/airindex/:city", api.AirIndex)
	apiv1.Get("/pistatus", api.PiStatus)
}
