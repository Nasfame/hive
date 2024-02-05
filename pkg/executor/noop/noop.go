package noop

import (
	"fmt"
	"path/filepath"

	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/executor"
	"github.com/CoopHive/hive/pkg/system"
)

const RESULTS_DIR = "noop-results"

type NoopExecutorOptions struct {
	BadActor         bool
	ResultsCID       string
	Stdout           string
	Stderr           string
	ExitCode         string
	InstructionCount int
}

type NoopExecutor struct {
	Options NoopExecutorOptions
}

const NOOP_RESULTS_CID = "123"

func NewNoopExecutorOptions() NoopExecutorOptions {
	return NoopExecutorOptions{
		BadActor:         false,
		ResultsCID:       NOOP_RESULTS_CID,
		Stdout:           "Hello World!",
		Stderr:           "",
		ExitCode:         "0",
		InstructionCount: 1,
	}
}

func NewNoopExecutor(options NoopExecutorOptions) (*NoopExecutor, error) {
	return &NoopExecutor{
		Options: options,
	}, nil
}

func (e *NoopExecutor) RunJob(
	deal dto.DealContainer,
	module dto.Module,
) (*executor.ExecutorResults, error) {
	resultsDir, err := system.EnsureDataDir(filepath.Join(RESULTS_DIR, deal.ID))
	if err != nil {
		return nil, fmt.Errorf("error creating a local folder of results %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stdout"), []byte(e.Options.Stdout))
	if err != nil {
		return nil, fmt.Errorf("error creating stdout file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "stderr"), []byte(e.Options.Stdout))
	if err != nil {
		return nil, fmt.Errorf("error creating stderr file %s -> %s", deal.ID, err.Error())
	}
	err = system.WriteFile(filepath.Join(resultsDir, "exitCode"), []byte(e.Options.ExitCode))
	if err != nil {
		return nil, fmt.Errorf("error creating exitCode file %s -> %s", deal.ID, err.Error())
	}
	results := &executor.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       e.Options.ResultsCID,
		InstructionCount: e.Options.InstructionCount,
	}
	return results, nil
}

// Compile-time interface check:
var _ executor.Executor = (*NoopExecutor)(nil)
