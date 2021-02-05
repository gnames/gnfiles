package files

type Files interface {
	SetMetaData() error
	PublishMetaData() error
	Dump() error
	Update() error
}
