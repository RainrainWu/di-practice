package server

import (
	"log"

	"Repositories/di-practice/config"
	"Repositories/di-practice/db"
	"Repositories/di-practice/handler"
	"Repositories/di-practice/job"
	"Repositories/di-practice/logger"
)

// Server extrude the public methods of server object
type Server interface {
	Start()
	Stop()
}

type server struct {
	log       logger.Logger
	dbHandler db.Handler
	conf      config.Config
	handler   handler.Handler
	queue     job.JobQueue
}

// Option is the abstract configure option
type Option interface {
	apply(*server)
}

type optionFunc func(*server)

func (f optionFunc) apply(s *server) {

	f(s)
}

// LoggerOption construct a attribute setter for logger.Logger
func LoggerOption(l logger.Logger) Option {
	return optionFunc(func(s *server) {
		s.log = l
	})
}

// DBHandlerOption construct a attribute setter for handler.dbHandler
func DBHandlerOption(db db.Handler) Option {
	return optionFunc(func(s *server) {
		s.dbHandler = db
	})
}

// ConfigOption construct a attribute setter for handler.conf
func ConfigOption(conf config.Config) Option {
	return optionFunc(func(s *server) {
		s.conf = conf
	})
}

// HandlerOption construct a attribute setter for handler.handler
func HandlerOption(h handler.Handler) Option {
	return optionFunc(func(s *server) {
		s.handler = h
	})
}

// QueueOption construct a attribute setter for handler.queue
func QueueOption(q job.JobQueue) Option {
	return optionFunc(func(s *server) {
		s.queue = q
	})
}

// NewServer instantiate a new server
func NewServer(opts ...Option) Server {

	instance := &server{}
	log.Println("Instantiate server instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.dbHandler == nil {
		instance.dbHandler = db.NewHandler()
	}
	if instance.conf == nil {
		instance.conf = config.NewConfig()
	}
	if instance.handler == nil {
		instance.handler = handler.NewHandler()
	}
	if instance.queue == nil {
		instance.queue = job.NewJobQueue()
	}
	return instance
}

func (s *server) Start() {
	log.Println("Start running server")
}

func (s *server) Stop() {
	log.Println("Stop running server")
}
