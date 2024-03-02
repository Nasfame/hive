package bacalhau

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	bacalhau2 "github.com/CoopHive/hive/pkg/bacalhau"
	"github.com/CoopHive/hive/pkg/dto"
	executorlib "github.com/CoopHive/hive/pkg/executor"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/utils"
)

type BacalhauExecutorOptions struct {
	ApiHost string
}

type BacalhauExecutor struct {
	Options     BacalhauExecutorOptions
	bacalhauEnv []string
}

func NewBacalhauExecutor(options BacalhauExecutorOptions) (*BacalhauExecutor, error) {
	additionalBacalhauEnv := config.Conf.GetStringSlice(enums.BACALHAU_ENV)

	log.Debug().Msgf("additionalBacalhauEnv: %q+", additionalBacalhauEnv)

	bacalhauEnv := []string{
		fmt.Sprintf("BACALHAU_API_HOST=%s", options.ApiHost),
		fmt.Sprintf("HOME=%s", config.Conf.GetString(enums.APP_DIR)),
	}

	bacalhauEnv = append(bacalhauEnv, additionalBacalhauEnv...)

	log.Debug().Msgf("bacalhauEnv: %s", bacalhauEnv)
	return &BacalhauExecutor{
		Options:     options,
		bacalhauEnv: bacalhauEnv,
	}, nil
}

func (executor *BacalhauExecutor) RunJob(
	deal dto.DealContainer,
	module dto.Module,
) (*executorlib.ExecutorResults, error) {
	id, err := executor.getJobID(deal, module) // runs the job and returns the job ID
	if err != nil {
		log.Error().Err(err).Msg("Run job")
		return nil, err
	}

	resultsDir, err := executor.copyJobResults(deal.ID, id)
	if err != nil {
		log.Error().Err(err).Msg("Run job")
		return nil, err
	}

	jobState, err := executor.getJobState(deal.ID, id)
	if err != nil {
		log.Error().Err(err).Msg("Run job")
		return nil, err
	}

	if len(jobState.State.Executions) == 0 {
		err := fmt.Errorf("no executions found for job %s", id)
		log.Error().Err(err).Msg("Run job")
		return nil, err
	}

	if jobState.State.State != bacalhau2.JobStateCompleted {
		err := fmt.Errorf("job %s did not complete successfully: %s", id, jobState.State.State.String())
		log.Error().Err(err).Msgf("job %s did not complete successfully", id)
		return nil, err
	}

	// TODO: we should think about WASM and instruction count here
	results := &executorlib.ExecutorResults{
		ResultsDir:       resultsDir,
		ResultsCID:       jobState.State.Executions[0].PublishedResult.CID,
		InstructionCount: 1,
	}

	return results, nil
}

// run the bacalhau job and return the job ID
func (executor *BacalhauExecutor) getJobID(
	deal dto.DealContainer,
	module dto.Module,
) (string, error) {
	// get a JSON string of the job
	jsonBytes, err := json.Marshal(module.Job)
	if err != nil {
		log.Error().Err(err).Msgf("error marshalling job JSON for deal %s -> %s", deal.ID, err.Error())
		return "", fmt.Errorf("error getting job JSON for deal %s -> %s", deal.ID, err.Error())
	}

	p := filepath.Join(config.Conf.GetString(enums.BACALHAU_SPECS_DIR), deal.ID)
	bacalhauJobSpecDir, err := utils.EnsureDir(p)
	if err != nil {
		log.Error().Err(err).Msgf("error creating a local folder for job specs %s -> %s", deal.ID, err.Error())
		return "", fmt.Errorf("error creating a local folder for job specs %s -> %s", deal.ID, err.Error())
	}
	jobPath := filepath.Join(bacalhauJobSpecDir, "job.json")
	err = system.WriteFile(jobPath, jsonBytes)
	if err != nil {
		log.Error().Err(err).Msgf("error writing job JSON %s -> %s", deal.ID, err.Error())
		return "", fmt.Errorf("error writing job JSON %s -> %s", deal.ID, err.Error())
	}

	// TODO:  try to use official bacalhau pkg
	runCmd := exec.Command(
		config.Conf.GetString(enums.BACALHAU_BIN),
		"create",
		"--id-only",
		"--wait",
		jobPath,
	)
	runCmd.Env = executor.bacalhauEnv
	log.Debug().Msgf("cmd: %+v", runCmd)

	runOutput, err := runCmd.CombinedOutput()
	log.Error().Err(err).Msgf("runOutput: %s", runOutput)
	if err != nil {
		log.Error().Err(err).Msgf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
		return "", fmt.Errorf("error running command %s -> %s, %s", deal.ID, err.Error(), runOutput)
	}

	id := strings.TrimSpace(string(runOutput))
	fmt.Printf("Got bacalhau job ID: %s\n", id)
	log.Info().Msgf("bacalhau jobID: %s\n", id)

	return id, nil
}

func (executor *BacalhauExecutor) copyJobResults(dealID string, jobID string) (string, error) {
	resultsDir, err := utils.EnsureDir(filepath.Join(config.Conf.GetString(enums.BACALHAU_RESULTS_DIR), dealID))
	if err != nil {
		log.Error().Err(err).Msgf("error creating a local folder of results %s -> %s", dealID, err.Error())
		return "", fmt.Errorf("error creating a local folder of results %s -> %s", dealID, err.Error())
	}

	copyResultsCmd := exec.Command(
		config.Conf.GetString(enums.BACALHAU_BIN),
		"get",
		jobID,
		"--output-dir", resultsDir,
	)
	copyResultsCmd.Env = executor.bacalhauEnv

	log.Debug().Msgf("cmd: %+v", copyResultsCmd)

	output, err := copyResultsCmd.CombinedOutput()
	log.Debug().Err(err).Msgf("output:%s", output)

	if err != nil {
		log.Error().Err(err).Msgf("error copying results %s -> %s", dealID, err)
		return "", fmt.Errorf("error copying results %s -> %s", dealID, err)
	}

	return resultsDir, nil
}

func (executor *BacalhauExecutor) getJobState(dealID string, jobID string) (*bacalhau2.JobWithInfo, error) {
	describeCmd := exec.Command(
		config.Conf.GetString(enums.BACALHAU_BIN),
		"describe",
		"--json",
		jobID,
	)
	describeCmd.Env = executor.bacalhauEnv

	log.Debug().Msgf("cmd: %+v", describeCmd)

	output, err := describeCmd.CombinedOutput()
	log.Debug().Err(err).Msgf("output:%s", output)
	if err != nil {
		log.Error().Err(err).Msgf("error calling describe command results %s -> %s", dealID, err.Error())
		return nil, fmt.Errorf("error calling describe command results %s -> %s", dealID, err.Error())
	}

	var job bacalhau2.JobWithInfo
	err = json.Unmarshal(output, &job)
	if err != nil {
		log.Error().Err(err).Msgf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
		return nil, fmt.Errorf("error unmarshalling job JSON %s -> %s", dealID, err.Error())
	}

	return &job, nil
}

// Compile-time interface check:
var _ executorlib.Executor = (*BacalhauExecutor)(nil)
