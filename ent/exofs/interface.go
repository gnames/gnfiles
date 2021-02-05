package exofs

import (
	"github.com/gnames/gnfiles/ent/metadata"
	api "github.com/ipfs/go-ipfs-api"
)

type ExoFS interface {
	Connect(string) error
	Add(path string) (id string, err error)
	Get(id, path string) error
	Pin(path string) error
	Unpin(id string) error
	PinExists(id string) (bool, error)
	KeyIPNS(keyName string) (api.Key, error)
	Publish(ipnsKeyName, id string) (ipnsKey string, err error)
	MetaData(ipnsPath, path string) (metadata.MetaFiles, error)
}
