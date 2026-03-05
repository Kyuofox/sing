package socks

import (
	"net"

	"github.com/kyuofox/sing/common"
	"github.com/kyuofox/sing/common/buf"
	"github.com/kyuofox/sing/common/bufio"
	M "github.com/kyuofox/sing/common/metadata"
	N "github.com/kyuofox/sing/common/network"
)

var _ N.VectorisedPacketWriter = (*VectorisedAssociatePacketConn)(nil)

type VectorisedAssociatePacketConn struct {
	AssociatePacketConn
	N.VectorisedPacketWriter
}

func NewVectorisedAssociateConn(conn net.Conn, writer N.VectorisedWriter, remoteAddr M.Socksaddr, underlying net.Conn) *VectorisedAssociatePacketConn {
	return &VectorisedAssociatePacketConn{
		AssociatePacketConn{
			AbstractConn: conn,
			conn:         bufio.NewExtendedConn(conn),
			remoteAddr:   remoteAddr,
			underlying:   underlying,
		},
		&bufio.UnbindVectorisedPacketWriter{VectorisedWriter: writer},
	}
}

func (c *VectorisedAssociatePacketConn) WriteVectorisedPacket(buffers []*buf.Buffer, destination M.Socksaddr) error {
	header := buf.NewSize(3 + M.SocksaddrSerializer.AddrPortLen(destination))
	defer header.Release()
	common.Must(header.WriteZeroN(3))
	err := M.SocksaddrSerializer.WriteAddrPort(header, destination)
	if err != nil {
		return err
	}
	return c.VectorisedPacketWriter.WriteVectorisedPacket(append([]*buf.Buffer{header}, buffers...), destination)
}

func (c *VectorisedAssociatePacketConn) FrontHeadroom() int {
	return 0
}
