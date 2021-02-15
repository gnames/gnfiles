package gnfiles

import "github.com/gnames/gnlib/ent/gnvers"

type GNfiles interface {
	Download() error
	Upload() error
	Version() gnvers.Version
}
