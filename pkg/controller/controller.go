package controller

import (
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"github.com/bee-projects/azure-service-operator/pkg/azure"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var log = logf.Log.WithName("controller")

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager
var AddToManagerFuncs []func(manager.Manager, azure.Config) error

var azureConfig azure.Config

func init() {
	var err error
	azureConfig, err = azure.GetConfigFromEnvironment()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	log.Info("Azure config", "value", azureConfig)
}

// AddToManager adds all Controllers to the Manager
func AddToManager(m manager.Manager) error {
	for _, f := range AddToManagerFuncs {
		if err := f(m, azureConfig); err != nil {
			return err
		}
	}
	return nil
}
