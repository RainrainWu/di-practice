package job

import (
	"log"
	"sync"
)

// JobQueue extrude the public methods of logger object
type JobQueue interface {
	Append()
	Pop()
}

type jobQueue struct {
	queue chan string
	size  int
	mutex sync.Mutex
}

// Option is the abstract configure option
type Option interface {
	apply(*jobQueue)
}

type optionFunc func(*jobQueue)

func (f optionFunc) apply(q *jobQueue) {

	f(q)
}

// SizeOption construct a attribute setter for jobQueue.size
func SizeOption(size int) Option {
	return optionFunc(func(q *jobQueue) {
		q.size = size
	})
}

// NewJobQueue instantiate a new config
func NewJobQueue(opts ...Option) JobQueue {

	instance := &jobQueue{
		size: 5,
	}
	log.Println("Instantiate jobqueue instance")
	for _, opt := range opts {
		opt.apply(instance)
	}
	if instance.queue == nil {
		instance.queue = make(chan string, instance.size)
	}
	return instance
}

func (q *jobQueue) Append() {
	log.Println("Append new job")
}

func (q *jobQueue) Pop() {
	log.Println("Pop next job")
}
