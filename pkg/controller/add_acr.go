package controller

import (
	"github.com/techniumlabs/azure-service-operator/pkg/controller/acr"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, acr.Add)
}
