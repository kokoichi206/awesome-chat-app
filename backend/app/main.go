package main

import (
	"context"
	"net"
	"os"

	"github.com/opentracing/opentracing-go"

	"github.com/kokoichi206/awesome-chat-app/backend/config"
	"github.com/kokoichi206/awesome-chat-app/backend/handler"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/database"
	"github.com/kokoichi206/awesome-chat-app/backend/repository/firebase"
	"github.com/kokoichi206/awesome-chat-app/backend/usecase"
	"github.com/kokoichi206/awesome-chat-app/backend/util"
	"github.com/kokoichi206/awesome-chat-app/backend/util/logger"
)

const (
	service = "awesome-chat-app"
)

func main() {
	// config
	cfg := config.New()

	// logger
	logger := logger.NewBasicLogger(os.Stdout, "ubuntu", service)

	// tracer
	tracer, traceCloser, err := util.NewJaegerTracer(cfg.AgentHost, cfg.AgentPort, service)
	if err != nil {
		logger.Errorf(context.Background(), "cannot initialize jaeger tracer: ", err)
	}
	defer traceCloser.Close()

	opentracing.SetGlobalTracer(tracer)

	// database
	database, err := database.New(
		cfg.DbDriver, cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword,
		cfg.DbName, cfg.DbSSLMode, logger,
	)
	if err != nil {
		logger.Errorf(context.Background(), "failed to db.New: ", err)
	}

	authClient, err := firebase.New(context.Background(), cfg.CredentialPath)
	if err != nil {
		logger.Errorf(context.Background(), "failed to firebase.New: ", err)
	}

	// usecase
	usecase := usecase.New(database, authClient, logger)

	// handler
	h := handler.New(logger, usecase)
	addr := net.JoinHostPort(cfg.ServerHost, cfg.ServerPort)

	// run
	if err := h.Engine.Run(addr); err != nil {
		logger.Critical(context.Background(), "failed to serve http")
	}
}
