package etcd

import (
	"github.com/yadisnel/go-ms/v2/registry"
	"github.com/yadisnel/go-ms/v2/registry/etcd"
)

// Auth allows you to specify username/password
func Auth(username, password string) registry.Option {
	return etcd.Auth(username, password)
}
