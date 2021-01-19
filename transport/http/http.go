// Package http returns a http2 transport using net/http
package http

import (
	"github.com/yadisnel/go-ms/v2/transport"
)

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewHTTPTransport(opts...)
}
