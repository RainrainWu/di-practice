package config

import "log"

// Config extrude the public methods of config object
type Config interface {
	GetDatabaseConfig()
	GetServerConfig()
}

type config struct {
	databaseConfig map[string]string
	serverConfig   map[string]string
}

// Option is the abstract configure option
type Option interface {
	apply(*config)
}

type optionFunc func(*config)

func (f optionFunc) apply(c *config) {

	f(c)
}

// DatabaseConfigOption construct a attribute setter for config.databaseConfig
func DatabaseConfigOption(conf map[string]string) Option {
	return optionFunc(func(c *config) {
		c.databaseConfig = conf
	})
}

// ServerOption construct a attribute setter for config.serverConfig
func ServerOption(conf map[string]string) Option {
	return optionFunc(func(c *config) {
		c.serverConfig = conf
	})
}

// NewConfig instantiate a new config
func NewConfig(opts ...Option) Config {

	instance := &config{
		databaseConfig: map[string]string{"host": "127.0.0.1"},
		serverConfig:   map[string]string{"port": "5050"},
	}
	log.Println("Instantiate config instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	return instance
}

func (c *config) GetDatabaseConfig() {
	log.Println("Return database config")
}

func (c *config) GetServerConfig() {
	log.Println("Return server config")
}
