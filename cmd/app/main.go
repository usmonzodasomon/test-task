package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	test_task "github.com/usmonzodasomon/test-task"
	"github.com/usmonzodasomon/test-task/internal/handler"
	"github.com/usmonzodasomon/test-task/internal/repository"
	"github.com/usmonzodasomon/test-task/internal/service"
	"github.com/usmonzodasomon/test-task/pkg/logger"
)

func main() {
	logg := logger.New()
	if err := godotenv.Load(); err != nil {
		logg.Error("Failed to load godotenv", logger.Err(err))
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	srv := new(test_task.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handler.InitRoutes()); !errors.Is(err, http.ErrServerClosed) && err != nil {
			logg.Error("Error occured while starting server: ", logger.Err(err))
			return
		}
	}()
	logg.Info("Server started...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	logg.Info("Server closed...")
	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Shutdown(ctx); err != nil {
		logg.Error("Error server shutting down: ", logger.Err(err))
		return
	}
}
