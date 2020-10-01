package handler

import (
	"log"

	"Repositories/di-practice/db"
)

// Handler extrude the public methods of handler object
type Handler interface {
	Get()
	Post()
	Update()
	Delete()
}

type handler struct {
	dbHandler db.Handler
}

// Option is the abstract configure option
type Option interface {
	apply(*handler)
}

type optionFunc func(*handler)

func (f optionFunc) apply(c *handler) {

	f(c)
}

// DBHandlerOption construct a attribute setter for handler.dbHandler
func DBHandlerOption(db db.Handler) Option {
	return optionFunc(func(h *handler) {
		h.dbHandler = db
	})
}

// NewHandler instantiate a new handler
func NewHandler(opts ...Option) Handler {

	instance := &handler{}
	log.Println("Instantiate handler instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.dbHandler == nil {
		instance.dbHandler = db.NewHandler()
	}
	return instance
}

func (h *handler) Get() {
	log.Println("Handle get method")
}

func (h *handler) Post() {
	log.Println("Handle post method")
}

func (h *handler) Update() {
	log.Println("Handle update method")
}

func (h *handler) Delete() {
	log.Println("Handle delete method")
}
