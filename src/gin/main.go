package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	app := gin.Default()

	app.GET("/ws", WebSocketHandler) // 注册websocket路由

	if err := app.Run(":9090"); err != nil {
		panic(err)
	}
}

func WebSocketHandler(c *gin.Context) {
	ws, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		panic(err)
	}

	for {
		mt, p, err := ws.ReadMessage()
		if err != nil {
			break
		}

		// websocket.TextMessage
		fmt.Println("messageType:", mt)
		fmt.Println("p:", string(p))

		c.Writer.Write(p)
	}

	_ = ws.Close()
}
