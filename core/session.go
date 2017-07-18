package core

import (
	"net"
	"sync/atomic"
)

type CryptFunc func(dst, src []byte)
type DecryptFunc func(dst, src []byte)

type Packet interface {
	Len() int
	Bytes() []byte
}

type ReadSession struct {
	headerSize int
	crypt      CryptFunc
}

func NewReadSession() *ReadSession {
	return &ReadSession{}
}

func (r *ReadSession) ReadPacket() (int, error) {
	if  {
		
	}
	if r.crypt != nil {
		r.crypt()
	}
	return 0, nil
}

func (r *ReadSession) SetHeaderSize(size int) {
	r.headerSize = size
}
func (r *ReadSession) SetCryptFunc(crypt CryptFunc) {
	r.crypt = crypt
}

type WriteSession struct {
	addr    string
	conn    net.Conn
	closed  int32
	decrypt DecryptFunc

	quitChan  chan struct{}
	writeChan chan Packet
}

func NewWriteSession() *WriteSession {
	return &WriteSession{}
}

func (w *WriteSession) setClosed() {
	atomic.StoreInt32(&w.closed, 1)
}
func (w *WriteSession) getClosed() bool {
	return atomic.LoadInt32(&w.closed) == 1
}
func (w *WriteSession) Send(p Packet) {
	if p.Len() > 0 && !w.getClosed() {
		w.writeChan <- p
	}
}

type RWSession struct {
	*ReadSession
	*WriteSession
}

func NewRWSession(r ReadSession, w WriteSession) *RWSession {
	return &RWSession{ReadSession: r, WriteSession: w}
}

func (rw *RWSession) startReadLoop(start, end chan<- struct{}) {
	start <- struct{}{}

	for {
		_, err := rw.ReadSession.ReadPacket()
		if err != nil {
			rw.setClosed()
		}
		if rw.getClosed() {
			break
		}
	}
	end <- struct{}{}
}

func (rw *RWSession) startWriteLoop(start, end chan<- struct{}) {

}

func (rw *RWSession) Run(onStartSession, onEndSession func()) {
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
