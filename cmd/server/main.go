package main

import (
	"fmt"
	logStd "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/juicypy/todo_list_service/src/config"
	"github.com/juicypy/todo_list_service/src/handlers"
	"github.com/juicypy/todo_list_service/src/infrastructure"
	"github.com/juicypy/todo_list_service/src/repo/dataservice"
	"github.com/juicypy/todo_list_service/src/server"
	authUC "github.com/juicypy/todo_list_service/src/usecase/auth"
	tasksUC "github.com/juicypy/todo_list_service/src/usecase/tasks"
)

func main() {

	cfg, err := config.ConfigFromEnv()
	if err != nil {
		logStd.Fatalf("error reading config: %s", err)
	}

	storageCfg, err := config.StorageConfigFromEnv()
	if err != nil {
		logStd.Fatalf("error reading config: %s", err)
	}
	_, goqu, err := infrastructure.ConnectDB(storageCfg)
	if err != nil {
		logStd.Fatalf("error reading config: %s", err)
	}

	logger, err := infrastructure.NewLogger(cfg.LogLevel)
	if err != nil {
		logStd.Fatalf("error creating logger: %s", err)
	}

	// repos
	userRepo := dataservice.NewUserDBRepo(goqu)
	tasksRepo := dataservice.NewTaskDBRepo(goqu)
	taskCommentsRepo := dataservice.NewTaskCommentsRepo(goqu)
	labelsRepo := dataservice.NewLabelsRepo(goqu)

	// usecases
	authUsecase := authUC.NewAuthUsecase(userRepo)
	tasksUsecase := tasksUC.NewTasksUseCase(tasksRepo, taskCommentsRepo, labelsRepo)

	//handlers
	auth := handlers.NewAuthHandler(authUsecase, *logger)
	users := handlers.NewUserHandler(authUsecase, *logger)
	tasks := handlers.NewTasksHandler(tasksUsecase, *logger)
	taskComments := handlers.NewTaskCommentsHandler(taskCommentsRepo, *logger)

	s := server.Server{Auth: auth, User: users, Tasks: tasks, TaskComments: taskComments}

	done := make(chan error)
	go func(done chan error) {
		done <- http.ListenAndServe(":9434", s.NewRouter())
		if err != nil {
			logger.Fatalf("error starting server: %v\n", err)
			return
		}
	}(done)

	logger.Info("------- SERVER STARTED --------")

	onExit(done, deferWithErr(logger.Sync))
}

func deferWithErr(f func() error) func() {
	return func() {
		logStd.Printf("error on exit: %s", f())
	}
}

func onExit(done chan error, run ...func()) {
	killSignal := make(chan os.Signal)
	signal.Notify(killSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case s := <-killSignal:
		fmt.Printf("exited with signal: %s", s.String())
	case err := <-done:
		fmt.Printf("exited with error: %s", err)
	}

	for _, r := range run {
		r()
	}
}
