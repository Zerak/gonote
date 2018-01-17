package core

type EncryptFunc func(dst, src []byte)
type DecryptFunc func(dst, src []byte)

type Packet interface {
	Len() int
	Bytes() []byte
}

type BytesPacket []byte

func (p BytesPacket) Len() int      { return len(p) }
func (p BytesPacket) Bytes() []byte { return []byte(p) }
