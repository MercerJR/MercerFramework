package MercerServer

import (
	"MercerFrame/MercerLog"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func (ho *HandlerObject) IntoWebsocket(w http.ResponseWriter, r *http.Request, config WebSocketFunc) {
	key := r.Header.Get("Sec-WebSocket-Key")

	//用sha1进行加密算法调用
	s := sha1.New()
	s.Write(QuickStringToBytes(key +
		"258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))
	b := s.Sum(nil)

	//返回base64的编码
	accept := base64.StdEncoding.EncodeToString(b)

	upgrade := "HTTP/1.1 101 Switching Protocols\r\n" +
		"Upgrade: websocket\r\n" +
		"Connection: Upgrade\r\n" +
		"Sec-WebSocket-Accept: " + accept + "\r\n\r\n"
	log.Println("response:", upgrade)

	//调用劫持，升级HTTP协议
	hijack := w.(http.Hijacker)
	con, buf, err := hijack.Hijack()
	if err != nil {
		MercerLog.Error.Println(err)
	}

	if con == nil {
		fmt.Println("nil")
	}

	//c := &WebSocketContext{
	//	Conn: con,
	//	buf:  buf,
	//}
	//con.Write(QuickStringToBytes(upgrade))
	//buf.Flush()
	if lenth, err := buf.Write([]byte(upgrade)); err != nil {
		log.Println(err)
	} else {
		log.Println("send len:", lenth)
	}
	buf.Flush()
	ho.conn = con

	fmt.Println("连接成功")
	content := make([]byte, 1024)
	_,errr := con.Read(content)
	log.Println(string(content))
	if err != nil {
		log.Println(errr)
	}
	//OnOpen方法
	//config.OnOpen(c)

	//对WebSocket的解码
	//ParseWebSocket(c, config)

}
