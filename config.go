package gnfiles

import (
	"log"
	"path/filepath"

	"github.com/gnames/gnsys"
)

type Config struct {
	ApiURL     string
	KeyName    string
	Source     string
	Dir        string
	WithUpload bool
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{
		ApiURL:     "localhost:5001",
		KeyName:    "self",
		Dir:        "untitled",
		WithUpload: false,
	}
	for i := range opts {
		opts[i](cfg)
	}
	return cfg
}

type Option func(*Config)

func OptApiURL(s string) Option {
	return func(c *Config) {
		c.ApiURL = s
	}
}

func OptKeyName(s string) Option {
	return func(c *Config) {
		c.KeyName = s
	}
}

func OptSource(s string) Option {
	return func(c *Config) {
		c.Source = s
	}
}

func OptDir(s string) Option {
	return func(c *Config) {
		c.Dir = prepareDir(s)
	}
}

func prepareDir(s string) string {
	var err error
	s, err = gnsys.ConvertTilda(s)
	if err == nil {
		s, err = filepath.Abs(s)
	}

	if err != nil {
		log.Fatal(err)
	}
	return s
}
