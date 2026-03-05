package network

import (
	"github.com/kyuofox/sing/common/buf"
	M "github.com/kyuofox/sing/common/metadata"
)

type VectorisedWriter interface {
	WriteVectorised(buffers []*buf.Buffer) error
}

type VectorisedPacketWriter interface {
	WriteVectorisedPacket(buffers []*buf.Buffer, destination M.Socksaddr) error
}
