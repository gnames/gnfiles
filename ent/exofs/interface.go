package exofs

import (
	"github.com/gnames/gnfiles/ent/metadata"
	api "github.com/ipfs/go-ipfs-api"
)

type ExoFS interface {
	metadata.MetaGetter
	FileGetter
	Add(path string) (id string, err error)
	Pin(path string) error
	Unpin(id string) error
	PinExists(id string) (bool, error)
	Keys() ([]*api.Key, error)
	KeyIPNS(keyName string) (*api.Key, error)
	Publish(ipnsKeyName, id string) (ipnsKey string, err error)
}

type FileGetter interface {
	Get(source, path string) error
}
