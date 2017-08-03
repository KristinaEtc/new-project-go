package main

// Function init() from this packege must execute first;
// do not move from the 1st position in import list.
import _ "github.com/KristinaEtc/slflog"

import (
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

	log.Error("----------------------------------------------")
	log.Infof("%s", globalOpt.Name)

	log.Infof("BuildDate=%s\n", BuildDate)
	log.Infof("GitCommit=%s\n", GitCommit)
	log.Infof("GitBranch=%s\n", GitBranch)
	log.Infof("GitState=%s\n", GitState)
	log.Infof("GitSummary=%s\n", GitSummary)
	log.Infof("VERSION=%s\n", Version)

	log.Info("Starting working...")

	tmpllib.PrintsTmpl("Yo")
}
