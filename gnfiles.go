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
	if cfg.KeyName == "" {
		log.Fatal("Cannot download data without a key")
	}
	log.Printf("Files will be downloaded to '%s' directory.", cfg.Dir)
	return &gnfiles{
		cfg: cfg,
	}
}

func (gnf *gnfiles) Sync() (err error) {
	exo := ipfs.NewExoFS(gnf.cfg.ApiURL)
	local := fs.NewLocalFS(gnf.cfg.Dir)

	fcfg := files.Config{
		Root:    gnf.cfg.Dir,
		KeyName: gnf.cfg.KeyName,
		Source:  gnf.cfg.Source,
	}
	f := files.New(fcfg, exo, local)

	err = gnsys.MakeDir(gnf.cfg.Dir)
	if err == nil {
		err = f.SetMetaData()
	}
	if err == nil {
		downloadAll := !gnf.cfg.WithUpload
		err = f.Dump(downloadAll)
	}
	if err == nil && gnf.cfg.WithUpload {
		err = f.Update()
	}
	return err
}

func (gnf *gnfiles) Version() gnvers.Version {
	return gnvers.Version{}
}

func (gnf *gnfiles) key() api.Key {
	return api.Key{
		Name: gnf.cfg.KeyName,
		Id:   gnf.cfg.Source,
	}
}
