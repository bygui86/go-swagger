// Package classification BlogPost API
//
// Here we provide a detailed overview of how to use BlogPost API.
//
// Terms Of Service:
//
//     Schemes: http
//     Host: localhost:8080
//     Version: 1.0.0
//     Contact: Matteo Baiguini<bygui86@github.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-swagger/go-swagger-rest-api/blogpost"
	"github.com/bygui86/go-swagger/go-swagger-rest-api/logging"
)

var restServer *blogpost.Server

func main() {
	logging.Log.Info("Start go-swagger-rest-api")

	restServer = startBlogPostServer()

	logging.Log.Info("go-swagger-rest-api up and running")

	startSysCallChannel()

	shutdownAndWait(3)
}

func startBlogPostServer() *blogpost.Server {
	logging.Log.Debug("Start BlogPost server")

	server := blogpost.NewRestServer()
	logging.Log.Debug("BlogPost server successfully created")

	server.Start()
	logging.Log.Debug("BlogPost server successfully started")

	return server
}

func startSysCallChannel() {
	syscallCh := make(chan os.Signal)
	signal.Notify(syscallCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-syscallCh
}

func shutdownAndWait(timeout int) {
	logging.SugaredLog.Warnf("Termination signal received! Timeout %d", timeout)

	if restServer != nil {
		restServer.Shutdown(timeout)
	}

	time.Sleep(time.Duration(timeout+1) * time.Second)
}
