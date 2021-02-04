package exofs

type ExoFS interface {
	Connect(string) error
	Add(path string) (id string, err error)
	Get(id, path string) error
	Pin(path string) error
	Unpin(id string) error
	PinExists(id string) (bool, error)
	Publish(key, id string) (keyID string, err error)
}
