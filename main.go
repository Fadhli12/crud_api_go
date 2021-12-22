package main

import (
	"crud_api_go/brands"
	"crud_api_go/config"
	"crud_api_go/migrations"
	"crud_api_go/utilities"
	"flag"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	version = flag.Int("version", 0, "Version")
	port    = flag.Int("port", 9090, "Port")
)

func main() {
	flag.Parse()
	utilities.Logs()
	if *version > 0 {
		if err := migrations.Migration(&config.Config.PostgreSQL, *version); err != nil {
			utilities.Logger.Panic(err)
		}
		os.Exit(0)
	}
	db, err := config.PostgreConnection(&config.Config.PostgreSQL)
	if err != nil {
		utilities.Logger.Panic(err)
	}
	defer db.Close()
	routes := mux.NewRouter()
	routes.Use(utilities.LogsMiddleware)

	brands.Routes(routes, config.PostgreConn)

	portStr := strconv.Itoa(*port)
	srv := &http.Server{
		Handler:      routes,
		Addr:         ":" + portStr,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	logrus.Infof("Server Running in Port : ", portStr)
	utilities.Logger.Fatal(srv.ListenAndServe())

}
