package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bygui86/go-swagger/swaggo-rest-api/blogpost"
	"github.com/bygui86/go-swagger/swaggo-rest-api/logging"
)

var restServer *blogpost.Server

// @title BlogPost API
// @version 1.0.0
// @description Here we provide a detailed overview of how to use BlogPost API.
// @termsOfService http://swagger.io/terms/
// @contact.name Matteo Baiguini
// @contact.email bygui86@github.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
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
