package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/phob0s-pl/weatherllo/graphql"
	"github.com/phob0s-pl/weatherllo/owm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"time"
)

const (
	tokenFlag = "token"
	addrFlag  = "listen-address"
	debugFlag = "debug"
)

func listen(cmd *cobra.Command, args []string) {
	initializeLogger()
	log.WithField("version", owm.Version).Info("starting application")

	token, err := cmd.Flags().GetString(tokenFlag)
	if err != nil {
		log.WithField("error", err).Fatal("failed to parse flags")
	}

	addr, err := cmd.Flags().GetString(addrFlag)
	if err != nil {
		log.WithField("error", err).Fatal("failed to parse flags")
	}

	debug, err := cmd.Flags().GetBool(debugFlag)
	if err != nil {
		log.WithField("error", err).Fatal("failed to parse flags")
	}
	enableDebugLogs(debug)

	resolver := graphql.NewResolver(owm.New(token))
	router := mux.NewRouter()
	router.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))
	router.Handle("/", handler.Playground("weatherllo GraphQL playground", "/query"))

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 120,
		ReadTimeout:  time.Second * 120,
		IdleTimeout:  time.Second * 120,
		Handler:      router,
	}

	log.WithField("address", addr).Info("starting listener")

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.WithField("error", err).Fatal("listener failed")
	}
}

func initializeLogger() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func enableDebugLogs(enabled bool) {
	if enabled {
		log.SetLevel(log.DebugLevel)
		log.Debug("debug log level enabled")
	}

}
