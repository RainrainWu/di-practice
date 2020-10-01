package logger

import "log"

// Logger extrude the public methods of logger object
type Logger interface {
	Debug()
	Info()
	Warning()
	Error()
}

type logger struct {
	directory string
	maxSize   int
}

// Option is the abstract configure option
type Option interface {
	apply(*logger)
}

type optionFunc func(*logger)

func (f optionFunc) apply(l *logger) {

	f(l)
}

// DirectoryOption construct a attribute setter for logger.directory
func DirectoryOption(dir string) Option {
	return optionFunc(func(l *logger) {
		l.directory = dir
	})
}

// MaxSizeOption construct a attribute setter for logger.maxSize
func MaxSizeOption(size int) Option {
	return optionFunc(func(l *logger) {
		l.maxSize = size
	})
}

// NewLogger instantiate a new logger
func NewLogger(opts ...Option) Logger {

	instance := &logger{
		directory: "./log/",
		maxSize:   20,
	}
	log.Println("Instantiate logger instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	return instance
}

func (l *logger) Debug() {
	log.Println("Log debug message")
}

func (l *logger) Info() {
	log.Println("Log info message")
}

func (l *logger) Warning() {
	log.Println("Log warning message")
}

func (l *logger) Error() {
	log.Println("Log error message")
}
