package MercerServer

import (
	"MercerFrame/MercerLog"
	"fmt"
	"log"
	"net/http"
)

const space = "  "

type Deal interface {
	NewRouter(HandlerObject)
	MakeRouter()
	DealRouter(http.ResponseWriter, *http.Request)
}

type Request struct {
	D Deal
	//r  *http.Request
	//rw http.ResponseWriter
}

func Default() *Request {
	MercerLog.Info.Println("Welcome to use Mercer's Frame")
	r := &Request{&Router{}}
	r.D.MakeRouter()
	return r
}

func (r *Request) Play(PortNum int) {
	MercerLog.Info.Println("Listening and serving HTTP on :" + fmt.Sprintf("%d", PortNum))
	http.Handle("/", r)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PortNum), r)
	if err != nil {
		log.Fatal(err)
	}
}

func (r *Request) ServeHTTP(w http.ResponseWriter, r2 *http.Request) {
	r.D.DealRouter(w, r2)
}

func (r *Request) NewRouter(ho HandlerObject) {
	r.D.NewRouter(ho)
}

func (r *Request) Get(uri string, f func(*Context)) {
	ho := HandlerObject{
		Uri:         uri,
		Method:      "GET",
		Fun:         f,
		IsWebSocket: false,
	}
	r.D.NewRouter(ho)
	MercerLog.Info.Println(ho.Method + space + ho.Uri)
}

func (r *Request) Post(uri string, f func(*Context)) {
	ho := HandlerObject{
		Uri:         uri,
		Method:      "POST",
		Fun:         f,
		IsWebSocket: false,
	}
	r.D.NewRouter(ho)
	MercerLog.Info.Println(ho.Method + space + ho.Uri)
}

func (r *Request) Put(uri string, f func(*Context)) {
	ho := HandlerObject{
		Uri:         uri,
		Method:      "PUT",
		Fun:         f,
		IsWebSocket: false,
	}
	r.D.NewRouter(ho)
	MercerLog.Info.Println(ho.Method + space + ho.Uri)
}

func (r *Request) WebSocket(uri string, f func(ctx *WebSocketContext)) {
	ho := HandlerObject{
		Uri:         uri,
		Method:      "GET",
		WSFun:       f,
		IsWebSocket: true,
		//SocketFunc:  socketFunc,
	}
	r.D.NewRouter(ho)
	MercerLog.Info.Println(ho.Method + space + ho.Uri)
}

func (r *Request) WebSocket1(uri string, socketFunc WebSocketFunc) {
	ho := HandlerObject{
		Uri:         uri,
		Method:      "GET",
		//WSFun:       f,
		IsWebSocket: true,
		SocketFunc:  socketFunc,
	}
	r.D.NewRouter(ho)
	MercerLog.Info.Println(ho.Method + space + ho.Uri)
}