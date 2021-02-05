package gnfiles

import "github.com/gnames/gnlib/ent/gnvers"

type GNfiles interface {
	Sync() error
	Version() gnvers.Version
}
