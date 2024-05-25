package main

import (
	http2 "net/http"
	"simple-job-board/internal/http"
	"simple-job-board/internal/job"
	db "simple-job-board/internal/sqlc"
	"simple-job-board/pkg/config"
	"simple-job-board/pkg/database"

	log "github.com/sirupsen/logrus"
)

const (
	BaseServerPath  = "/api/v1"
	HealthCheckPath = "/healthcheck"

	// offers
	ListOfferPath   = "/offers"
	DetailOfferPath = "/offer/{offerId}"
	ApplyOfferPath  = "/offer/{offerId}/apply"

	// employers
	ListEmployerPath   = "/employers"
	DetailEmployerPath = "/employer/{employerId}"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	// Load configuration
	config, err := config.NewConfiguration("config")
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	conn, err := database.GetDbConn(config.PostgresUser, config.PostgresPassword, config.PostgresHost, config.PostgresPort, config.PostgresDatabase, config.PostgresSSLMode)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := db.New(conn)
	offerHandler := job.NewOfferHandler(job.NewJobService(*repo))

	server := http.NewHttpServer(config)
	server.RegisterRouter(HealthCheckPath, func(w http2.ResponseWriter, r *http2.Request) {
		w.WriteHeader(http2.StatusOK)
		w.Write([]byte("OK"))
	})
	server.GroupRouter(BaseServerPath, func(r *http2.ServeMux) {
		// offer
		r.HandleFunc(ListOfferPath, offerHandler.ListOffers)
		r.HandleFunc(DetailOfferPath, offerHandler.GetOffer)
		r.HandleFunc(ApplyOfferPath, offerHandler.CreateOffer)
	})
	server.Start()
}
