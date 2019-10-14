package soap

import (
	"io"
	"sync"
)

// Conn provides read, write, close network stream functionality.
type Conn struct {
	r io.Reader
	w io.Writer

	wLock sync.Mutex
	rLock sync.Mutex
}

func NewConn(r io.Reader, w io.Writer) *Conn {
	return &Conn{
		r: r,
		w: w,
	}
}

// Writes to network stream.
func (c *Conn) Write(data []byte) (int, error) {
	c.wLock.Lock()
	defer c.wLock.Unlock()
	return c.w.Write(data)
}

// Reads from network stream.
func (c *Conn) Read(data []byte) (int, error) {
	c.rLock.Lock()
	defer c.rLock.Unlock()
	return c.r.Read(data)
}

// Closes network stream.
func (c *Conn) Close() error {
	return nil
}
