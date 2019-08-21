# MercerFramework

Golang轻量级web框架  
  
整体架构：  
1.Mlog包中是日志  
2.Mserver包中的Mrouter是路由的存放和调用等功能  
3.Mserver包中的Mcontext是对请求上下文的处理  
4.Mserver包中的Mmatcher是对参数进行匹配的各种发方法  
5.Mserver包中的Mserver是对接收请求的处理  
6.Mserver包中的Mwebsocket是对websocket功能的实现方法以及上下文  
7.Mserver包中的PraseWebsocket是对websocket连接的解析（有问题，导致连接一直在pending状态，找到后改进）  
  
实现的基本功能：  
 1.取GET的RestFul风格参数，比如/test?name=&age=20，可设置默认值  
 2.取GET动态路由的参数，比如/user/age/:name  
 3.取POST的表单参数，可设置默认值  
 4.websocket的发送和接收消息（ws连接还有些问题，找出问题后以后更改）  
  
代码示例 
  
```Golang
package main

import (
	"MercerFrame/MercerServer"
	"fmt"
	"log"
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
		OnMessage: func(context *MercerServer.WebSocketContext) {
			data, err := context.Read()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(data)
			log.Println(string(data))
			err = context.Send("helloa")
			if err != nil {
				log.Println("send err:" , err)
			}
			log.Println("send data")
		},
		OnClose: func(context *MercerServer.WebSocketContext) {
			context.Conn.Close()
		},
		OnError:nil,
	})
	r.Play(8080)
}

```
