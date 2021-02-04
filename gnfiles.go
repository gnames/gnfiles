package gnfiles

import "github.com/gnames/gnlib/ent/gnvers"

type gnfiles struct {
	cfg Config
}

func New(cfg Config) GNfiles {
	return &gnfiles{cfg: cfg}
}

func (gnf *gnfiles) Dump() error {
	return nil
}

func (gnf *gnfiles) Update() error {
	return nil
}

func (gnf *gnfiles) Version() gnvers.Version {
	return gnvers.Version{}
}
