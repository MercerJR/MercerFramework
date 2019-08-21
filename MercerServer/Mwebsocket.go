package MercerServer

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"
)

type WebSocketFunc struct {
	OnOpen    func(*WebSocketContext)
	OnClose   func(*WebSocketContext)
	OnMessage func(*WebSocketContext)
	OnError   func(*WebSocketContext)
}

type WebSocketContext struct {
	Conn       net.Conn
	buf        *bufio.ReadWriter
	req        *http.Request
	rw         http.ResponseWriter
	MaskingKey []byte
}

func (ho *HandlerObject) DealWSContext(rw http.ResponseWriter, r *http.Request) {
	c := &WebSocketContext{}
	c.rw = rw
	c.req = r
	c.Conn = ho.conn
	onopen := ho.SocketFunc.OnOpen
	//onmessage := ho.SocketFunc.OnMessage
	//onclose := ho.SocketFunc.OnClose
	//onerror := ho.SocketFunc.OnError
	onopen(c)
	//onmessage(c)
	//onclose(c)
	//onerror(c)
}

func (c *WebSocketContext) Send(s string) error {
	data := []byte(s)
	if len(data) >= 125 {
		return errors.New("send data error")
	}
	lenth := len(data)
	maskedData := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		if c.MaskingKey != nil {
			maskedData[i] = data[i] ^ c.MaskingKey[i%4]
		} else {
			maskedData[i] = data[i]
		}
	}
	c.Conn.Write([]byte{0x81})
	var payLenByte byte
	if c.MaskingKey != nil && len(c.MaskingKey) != 4 {
		payLenByte = byte(0x80) | byte(lenth)
		c.Conn.Write([]byte{payLenByte})
		c.Conn.Write(c.MaskingKey)
	} else {
		payLenByte = byte(0x00) | byte(lenth)
		c.Conn.Write([]byte{payLenByte})
	}
	c.Conn.Write(data)
	return nil
}

func (c *WebSocketContext) Read() (data []byte, err error) {
	err = nil
	opcodeByte := make([]byte, 1)
	c.Conn.Read(opcodeByte)
	FIN := opcodeByte[0] >> 7
	RSV1 := opcodeByte[0] >> 6 & 1
	RSV2 := opcodeByte[0] >> 5 & 1
	RSV3 := opcodeByte[0] >> 4 & 1
	OPCODE := opcodeByte[0] & 15
	log.Println(RSV1, RSV2, RSV3, OPCODE)
	payloadLenByte := make([]byte, 1)
	c.Conn.Read(payloadLenByte)
	payloadLen := int(payloadLenByte[0] & 0x7F)
	mask := payloadLenByte[0] >> 7
	if payloadLen == 127 {
		extendedByte := make([]byte, 8)
		c.Conn.Read(extendedByte)
	}
	maskingByte := make([]byte, 4)
	if mask == 1 {
		//c.Conn.Read(maskingByte)
		c.MaskingKey = maskingByte
	}
	payloadDataByte := make([]byte, payloadLen)
	//c.Conn.Read(payloadDataByte)
	log.Println("data:", payloadDataByte)
	dataByte := make([]byte, payloadLen)
	for i := 0; i < payloadLen; i++ {
		if mask == 1 {
			dataByte[i] = payloadDataByte[i] ^ maskingByte[i%4]
		} else {
			dataByte[i] = payloadDataByte[i]
		}
	}
	if FIN == 1 {
		data = dataByte
		return
	}
	nextData, err := c.Read()
	if err != nil {
		return
	}
	data = append(data, nextData...)
	return
}
