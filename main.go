package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	"net/http"
	"os"
)

const serviceName = "D3 WebService"
const serviceDescription = "D3 web interface"

type program struct{}

type HomePageData struct {
	Version string
}

func main() {
	logInfo("MAIN", serviceName+" starting...")
	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		logError("MAIN", "Cannot start: "+err.Error())
	}
	err = s.Run()
	if err != nil {
		logError("MAIN", "Cannot start: "+err.Error())
	}
}

func (p *program) Start(service.Service) error {
	logInfo("MAIN", serviceName+" started")
	go p.run()
	return nil
}

func (p *program) Stop(service.Service) error {
	logInfo("MAIN", serviceName+" stopped")
	return nil
}

func (p *program) run() {
	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.GET("/", homepage)
	err := http.ListenAndServe(":80", router)
	if err != nil {
		logError("MAIN", "Problem starting service: "+err.Error())
		os.Exit(-1)
	}
	logInfo("MAIN", serviceName+" running")
}
