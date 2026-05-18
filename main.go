package main

import (
	"context"

	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

var model = resource.NewModel("cecilyViam", "build-test", "stub")

func main() {
	resource.RegisterComponent(generic.API, model, resource.Registration[resource.Resource, resource.NoNativeConfig]{
		Constructor: newStub,
	})
	module.ModularMain(resource.APIModel{generic.API, model})
}

type stub struct {
	resource.Named
	resource.AlwaysRebuild
	resource.TriviallyCloseable
}

func newStub(_ context.Context, _ resource.Dependencies, conf resource.Config, _ logging.Logger) (resource.Resource, error) {
	return &stub{Named: conf.ResourceName().AsNamed()}, nil
}

func (s *stub) DoCommand(_ context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return cmd, nil
}
