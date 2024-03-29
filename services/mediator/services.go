package mediator

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/executor"
	"github.com/CoopHive/hive/pkg/executor/bacalhau"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
)

type MediatorOptions struct {
	Bacalhau bacalhau.BacalhauExecutorOptions
	Services dto.ServiceConfig
	Web3     web3.Web3Options
}

type Mediator struct {
	web3SDK    *web3.Web3SDK
	controller *MediatorController
}

func NewMediator(
	options MediatorOptions,
	web3SDK *web3.Web3SDK,
	executor executor.Executor,
) (*Mediator, error) {
	log.Debug().Msgf("begin NewMediatorController")
	controller, err := NewMediatorController(options, web3SDK, executor)
	log.Debug().Msgf("end NewMediatorController")
	if err != nil {
		log.Error().Msgf("error NewMediatorController")
		return nil, err
	}
	log.Debug().Msgf("begin Mediator")
	solver := &Mediator{
		controller: controller,
		web3SDK:    web3SDK,
	}
	return solver, nil
}

func (mediator *Mediator) Start(ctx context.Context, cm *system.CleanupManager) chan error {
	return mediator.controller.Start(ctx, cm)
}
