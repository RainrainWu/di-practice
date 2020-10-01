package main

import (
	"Repositories/di-practice/config"
	"Repositories/di-practice/db"
	"Repositories/di-practice/handler"
	"Repositories/di-practice/job"
	"Repositories/di-practice/logger"
	"Repositories/di-practice/server"
	"context"

	"go.uber.org/fx"
)

// ProvideConfig provide a instance of config.Config
func ProvideConfig() config.Config {
	return config.NewConfig()
}

// ProvideLogger provide a instance of logger.Logger
func ProvideLogger() logger.Logger {
	return logger.NewLogger()
}

// ProvideJobQueue provide a instance of queue.Jobqueue
func ProvideJobQueue() job.JobQueue {
	return job.NewJobQueue()
}

// ProvideDBHandler provide a instance of db.Handler
func ProvideDBHandler(conf config.Config) db.Handler {
	return db.NewHandler(
		db.TimeoutOption(5),
		db.RetryOption(3),
		db.ConfigOption(conf),
	)
}

// ProvideHandler provide a instance of handler.Handler
func ProvideHandler(db db.Handler) handler.Handler {
	return handler.NewHandler(
		handler.DBHandlerOption(db),
	)
}

// ProvideServer provide a instance of handler.Handler
func ProvideServer(
	l logger.Logger, d db.Handler, c config.Config, h handler.Handler, q job.JobQueue,
) server.Server {

	return server.NewServer(
		server.LoggerOption(l),
		server.DBHandlerOption(d),
		server.ConfigOption(c),
		server.HandlerOption(h),
		server.QueueOption(q),
	)
}

func register(lifecycle fx.Lifecycle, s server.Server, d db.Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				d.CreateTable()
				s.Start()
				return nil
			},
			OnStop: func(context.Context) error {
				s.Stop()
				return nil
			},
		},
	)
}

func main() {
	fx.New(
		fx.Provide(
			ProvideConfig,
			ProvideLogger,
			ProvideJobQueue,
			ProvideDBHandler,
			ProvideHandler,
			ProvideServer,
		),
		fx.Invoke(register),
	).Run()
}
