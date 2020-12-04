// Package http returns a http2 transport using net/http
package http

import (
	"github.com/yadisnel/go-ms/v2/config/cmd"
	"github.com/yadisnel/go-ms/v2/transport"
)

func init() {
	cmd.DefaultTransports["http"] = NewTransport
}

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewTransport(opts...)
}
