// Package quic provides a QUIC based transport
package quic

import (
	"github.com/yadisnel/go-ms/v2/config/cmd"
	"github.com/yadisnel/go-ms/v2/transport"
	"github.com/yadisnel/go-ms/v2/transport/quic"
)

func init() {
	cmd.DefaultTransports["quic"] = NewTransport
}

func NewTransport(opts ...transport.Option) transport.Transport {
	return quic.NewTransport(opts...)
}
