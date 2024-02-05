package jobCreatorService

import (
	"fmt"
	"time"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/module"
)

// this will load the module in the offer options
// and hoist the machine spec from the module into the offer
func getJobOfferFromOptions(options JobCreatorOfferOptions, jobCreatorAddress string) (dto.JobOffer, error) {
	// process the given module so we know what spec the job is asking for
	// this will also validate the module the user is asking for
	loadedModule, err := module.LoadModule(options.Module, options.Inputs)
	if err != nil {
		return dto.JobOffer{}, fmt.Errorf("error loading module: %s opts=%+v", err.Error(), options)
	}

	return dto.JobOffer{
		// assign CreatedAt to the current millisecond timestamp
		CreatedAt:  int(time.Now().UnixNano() / int64(time.Millisecond)),
		JobCreator: jobCreatorAddress,
		Module:     options.Module,
		Spec:       loadedModule.Machine,
		Inputs:     options.Inputs,
		Mode:       options.Mode,
		Pricing:    options.Pricing,
		Timeouts:   options.Timeouts,
		Services:   options.Services,
	}, nil
}
