package MercerServer

import (
	"fmt"
	"net/http"
)

var fun WebSocketFunc

type Context struct {
	Header    http.Header
	Req       *http.Request
	RW        http.ResponseWriter
	Param     map[string]string
	Uri       string
	RouterUri string
	ParamNum  int
	ParamName []string
	RNum      int
}

func (ho *HandlerObject) DealContext(rw http.ResponseWriter, req *http.Request) {
	url := req.RequestURI

	cont := &Context{}
	cont.RW = rw
	cont.Req = req
	cont.Uri = url
	cont.Header = req.Header
	cont.RouterUri = ho.Uri
	cont.ParamNum = ho.ParamNum
	cont.ParamName = ho.ParamName
	cont.RNum = ho.RNum

	ho.Fun(cont)
}

func (c *Context) WriteOnWeb(s string) {
	_, err := c.RW.Write([]byte(s))
	if err != nil {
		fmt.Println(err.Error())
	}
}
