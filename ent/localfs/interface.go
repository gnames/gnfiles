package localfs

type LocalFS interface {
	MetaData() []MetaData
	Dump() error
	Update() error
}
