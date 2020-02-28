package controller

import (
	"github.com/NautiluX/presentation-example-operator/pkg/controller/presentation"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, presentation.Add)
}
