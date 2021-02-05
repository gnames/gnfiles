package gnfiles

type Config struct {
	apiURL     string
	keyName    string
	keyID      string
	root       string
	withUpload bool
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		apiURL:     "localhost:5001",
		keyName:    "self",
		root:       "untitled",
		withUpload: false,
	}
	for i := range opts {
		opts[i](cfg)
	}
	return cfg
}

type Option func(*Config)

func OptApiURL(s string) Option {
	return func(c *Config) {
		c.apiURL = s
	}
}

func OptKeyName(s string) Option {
	return func(c *Config) {
		c.keyName = s
	}
}

func OptKeyID(s string) Option {
	return func(c *Config) {
		c.keyID = s
	}
}

func OptRoot(s string) Option {
	return func(c *Config) {
		c.root = s
	}
}

func OptWithUpload(b bool) Option {
	return func(c *Config) {
		c.withUpload = b
	}
}
