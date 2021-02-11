package ipfs

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gnames/gnfiles/ent/exofs"
	"github.com/gnames/gnfiles/ent/metadata"
	"github.com/gnames/gnfiles/ent/paths"
	"github.com/gnames/gnfmt"
	api "github.com/ipfs/go-ipfs-api"
)

type efs struct {
	ipfs *api.Shell
}

func NewExoFS(url string) exofs.ExoFS {
	return &efs{
		ipfs: api.NewShell(url),
	}
}

func (e *efs) GetMetaData(
	source, metaPath string,
) (metadata.MetaFiles, error) {
	var meta metadata.MetaFiles
	err := e.Get(source, metaPath)
	if err != nil {
		return meta, err
	}

	txtJSON, err := ioutil.ReadFile(metaPath)
	if err != nil {
		return meta, err
	}

	enc := gnfmt.GNjson{}
	// skipping the error for cases where txtJSON does not contain
	// metadata information
	err = enc.Decode(txtJSON, &meta)
	if err != nil {
		log.Printf("Cannod decode _META.json: %v", err)
	}
	res := make(map[string]*metadata.MetaData)
	for k, v := range meta {
		if v.ID == "" || strings.HasPrefix(k, "/_META") {
			continue
		}
		res[k] = v
	}
	mf := metadata.MetaFiles(res)
	return mf, err
}

func (e *efs) Add(path string) (id string, err error) {
	var f io.Reader
	f, err = os.Open(path)
	if err != nil {
		return id, err
	}
	return e.ipfs.Add(f)
}

func (e *efs) Get(source, path string) error {
	if isFile(source) {
		return GetterFile{}.Get(source, path)
	}

	if isURL(source) {
		return GetterHTTP{}.Get(source, path)
	}

	if isPathIPNS(source) {
		return e.ipfs.Get(source, path)
	}
	if isKeyIPNS(source) {
		return e.ipfs.Get(paths.IPNSPath(source), path)
	}
	if isPathIPFS(source) {
		source = strings.TrimLeft(source, "/ipfs/")
	}
	if isCID(source) {
		return e.ipfs.Get(source, path)
	}
	return fmt.Errorf("unknown source type '%s'", source)
}

func (e *efs) Pin(id string) error {
	return e.ipfs.Pin(paths.IPFSPath(id))
}

func (e *efs) Unpin(id string) error {
	return e.ipfs.Unpin(paths.IPFSPath(id))
}

func (e *efs) PinExists(id string) (bool, error) {
	res, err := e.ipfs.Pins()
	if err != nil {
		return false, err
	}
	for k, v := range res {
		if k == id && v.Type == "recursive" {
			return true, nil
		}
	}
	return false, nil
}

func (e *efs) Keys() ([]*api.Key, error) {
	return e.ipfs.KeyList(context.Background())
}

func (e *efs) KeyIPNS(keyName string) (*api.Key, error) {
	var res *api.Key
	keys, err := e.Keys()

	if err != nil {
		return res, err
	}
	for i := range keys {
		if keys[i].Name == keyName {
			return keys[i], nil
		}
	}
	return res, fmt.Errorf("cannot find IPNS key '%s'", keyName)
}

func (e *efs) Publish(name, id string) (key string, err error) {
	resp, err := e.ipfs.PublishWithDetails(id, name, 0, 10*time.Second, false)
	if err != nil {
		return "", err
	}
	return resp.Name, err
}
