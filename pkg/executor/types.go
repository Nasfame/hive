package executor

import (
	"github.com/CoopHive/hive/pkg/dto"
)

type ExecutorResults struct {
	ResultsDir       string
	ResultsCID       string
	InstructionCount int
}

type Executor interface {
	// run the given job and return a local folder
	// that contains the results
	RunJob(
		deal dto.DealContainer,
		module dto.Module,
	) (*ExecutorResults, error)
}
