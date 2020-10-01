package db

import (
	"log"

	"Repositories/di-practice/config"
)

// Handler extrude the public methods of handler object
type Handler interface {
	CreateTable()
	DropTable()
}

type handler struct {
	timeout int
	retry   int
	conf    config.Config
}

// Option is the abstract configure option
type Option interface {
	apply(*handler)
}

type optionFunc func(*handler)

func (f optionFunc) apply(c *handler) {

	f(c)
}

// TimeoutOption construct a attribute setter for handler.timeout
func TimeoutOption(time int) Option {
	return optionFunc(func(h *handler) {
		h.timeout = time
	})
}

// RetryOption construct a attribute setter for handler.retry
func RetryOption(count int) Option {
	return optionFunc(func(h *handler) {
		h.retry = count
	})
}

// ConfigOption construct a attribute setter for handler.conf
func ConfigOption(conf config.Config) Option {
	return optionFunc(func(h *handler) {
		h.conf = conf
	})
}

// NewHandler instantiate a new handler
func NewHandler(opts ...Option) Handler {

	instance := &handler{
		timeout: 3,
		retry:   2,
	}
	log.Println("Instantiate db instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.conf == nil {
		instance.conf = config.NewConfig()
	}
	return instance
}

func (h *handler) CreateTable() {
	log.Println("Create database table")
}

func (h *handler) DropTable() {
	log.Println("Drop database table")
}
