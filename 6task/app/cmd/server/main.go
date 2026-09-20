package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	httpadapter "web-app/internal/adapters/http"
	"web-app/internal/adapters/metrics"
	"web-app/internal/application"
	"web-app/internal/domain/thing"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelInfo,
			},
		),
	)

	slog.SetDefault(logger)

	metricsAdapter := metrics.NewPrometheus()

	domainService := thing.NewService()

	thingService := application.NewThingService(
		domainService,
		metricsAdapter,
	)

	handler := httpadapter.NewHandler(thingService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /things", handler.CreateThing)
	mux.HandleFunc("GET /things", handler.GetAll)

	mux.Handle(
		"/metrics",
		promhttp.Handler(),
	)

	log.Println("server started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
