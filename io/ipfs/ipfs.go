package ipfs

import (
	"io"
	"os"
	"time"

	"github.com/gnames/gnfiles/ent/exofs"
	api "github.com/ipfs/go-ipfs-api"
)

type efs struct {
	ipfs *api.Shell
}

func NewExoFS(URL string) exofs.ExoFS {
	return &efs{
		ipfs: api.NewShell(URL),
	}
}

func (fs *efs) Connect(URL string) error {
	return nil
}

func (e *efs) Add(path string) (id string, err error) {
	var f io.Reader
	f, err = os.Open(path)
	if err != nil {
		return id, err
	}
	return e.ipfs.Add(f)
}

func (e *efs) Get(id, path string) error {
	return e.ipfs.Get(id, path)
}

func (e *efs) Pin(id string) error {
	return e.ipfs.Pin(idPath(id))
}

func (e *efs) Unpin(id string) error {
	return e.ipfs.Unpin(idPath(id))
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

func (fs *efs) Publish(key, id string) (name string, err error) {
	resp, err := fs.ipfs.PublishWithDetails(id, key, 0, 10*time.Second, false)
	if err != nil {
		return "", err
	}
	return resp.Name, err
}

func idPath(id string) string {
	return "/ipfs/" + id
}
