package gnfiles

import "github.com/gnames/gnlib/ent/gnvers"

type GNfiles interface {
	Dump() error
	Update() error
	Version() gnvers.Version
}
