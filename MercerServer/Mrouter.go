package MercerServer

import (
	"MercerFrame/MercerLog"
	"net"
	"net/http"
	"strings"
	"unsafe"
)

//请求实体
type HandlerObject struct {
	Header      http.Header
	Uri         string
	Method      string
	Fun         func(cont *Context)
	WSFun       func(cont *WebSocketContext)
	ParamNum    int
	ParamName   []string
	IsWebSocket bool
	SocketFunc  WebSocketFunc
	Req         *http.Request
	RW          http.ResponseWriter
	RNum        int
	WS          WebSocketFunc
	conn        net.Conn
}

//type WSParam = map[string]interface{}

//路由的结构体
type Router struct {
	RouterMap *map[string]HandlerObject

	num int
	//临时测试websocket
	WSRouter *map[string]WebSocketFunc
}

//创建新的路由
func (r *Router) MakeRouter() {
	routermap := make(map[string]HandlerObject)
	r.RouterMap = &routermap

	//临时测试websocket
	wsrouter := make(map[string]WebSocketFunc)
	r.WSRouter = &wsrouter
}

//添加路由
func (r *Router) NewRouter(ho HandlerObject) {
	router := *r.RouterMap
	wsrouter := *r.WSRouter

	//测试
	if strings.Count(ho.Uri, "/") > 1 {
		temp := strings.Split(ho.Uri, "/:")[0]
		r.num = strings.Count(temp, "/")
		ho.RNum = r.num
	}

	//如果路径中有":"，则只存":"之前的部分进路由
	if strings.Contains(ho.Uri, ":") {
		temp := strings.Split(ho.Uri, "/:")[1]
		ho.ParamName = strings.SplitN(temp, "/", -1)
		//ho.ParamName = strings.SplitN(ho.Uri, "/", -1)[2:]
		t1 := "/" + temp
		ho.ParamNum = strings.Count(t1, "/")
		//ho.ParamNum = strings.Count(ho.Uri, "/") - 2
		uri := strings.Split(ho.Uri, "/:")[0] //uri = "/user"
		ho.Uri = uri
	}

	_, ok := router[ho.Uri]
	if ok {
		panic("The router has existed")
	}
	router[ho.Uri] = ho

	//websocket临时测试
	if ho.IsWebSocket {
		wsrouter[ho.Uri] = ho.SocketFunc
	}
}

//查找路由里的方法，有就执行，没有就判断错误，404，405
func (r *Router) DealRouter(w http.ResponseWriter, r2 *http.Request) {
	uri := strings.Split(r2.RequestURI, "?")[0]
	route := *r.RouterMap
	//测试
	//wsroute := *r.WSRouter

	if strings.Count(uri, "/") > 1 {
		temp := strings.SplitN(uri, "/", -1)
		uri1 := ""
		for i := 1; i <= r.num; i++ {
			uri1 += ("/" + temp[i])
		}
		uri = uri1
	}

	ho, ok := route[uri]
	if ok {
		if ho.Method == r2.Method {
			if ho.IsWebSocket {
				ho.IntoWebsocket(w, r2, ho.SocketFunc)
				ho.DealWSContext(w, r2)
				return
			}
			ho.DealContext(w, r2)
			MercerLog.Run.Println("     " + r2.Method + "     " + uri)
		} else {
			DealBy405(w, r2)
		}
	} else {
		DealBy404(w, r2)
	}
}

func DealBy404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte ("404 not found"))
}

func DealBy405(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("405 error method"))
}

func QuickBytesToString(b []byte) (s string) {
	return *(*string)(unsafe.Pointer(&b))
}

func QuickStringToBytes(s string) (b []byte) {
	return *(*[]byte)(unsafe.Pointer(&s))
}
