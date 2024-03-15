package in_test

import (
	"context"
	"fmt"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/internal"
	"github.com/CoopHive/hive/services"
	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/internal/jobCreatorService"
	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/executor/noop"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/pkg/web3"
	"github.com/CoopHive/hive/services/dealmaker"
	"github.com/CoopHive/hive/services/jobcreator"
	"github.com/CoopHive/hive/services/mediator"
	"github.com/CoopHive/hive/services/resourceprovider"
	"github.com/CoopHive/hive/services/solver"
	solver2 "github.com/CoopHive/hive/services/solver/solver"
	solvermemorystore "github.com/CoopHive/hive/services/solver/solver/store/memory"
)

type testOptions struct {
	moderationChance int
	executor         noop.NoopExecutorOptions
}

var dealMaker *dealmaker.Service

func getWeb3Options() (w *web3.Web3Options) {
	conf := config.Conf
	w = &web3.Web3Options{
		RpcURL: conf.GetString(enums.HIVE_RPC_WS),
	}

	return
}
func getSolver(t *testing.T, options testOptions) (*solver2.Solver, error) {
	solverOptions := solver.NewSolverOptions()

	//solverOptions.Web3 = *getWeb3Options()
	solverOptions.Web3.PrivateKey = os.Getenv("SOLVER_PRIVATE_KEY")
	solverOptions.Server.Port = 8080
	solverOptions.Server.URL = "http://localhost:8080"

	// test that the solver private key is defined
	if solverOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("SOLVER_PRIVATE_KEY is not defined")
	}

	web3SDK, err := web3.NewContractSDK(solverOptions.Web3)
	if err != nil {
		return nil, err
	}

	solverStore, err := solvermemorystore.NewSolverStoreMemory(config.Conf) // fx probably won't work fx injections TODO: to test FIXME
	if err != nil {
		return nil, err
	}

	return solver2.NewSolver(solverOptions, solverStore, web3SDK)
}

func getResourceProvider(
	t *testing.T,
	systemContext *system.CommandContext,
	options testOptions,
) (*resourceprovider.ResourceProvider, error) {
	resourceProviderOptions := resourceprovider.NewResourceProviderOptions()
	resourceProviderOptions.Web3.PrivateKey = os.Getenv("RP_PRIVATE_KEY")
	if resourceProviderOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("RESOURCE_PROVIDER_PRIVATE_KEY is not defined")
	}
	resourceProviderOptions, err := resourceprovider.ProcessResourceProviderOptions(resourceProviderOptions)
	if err != nil {
		return nil, err
	}

	web3SDK, err := web3.NewContractSDK(resourceProviderOptions.Web3)
	if err != nil {
		return nil, err
	}

	executor, err := noop.NewNoopExecutor(options.executor)
	if err != nil {
		return nil, err
	}

	// FIXME:pass dealmaker service
	return resourceprovider.NewResourceProvider(resourceProviderOptions, web3SDK, executor, dealMaker)
}

func getMediator(
	t *testing.T,
	systemContext *system.CommandContext,
	options testOptions,
) (*mediator.Mediator, error) {
	mediatorOptions := mediator.NewMediatorOptions()
	mediatorOptions.Web3.PrivateKey = os.Getenv("MEDIATOR_PRIVATE_KEY")
	if mediatorOptions.Web3.PrivateKey == "" {
		return nil, fmt.Errorf("MEDIATOR_PRIVATE_KEY is not defined")
	}
	mediatorOptions, err := mediator.ProcessMediatorOptions(mediatorOptions)
	if err != nil {
		return nil, err
	}

	web3SDK, err := web3.NewContractSDK(mediatorOptions.Web3)
	if err != nil {
		return nil, err
	}

	options.executor.BadActor = false

	executor, err := noop.NewNoopExecutor(options.executor)
	if err != nil {
		return nil, err
	}

	return mediator.NewMediator(mediatorOptions, web3SDK, executor)
}

func getJobCreatorOptions(options testOptions) (jobCreatorService.JobCreatorOptions, error) {
	jobCreatorOptions := jobcreator.NewJobCreatorOptions()
	jobCreatorOptions.Web3.PrivateKey = os.Getenv("JC_PRIVATE_KEY")
	if jobCreatorOptions.Web3.PrivateKey == "" {
		return jobCreatorOptions, fmt.Errorf("JC_PRIVATE_KEY is not defined")
	}
	ret, err := jobcreator.ProcessJobCreatorOptions(jobCreatorOptions, []string{
		// this should point to the shortcut
		"cowsay:v0.0.2",
	})

	if err != nil {
		return ret, err
	}

	jobCreatorOptions.Mediation.CheckResultsPercentage = options.moderationChance
	return ret, nil
}

func testStackWithOptions(
	t *testing.T,
	commandCtx *system.CommandContext,
	options testOptions,
) (*jobCreatorService.RunJobResults, error) {

	solver, err := getSolver(t, options)
	if err != nil {
		return nil, err
	}

	solver.Start(commandCtx.Ctx, commandCtx.Cm)

	// give the solver server a chance to boot before we get all the websockets
	// up and trying to connect to it
	time.Sleep(100 * time.Millisecond)

	resourceProvider, err := getResourceProvider(t, commandCtx, options)
	if err != nil {
		return nil, err
	}

	resourceProvider.Start(commandCtx.Ctx, commandCtx.Cm)

	mediator, err := getMediator(t, commandCtx, options)
	if err != nil {
		return nil, err
	}

	mediator.Start(commandCtx.Ctx, commandCtx.Cm)

	jobCreatorOptions, err := getJobCreatorOptions(options)
	if err != nil {
		return nil, err
	}

	result, err := jobCreatorService.RunJob(commandCtx, jobCreatorOptions, dealMaker, func(evOffer dto.JobOfferContainer) {

	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func TestNoModeration(t *testing.T) {
	commandCtx := system.NewTestingContext()
	defer commandCtx.Cleanup()

	initApp(t)

	message := "hello apples this is a message"

	executorOptions := noop.NewNoopExecutorOptions()
	executorOptions.Stdout = message

	result, err := testStackWithOptions(t, commandCtx, testOptions{
		moderationChance: 0,
		executor:         executorOptions,
	})

	assert.NoError(t, err, "there was an error running the job")
	assert.Equal(t, noop.NOOP_RESULTS_CID, result.Result.DataID, "the data ID was correct")

	localPath := solver2.GetDownloadsFilePath(result.Result.DealID)

	fmt.Printf("result --------------------------------------\n")
	spew.Dump(localPath)
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal().Str("err", err.Error()).Msgf(".env not found")
	}
	//
	//app := cmd.Hive()
	//
	//go func() {
	//	<-app.Done()
	//	log.Info().Str("app", "exiting gracefully")
	//}()

}

var ctx = context.Background()

func initApp(t *testing.T) {
	type In struct {
		DealmakerService *dealmaker.Service `name:"dealmaker"`
	}

	app := fxtest.New(t,
		config.Module,
		internal.Module,
		services.ModuleWithoutRoot,
		fx.Invoke(func(d *dealmaker.Service) {
			dealMaker = d
		}),
		//fx.Provide(func(i *In) {
		//	dealMaker = i.DealmakerService
		//}),
	)

	// Start the application.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	// defer cancel()
	// if err := app.Start(ctx); err != nil {
	// 	t.Fatal(err)
	// }

	app.RequireStart()

	go func() {
		<-app.Done()
		t.Log("exiting app gracefully")
	}()

	//// Stop the application.
	//if err := app.Stop(ctx); err != nil {
	//	t.Fatal(err)
	//}
}
