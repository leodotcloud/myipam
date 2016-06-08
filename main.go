package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/ipam"
	"github.com/rancher/rancher-ipam/driver"
	"log"
)

const (
	DefaultPort = "8766"
)

func init() {
	//TODO: Take this as a flag during startup
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {

	launchPort := DefaultPort
	logrus.Info("rancher ipam driver: starting on port ", launchPort)

	d := driver.RancherIPAMDriver{}

	h := ipam.NewHandler(&d)
	log.Fatal(h.ServeTCP("rancher-ipam", ":"+launchPort))

	logrus.Info("rancher ipam driver: exiting")
}
