package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
)

type MsgHeartbeat struct {
	Header byte
	Cmd    int32
	Len    int32
	Data   []byte
}

func main() {
	fmt.Printf("Hell\n")

	var hb MsgHeartbeat
	hb.Header = 0x05
	hb.Cmd = 10010
	hb.Len = 0
	b, err := json.Marshal(&hb)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	if err := binary.Write(buf, binary.BigEndian, b); err != nil {
		panic(err)
	}

	fmt.Printf("write ok [%v]\n", buf)

}
