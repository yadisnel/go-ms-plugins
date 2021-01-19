// Package mdns provides a multicast dns registry
package mdns

import (
	"github.com/yadisnel/go-ms/v2/cmd"
	"github.com/yadisnel/go-ms/v2/registry"
	"github.com/yadisnel/go-ms/v2/registry/mdns"
)

func init() {
	cmd.DefaultRegistries["mdns"] = NewRegistry
}

// NewRegistry returns a new mdns registry
func NewRegistry(opts ...registry.Option) registry.Registry {
	return registry.NewRegistry(opts...)
}

// Domain sets the mdnsDomain
func Domain(d string) registry.Option {
	return mdns.Domain(d)
}
