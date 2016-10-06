package main

import (
	_ "github.com/KristinaEtc/slflog"

	"github.com/KristinaEtc/new-project-go/tmplib"

	"github.com/KristinaEtc/config"
	"github.com/ventu-io/slf"
)

var log = slf.WithContext("template.go")

var (
	// These fields are populated by govvv
	BuildDate  string
	GitCommit  string
	GitBranch  string
	GitState   string
	GitSummary string
	Version    string
)

// ConfFile is a file with all program options
type ConfFile struct {
	Name string
}

var globalOpt = ConfFile{
	Name: "config",
}

func main() {

	config.ReadGlobalConfig(&globalOpt, "template options")

	log.Infof("%s", globalOpt.Name)

	log.Error("----------------------------------------------")

	log.Infof("BuildDate=%s\n", BuildDate)
	log.Infof("GitCommit=%s\n", GitCommit)
	log.Infof("GitBranch=%s\n", GitBranch)
	log.Infof("GitState=%s\n", GitState)
	log.Infof("GitSummary=%s\n", GitSummary)
	log.Infof("VERSION=%s\n", Version)

	log.Info("Starting working...")

	tmpllib.PrintsTmpl("Yo")
}
