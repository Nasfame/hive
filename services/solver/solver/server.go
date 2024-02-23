package solver

import (
	"context"
	"encoding/json"
	"fmt"
	corehttp "net/http"
	"time"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/services/solver/solver/store"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/http"
	"github.com/CoopHive/hive/pkg/system"
)

type solverServer struct {
	options    http.ServerOptions
	controller *SolverController
	store      store.SolverStore
}

func NewSolverServer(
	options http.ServerOptions,
	controller *SolverController,
	store store.SolverStore,
) (*solverServer, error) {
	server := &solverServer{
		options:    options,
		controller: controller,
		store:      store,
	}
	return server, nil
}

/*
 *
 *
 *

 Routes

 *
 *
 *
*/

func (solverServer *solverServer) ListenAndServe(ctx context.Context, cm *system.CleanupManager) error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix(config.API_SUB_PATH).Subrouter()

	subrouter.Use(http.CorsMiddleware)

	subrouter.HandleFunc("/job_offers", http.GetHandler(solverServer.getJobOffers)).Methods("GET")
	subrouter.HandleFunc("/job_offers", http.PostHandler(solverServer.addJobOffer)).Methods("POST")

	subrouter.HandleFunc("/resource_offers", http.GetHandler(solverServer.getResourceOffers)).Methods("GET")
	subrouter.HandleFunc("/resource_offers", http.PostHandler(solverServer.addResourceOffer)).Methods("POST")

	subrouter.HandleFunc("/deals", http.GetHandler(solverServer.getDeals)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}", http.GetHandler(solverServer.getDeal)).Methods("GET")

	subrouter.HandleFunc("/deals/{id}/files", solverServer.downloadFiles).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/files", solverServer.uploadFiles).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/result", http.GetHandler(solverServer.getResult)).Methods("GET")
	subrouter.HandleFunc("/deals/{id}/result", http.PostHandler(solverServer.addResult)).Methods("POST")

	subrouter.HandleFunc("/deals/{id}/txs/resource_provider", http.PostHandler(solverServer.updateTransactionsResourceProvider)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/job_creator", http.PostHandler(solverServer.updateTransactionsJobCreator)).Methods("POST")
	subrouter.HandleFunc("/deals/{id}/txs/mediator", http.PostHandler(solverServer.updateTransactionsMediator)).Methods("POST")

	// this will fan out to all connected web socket connections
	// we read all events coming from inside the solver controller
	// and write them to anyone who is connected to us
	websocketEventChannel := make(chan []byte)

	log.Debug().Msgf("begin solverServer.controller.subscribeEvents")
	go solverServer.controller.subscribeEvents(func(ev SolverEvent) {
		evBytes, err := json.Marshal(ev)
		if err != nil {
			log.Error().Msgf("Error marshalling event: %s", err.Error())
		}
		websocketEventChannel <- evBytes
	})

	log.Info().Msgf("start websocket server")

	// websocket server to send deals matched
	http.StartWebSocketServer(
		subrouter,
		config.WEBSOCKET_SUB_PATH,
		websocketEventChannel,
		ctx,
	)

	log.Info().Msgf("started websocket server")

	srv := &corehttp.Server{
		Addr:              fmt.Sprintf("%s:%d", solverServer.options.Host, solverServer.options.Port),
		WriteTimeout:      time.Minute * 15,
		ReadTimeout:       time.Minute * 15,
		ReadHeaderTimeout: time.Minute * 15,
		IdleTimeout:       time.Minute * 60,
		Handler:           router,
	}

	// Create a channel to receive errors from ListenAndServe
	serverErrors := make(chan error, 1)

	// Run ListenAndServe in a goroutine because it blocks
	go func() {
		log.Info().Msgf("Starting solver server on %s:%d", solverServer.options.Host, solverServer.options.Port)
		serverErrors <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		// Create a context with a timeout for the server to close
		log.Debug().Msgf("closing server gracefully after 3 seconds")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		// Attempt to gracefully shut down the server
		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("failed to stop server: %w", err)
		}

	case err := <-serverErrors:
		log.Error().Msgf("Error running solver server: %s", err.Error())
		return err
	}

	return nil
}
