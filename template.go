package main

// Function init() from this packege must execute first;
// do not move from the 1st position in import list.
import _ "github.com/KristinaEtc/slflog"

import (
	"github.com/KristinaEtc/new-project-go/tmplib"

	"github.com/KristinaEtc/config"
	"github.com/KristinaEtc/slf"
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

	log.Infof("BuildDate=%s", BuildDate)
	log.Infof("GitCommit=%s", GitCommit)
	log.Infof("GitBranch=%s", GitBranch)
	log.Infof("GitState=%s", GitState)
	log.Infof("GitSummary=%s", GitSummary)
	log.Infof("VERSION=%s", Version)

	log.Info("Starting working...")

	tmpllib.PrintsTmpl("\nYo")
}
