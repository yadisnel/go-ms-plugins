// Package etcd provides an etcd v3 service registry
package etcd

import (
	"github.com/yadisnel/go-ms/v2/config/cmd"
	"github.com/yadisnel/go-ms/v2/registry"
	"github.com/yadisnel/go-ms/v2/registry/etcd"
)

func init() {
	cmd.DefaultRegistries["etcd"] = etcd.NewRegistry
}

func NewRegistry(opts ...registry.Option) registry.Registry {
	return etcd.NewRegistry(opts...)
}
