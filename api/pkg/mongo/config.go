package mongo

import (
	"fmt"
)

// Config is the configuration for the MongoDB client.
type Config struct {
	User            string   `yaml:"user"`
	Password        string   `yaml:"password"`
	Host            string   `yaml:"host"`
	DB              string   `yaml:"db"`
	AuthType        string   `yaml:"auth_type"`
	Port            int      `yaml:"port"`
	RetryTimeout    int      `yaml:"retry_timeout"`
	CollectionNames []string `yaml:"collections"`
}

// URL returns the connection URL.
func (c Config) URL() string {
	url := fmt.Sprintf("mongodb://%s:%d/%s", c.Host, c.Port, c.DB)
	if c.AuthType != "" && c.AuthType != "no" {
		url = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", c.User, c.Password, c.Host, c.Port, c.DB)
	}
	return url
}
