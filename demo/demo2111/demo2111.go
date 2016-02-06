package main

type PacketCnn interface {
	ReadFrom(b []byte) (n int, addr int, err error)
}

type dgramOpt struct {
	PacketCnn
}

type payloadHandler struct {
	PacketCnn
}

func (c *payloadHandler) /*def*/ ReadFrom(b []byte) (a, d, e, f int) {
	println("hello")
	return 1, 2, 3, 4
}

type ipv4PacketConn struct {
	dgramOpt
	payloadHandler
}

var a PacketCnn = &PacketConn{p4: &ipv4PacketConn{}}

type PacketConn struct {
	p4 *ipv4PacketConn
}

func (ip *ipv4PacketConn) ReadFrom(b []byte) (n int, add int, err error) {
	return 1, 1, nil
}

func (c *PacketConn) ReadFrom(b []byte) (n int, addr int, err error) {
	_, _, _, _ = c.p4.ReadFrom
	return 1, 1, nil
}

func main() {
	_, _, _ = a.ReadFrom([]byte{})
}
