package main

//Notes:
//MySQL user is admin, this is a bad practice but kept because the ip2location docker image has this user as default.
//The MySQL user password should be delivered by the orchestrator
//In config, the cert files are stored, those shouldn't be stored in any repository but for practical reasons they will be kept

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/nullc0rp/go-ip2proxy-api/config"
	"github.com/nullc0rp/go-ip2proxy-api/controller"
	"github.com/nullc0rp/go-ip2proxy-api/database"
	"github.com/nullc0rp/go-ip2proxy-api/service"
)

func main() {

	//Get configuration
	configuration := config.GetConfig()

	//Define services
	var serviceInstance service.Service
	var controllerInstance controller.Controller
	var databaseInstance database.Database

	//Create database connection TODO: use env variables
	databaseInstance = &database.DatabaseImpl{
		Server:   configuration.DBHOST,
		User:     configuration.DBUSERNAME,
		Password: configuration.DBPASSWORD,
		Database: configuration.DBNAME,
	}

	//Start connection. If it fails, the server should not be operational
	databaseInstance.Connect()

	//Instance Service
	serviceInstance = &service.ServiceImp{
		DB: databaseInstance,
	}

	//Instance Controller
	controllerInstance = &controller.ControllerImpl{
		Service: serviceInstance,
	}

	//Create Server and Route Handlers
	r := mux.NewRouter()

	//Define paths
	r.HandleFunc("/ip/{address:.*}", controllerInstance.GetIpInfo).Methods("GET")
	r.HandleFunc("/country/{country:[A-Z]+}", controllerInstance.GetIpList).Methods("GET")
	r.HandleFunc("/country/{country:[A-Z]+}/isp", controllerInstance.GetISPCountry).Methods("GET")
	r.HandleFunc("/country/{country:[A-Z]+}/total", controllerInstance.GetIPTotalCountry).Methods("GET")
	r.HandleFunc("/proxytypes", controllerInstance.GetMostProxyTypes).Methods("GET")

	//Create http service TODO: use config variables
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8443",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServeTLS("./config/certs/server.crt", "./config/certs/server.key"); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
