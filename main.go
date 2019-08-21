package main

import (
	"MercerFrame/MercerServer"
	"fmt"
)

func main() {
	r := MercerServer.Default()
	r.Get("/test", func(context *MercerServer.Context) {
		context.WriteOnWeb("hello")
		name := context.DefaultQuery("name", "no one")
		age := context.Query("age")
		fmt.Println("name = " + name + "age = " + age)
	})

	r.Get("/user/age/:name", func(context *MercerServer.Context) {
		context.WriteOnWeb("get param")
		name := context.PraParam("name")
		fmt.Println("name = "+name)
	})

	r.Post("/test2", func(context *MercerServer.Context) {
		context.WriteOnWeb("取post参数")
		message := context.PostForm("message")
		name := context.DefaultPostForm("name","no one")
		fmt.Println(message,name)
	})

	r.WebSocket1("/ws",MercerServer.WebSocketFunc{
		OnOpen: func(context *MercerServer.WebSocketContext) {
			fmt.Println("open")
		},
		//OnMessage: func(context *MercerServer.WebSocketContext) {
		//	data, err := context.Read()
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Println(data)
		//	log.Println(string(data))
		//	err = context.Send("helloa")
		//	if err != nil {
		//		log.Println("send err:" , err)
		//	}
		//	log.Println("send data")
		//},
		OnClose: func(context *MercerServer.WebSocketContext) {
			context.Conn.Close()
		},
		OnError:nil,
	})

	//r.WebSocket("/ws", func(ctx *MercerServer.WebSocketContext) {
	//	ctx.OnOpen()
	//})
	r.Play(8080)
}
