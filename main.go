package main

import (
	"gitea.repetitra.ru/StudBank/Backend/api"
	"gitea.repetitra.ru/StudBank/Backend/etc"
	"github.com/sirupsen/logrus"
)

func main() {
	gendb() // On start is empty

	etc.InitConfig()

	log := etc.GetLogger("main", logrus.TraceLevel)

	// GET DB

	log.Debug("Initializing api")
	api, err := (&api.API{}).Init()
	if err != nil {
		log.WithError(err).Fatal("Cant init API")
	}

	log.Debug("Running api")
	log.WithError(api.MRun()).Fatal("API exited")
}
