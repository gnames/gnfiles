package gnfiles

import (
	"log"

	"github.com/gnames/gnfiles/ent/files"
	"github.com/gnames/gnfiles/io/fs"
	"github.com/gnames/gnfiles/io/ipfs"
	"github.com/gnames/gnlib/ent/gnvers"
	"github.com/gnames/gnsys"
	api "github.com/ipfs/go-ipfs-api"
)

type gnfiles struct {
	cfg *Config
}

func New(cfg *Config) GNfiles {
	if cfg.keyName == "" {
		log.Fatal("Cannot download data without a key")
	}
	log.Printf("Files will be downloaded to '%s' directory.", cfg.root)
	return &gnfiles{
		cfg: cfg,
	}
}

func (gnf *gnfiles) Sync() (err error) {
	exo := ipfs.NewExoFS(gnf.cfg.apiURL)
	local := fs.NewLocalFS(gnf.cfg.root)
	f := files.New(gnf.cfg.root, gnf.key(), exo, local)

	err = gnsys.MakeDir(gnf.cfg.root)
	if err == nil {
		err = f.SetMetaData()
	}
	if err == nil {
		err = f.Dump()
	}
	if err == nil && gnf.cfg.withUpload {
		err = f.Update()
	}
	return err
}

func (gnf *gnfiles) Version() gnvers.Version {
	return gnvers.Version{}
}

func (gnf *gnfiles) key() api.Key {
	return api.Key{
		Name: gnf.cfg.keyName,
		Id:   gnf.cfg.keyID,
	}
}
