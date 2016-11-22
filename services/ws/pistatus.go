package ws

import "github.com/kataras/iris"

// PiStatusInit is
func PiStatusInit() {
	var room = "pistatus"
	iris.Websocket.OnConnection(func(c iris.WebsocketConnection) {
		c.Join(room)

		iris.Logger.Println(c.ID())

		c.On("refresh", func(msg string) {
			iris.Logger.Println(msg)
			c.To(room).Emit("refresh", c.ID()+"refreshing")
		})

	})
}
