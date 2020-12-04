// Package memory provides an in-memory registry
package memory

import (
	"github.com/yadisnel/go-ms/v2/config/cmd"
	"github.com/yadisnel/go-ms/v2/registry"
	"github.com/yadisnel/go-ms/v2/registry/memory"
)

func init() {
	cmd.DefaultRegistries["memory"] = NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return memory.NewRegistry(opts...)
}
