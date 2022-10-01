package presentation

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/pblanc5/headergetter/internal/opengraph"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HgApiConfig struct {
	Port int `mapstructure:"HG_PORT"`
}

func InitConfig() (HgApiConfig, error) {
	viper.SetDefault("HG_PORT", 8080)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	config := HgApiConfig{}
	err := viper.Unmarshal(&config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Unable to unmarshal config")
		return HgApiConfig{}, err
	}

	return config, nil
}

func RunApi() {

	logrus.SetFormatter(&logrus.JSONFormatter{})

	config, err := InitConfig()
	if err != nil {
		logrus.Fatal("Failed to load configuration: make sure hg.env exists or HG_PORT is defined")
	}

	r := mux.NewRouter()

	h := opengraph.OpenGraphHandler{
		Service: opengraph.OpenGraphService{
			Client: *http.DefaultClient,
		},
	}

	r.HandleFunc("/api/v1/og", h.GetOpenGraphTags).Methods("GET", "OPTIONS")
	r.Use(LoggingMiddleware)

	s := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", config.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		logrus.WithFields(logrus.Fields{
			"host": "0.0.0.0",
			"port": config.Port,
		}).WithTime(time.Now()).Info("Starting HeaderGetter server")

		if err := s.ListenAndServe(); err != nil {
			logrus.WithTime(time.Now()).Error(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	s.Shutdown(ctx)

	logrus.WithTime(time.Now()).Info("Shutting down HeaderGetter server")

	os.Exit(0)
}
