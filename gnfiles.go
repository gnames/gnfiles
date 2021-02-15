package gnfiles

import (
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
	return &gnfiles{
		cfg: cfg,
	}
}

func (gnf *gnfiles) Download() (err error) {
	var f files.Files
	f, err = gnf.initFiles()

	if err == nil {
		err = f.MetaDownload()
	}

	if err == nil {
		f.Download()
	}
	return err
}

func (gnf *gnfiles) Upload() (err error) {
	var f files.Files
	f, err = gnf.initFiles()

	if err == nil {
		err = f.MetaUpload()
	}

	if err == nil {
		err = f.Upload()
	}
	return err
}

func (gnf *gnfiles) initFiles() (f files.Files, err error) {
	exo := ipfs.NewExoFS(gnf.cfg.ApiURL)
	local := fs.NewLocalFS(gnf.cfg.Dir)

	fcfg := files.Config{
		Root:    gnf.cfg.Dir,
		KeyName: gnf.cfg.KeyName,
		Source:  gnf.cfg.Source,
	}
	f = files.New(fcfg, exo, local)

	err = gnsys.MakeDir(gnf.cfg.Dir)
	if err == nil {
		err = f.SetMetaData()
	}
	return f, err
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
