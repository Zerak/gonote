package main

import (
	//"bufio"
	//"bytes"
	"errors"
	"fmt"
	"io"

	"bytes"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

type ClientMessage struct {
	Body             *string `protobuf:"bytes,1,req,name=body" json:"body,omitempty"`
	Ext              *string `protobuf:"bytes,2,opt,name=ext" json:"ext,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ClientMessage) Reset()                    { *m = ClientMessage{} }
func (m *ClientMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientMessage) ProtoMessage()               {}
func (*ClientMessage) Descriptor() ([]byte, []int) { return []byte(""), []int{3} }

func (m *ClientMessage) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *ClientMessage) GetExt() string {
	if m != nil && m.Ext != nil {
		return *m.Ext
	}
	return ""
}

const DefaultByteNumForLength = 2

type FactoryFunc func() proto.Message

var (
	protocolFactory map[string]FactoryFunc

	ErrTooShort           = errors.New("too short")
	ErrUnknownMessageName = errors.New("unknown message name")
)

func DecodeLength(buf []byte) int {
	n := 0
	for i, b := range buf {
		n |= int(b) << (uint32(i) << 3)
	}
	return n
}

func EncodeLength(length int, buf []byte) []byte {
	for i := range buf {
		b := length & 0xFF
		fmt.Printf("i:%v b:%v byteb[%v]\n", i, b, byte(b))
		buf[i] = byte(b)
		length = length >> (uint32(i+1) << 3)
	}
	return buf
}

func encodeMessageHeader(w io.Writer, v string, bodySize int) error {
	name := v
	totalSize := DefaultByteNumForLength + (DefaultByteNumForLength + len(name)) + bodySize
	// 写包大小
	if _, err := w.Write(EncodeLength(totalSize, make([]byte, DefaultByteNumForLength))); err != nil {
		return err
	}
	// 写包名大小
	if _, err := w.Write(EncodeLength(len(name), make([]byte, DefaultByteNumForLength))); err != nil {
		return err
	}
	// 写包名
	_, err := io.WriteString(w, name)
	return err
}

func decodeMessageHeader(b []byte) (n int, name string, err error) {
	byteNum := DefaultByteNumForLength
	// 解包大小
	if len(b[n:]) < byteNum {
		err = ErrTooShort
		return
	}
	DecodeLength(b[n : n+byteNum])
	n += byteNum
	// 解包名大小
	if len(b[n:]) < byteNum {
		err = ErrTooShort
		return
	}
	nameLength := DecodeLength(b[n : n+byteNum])
	n += byteNum
	// 解包名
	if len(b[n:]) < nameLength {
		err = ErrTooShort
		return
	}
	name = string(b[n : n+nameLength])
	n += nameLength
	return
}

//func Encode(w io.Writer, v proto.Message) error {
//	data, err := proto.Marshal(v)
//	if err == nil {
//		encodeMessageHeader(w, v, len(data))
//		_, err = w.Write(data)
//	}
//	return err
//}

//func Decode(b []byte) (proto.Message, error) {
//	n, name, err := decodeMessageHeader(b)
//	if err != nil {
//		return nil, err
//	}
//	fn, ok := protocolFactory[name]
//	if !ok {
//		return nil, ErrUnknownMessageName
//	}
//	v := fn()
//	err = proto.Unmarshal(b[n:], v)
//	return v, err
//}

//func main() {
//	//l := 10
//	//reBuf := EncodeLength(l, []byte("ab"))
//	//fmt.Printf("rebuf:%v, lenRebuf[%v]\n", reBuf, len(reBuf))
//	//
//	//reLen := DecodeLength(reBuf)
//	//fmt.Printf("relen:%v\n", reLen)
//
//	buf := bufio.NewWriter(bytes.NewBuffer([]byte("a")))
//	encodeMessageHeader(buf, "name", 10)
//	fmt.Printf("buf:%v\n", buf)
//}

func main() {
	var data = []byte(`{"status": 200}`)

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Println("error:", err)
		return
	}

	var status = int(result["status"].(float64)) //ok
	fmt.Println("status value:", status)

	fmt.Println("ver2")
	ver2()
}

func ver2() {
	var data = []byte(`{"status": 200}`)

	var result struct {
		Status int `json:"status"`
	}

	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&result); err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Printf("result: %+v", result)
}
