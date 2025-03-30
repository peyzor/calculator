package main

import (
	"fmt"
	"github.com/peyzor/calculator/middleware"
	"github.com/rs/cors"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := http.NewServeMux()

	router.HandleFunc("POST /add", handleAdd)
	router.HandleFunc("POST /subtract", handleSubtract)
	router.HandleFunc("POST /multiply", handleMultiply)
	router.HandleFunc("POST /divide", handleDivide)
	router.HandleFunc("POST /sum", handleSum)

	srv := http.Server{
		Addr: ":8080",
		Handler: cors.AllowAll().Handler(
			middleware.Logging(logger, router),
		),
	}

	logger.Info(fmt.Sprintf("server listening on %s", srv.Addr))

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
