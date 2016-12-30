package main

import (
	log "github.com/Sirupsen/logrus"
)

func main(){
	fields := log.Fields{"splunk": 12, "akbar": true, "size": 234}()
	log.WithFields(fields).Info("doom!!")
}
