package graf

import (
	"fmt"
	"net"
	"time"
)

type Graf struct {
	conn net.Conn
}

func Dial(address string, timeout time.Duration) (g Graf, err error) {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return
	}
	g = Graf{conn: conn}
	return
}

func (g Graf) Send(path string, value int64, time time.Time) (err error) {
	_, err = fmt.Fprintf(g.conn, "%s %d %d\n", path, value, time.Unix())
	return
}
