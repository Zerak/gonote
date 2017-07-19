package core

import (
	"encoding/binary"
	"errors"
	"net"
	"sync/atomic"
	"time"
)

type PacketHandler func(b []byte)
type EncryptFunc func(dst, src []byte)
type DecryptFunc func(dst, src []byte)

type Packet interface {
	Len() int
	Bytes() []byte
}

type BytesPacket []byte

func (p BytesPacket) Len() int      { return len(p) }
func (p BytesPacket) Bytes() []byte { return []byte(p) }

type RWSession interface {
	// Conn return the real connection
	// of the current client
	Conn() net.Conn

	// Send send the packet
	Send(p Packet)

	// SetCryptFunc set crypt function
	// if not set crypt will use default crypt
	SetEncryptFunc(encrypt EncryptFunc)

	// SetDecryptFunc set decrypt function
	// if not set will use default decrypt
	SetDecryptFunc(decrypt DecryptFunc)

	// Run start read write session
	//
	// the nil val of start/end the service will
	// not notify when service start or end
	Run(start, end func())
}

var (
	errPacketTooBig = errors.New("packet msg too big")
)
var (
	HeaderSize  = 2
	MaxBodySize = 4 * 1024 * 1024
)

type rwSession struct {
	// addr is the connection address of the client
	addr string

	// conn is the real session about client
	conn net.Conn

	// packHandle is the call func when
	// receive packet from client
	packHandle PacketHandler

	// closed is a flag of the conn closed or not
	closed int32

	// timeout is the connection read timeout
	timeout time.Duration

	// encrypt/decrypt is the encryption/decryption function
	encrypt EncryptFunc
	decrypt DecryptFunc

	// quitChan is the channel of quit todo
	quitChan chan struct{}

	// writeChan is the channel of write
	writeChan chan Packet

	defaultHeaderLen int
	header           [HeaderSize]byte
	body             []byte
}

func NewRWSession(ph PacketHandler) *RWSession {
	return &rwSession{packHandle: ph}
}

func (w *rwSession) setClosed() {
	atomic.StoreInt32(&w.closed, 1)
}

func (w *rwSession) getClosed() bool {
	return atomic.LoadInt32(&w.closed) == 1
}

func (rw *rwSession) read(b []byte) (int, error) {
	length := len(b)
	readNum := 0

	for readNum < length {
		n, err := rw.conn.Read(b[readNum:length])
		readNum += n
		if err != nil {
			return readNum, err
		}
	}
	return readNum, nil
}

func (rw *rwSession) readPacket() (int, error) {
	total := 0
	if rw.timeout > 0 {
		rw.conn.SetReadDeadline(time.Now().Add(rw.timeout))
	}

	n, err := rw.read(rw.header[:])
	total += n
	if err != nil {
		return total, err
	}

	// parse header
	length := binary.BigEndian.Uint32(rw.header[:])
	if length > MaxBodySize {
		return total, errPacketTooBig
	}

	// parse body
	if len(rw.body) < length {
		rw.body = make([]byte, length)
	}
	n, err = rw.read(rw.body[:length])
	total += n
	if err != nil {
		return total, err
	}

	//if rw.encrypt != nil {
	//	rw.decrypt()
	//}
	rw.packHandle(rw.body[:length])
	return total, nil
}

func (rw *rwSession) startReadLoop(start, end chan<- struct{}) {
	start <- struct{}{}
	for {
		_, err := rw.readPacket()
		if err != nil {
			rw.setClosed()
		}
		if rw.getClosed() {
			break
		}
	}
	end <- struct{}{}
}

func (rw *rwSession) startWriteLoop(start, end chan<- struct{}) {
	start <- struct{}{}
	remain := 0
	for {
		if rw.getClosed() {
			remain = len(rw.writeChan)
			break
		}

		select {
		case p := <-rw.writeChan:
			_, err := rw.conn.Write(p.Bytes())
			if err != nil {
				rw.setClosed()
			}
		case <-time.After(time.Second):
		}
	}

	for i := 0; i < remain; i++ {
		p := <-rw.writeChan
		_, err := rw.conn.Write(p.Bytes())
		if err != nil {
			break
		}
	}

	rw.conn.Close()
	end <- struct{}{}
}

func (rw *rwSession) Conn() net.Conn {
	return rw.conn
}

// Send just put the p to the writeChan
func (w *rwSession) Send(p Packet) {
	if p.Len() > 0 && !w.getClosed() {
		w.writeChan <- p
	}
}

func (rw *rwSession) SetEncryptFunc(encrypt EncryptFunc) {
	rw.encrypt = encrypt
}

func (rw *rwSession) SetDecryptFunc(decrypt DecryptFunc) {
	rw.decrypt = decrypt
}

func (rw *rwSession) Run(onStartSession, onEndSession func()) {
	startRead := make(chan struct{})
	startWrite := make(chan struct{})
	endRead := make(chan struct{})
	endWrite := make(chan struct{})

	go rw.startReadLoop(startRead, endRead)
	go rw.startWriteLoop(startWrite, endWrite)

	<-startRead
	<-startWrite

	if onStartSession != nil {
		onStartSession()
	}

	<-endRead
	<-endWrite

	if rw.conn != nil {
		rw.conn.Close()
	}

	if onEndSession != nil {
		onEndSession()
	}
}
