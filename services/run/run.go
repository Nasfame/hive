package run

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/theckman/yacspin"

	"github.com/CoopHive/hive/config"
	"github.com/CoopHive/hive/enums"
	"github.com/CoopHive/hive/internal/genesis"
	"github.com/CoopHive/hive/internal/jobCreatorService"
	"github.com/CoopHive/hive/pkg/dto"
	"github.com/CoopHive/hive/pkg/system"
	"github.com/CoopHive/hive/services/dealmaker"
	"github.com/CoopHive/hive/services/solver/solver"
)

type service struct {
	dealMakerService *dealmaker.Service
	*genesis.Service
}

func (s *service) runJob(cmd *cobra.Command, options jobCreatorService.JobCreatorOptions, conf *viper.Viper) error {
	c := color.New(color.FgCyan).Add(color.Bold)
	header := `
  ___  __    __  ____  _  _  __  _  _  ____ 
 / __)/  \  /  \(  _ \/ )( \(  )/ )( \(  __)
( (__(  O )(  O )) __/) __ ( )( \ \/ / ) _) 
 \___)\__/  \__/(__)  \_)(_/(__) \__/ (____) v0

  Decentralized Compute Network  https://coophive.network


`
	version := conf.GetString(enums.VERSION)
	if version != "" {
		header = strings.Replace(header, "v0", version, 1)
	}
	c.Print(header)

	appName := s.Conf.GetString(enums.APP_NAME)

	spinner, err := createSpinner(appName+" submitting job", "🌟")
	if err != nil {
		s.Log.Fatalf("failed to make spinner from config struct: %v\n", err)
	}

	// start the spinner animation
	if err := spinner.Start(); err != nil {
		return fmt.Errorf("failed to start spinner: %w", err)
	}

	// update message
	// spinner.Message("uploading files")

	// let spinner render some more
	// time.Sleep(1 * time.Second)

	// if you wanted to print a failure message...
	//
	// if err := spinner.StopFail(); err != nil {
	// 	return fmt.Errorf("failed to stop spinner: %w", err)
	// }

	commandCtx := system.NewCommandContext(cmd)
	defer commandCtx.Cleanup()

	if options.Dealer != s.Conf.GetString(enums.DEALER) {
		if options.Dealer != config.DEFAULT_DEALER {
			if err := s.dealMakerService.LoadPlugin(options.Dealer); err != nil {
				s.Log.Errorf("failed to load dealer %s", options.Dealer)
			}
		} // TODO: should be refactored to the jobCreatorService
	}

	// TODO: inject this jobCreatorService to a service instead
	result, err := jobCreatorService.RunJob(commandCtx, options, s.dealMakerService, func(evOffer dto.JobOfferContainer) {
		if err := spinner.Stop(); err != nil {
			log.Fatalf("failed to stop spinner: %v", err)
		}
		st := dto.GetAgreementStateString(evOffer.State)
		var desc string
		var emoji string
		switch st {
		case "DealNegotiating":
			desc = "Job submitted. Negotiating deal..."
			emoji = "🤝"
		case "DealAgreed":
			desc = "Deal agreed. Running job..."
			emoji = "💌"
		case "ResultsSubmitted":
			//
			desc = "Results submitted. Awaiting verification..."
			emoji = "🤔"

		case "ResultsAccepted":
			desc = "Results accepted. Downloading result..."
			emoji = "✅"
		// 	jc call
		case "ResultsRejected":
			desc = "Results rejected! Getting refund..."
			emoji = "🙀"
		//
		default:
			desc = st
			emoji = "🌟"
		}
		spinner, err = createSpinner(desc, emoji)
		if err != nil {
			log.Fatalf("failed to make spinner from config struct: %v\n", err)
		}

		// start the spinner animation
		if err := spinner.Start(); err != nil {
			log.Fatalf("failed to start spinner: %s", err)
		}

		// UPDATE FUNCTION
		// fmt.Printf("evOffer: %s --------------------------------------\n")
		// spew.Dump(evOffer)

	})
	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}
	spinner.Stop()
	fmt.Printf("\n🍂 %s job completed, try 👇\n    open %s\n    cat %s/stdout\n    cat %s/stderr\n    https://ipfs.io/ipfs/%s\n",
		appName,
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		solver.GetDownloadsFilePath(result.JobOffer.DealID),
		result.Result.DataID,
	)
	return err
}

func createSpinner(message string, emoji string) (*yacspin.Spinner, error) {
	// build the configuration, each field is documented
	cfg := yacspin.Config{
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[69],
		Suffix:            " ", // puts a least one space between the animating spinner and the Message
		Message:           message,
		SuffixAutoColon:   true,
		ColorAll:          false,
		Colors:            []string{"fgMagenta"},
		StopCharacter:     emoji,
		StopColors:        []string{"fgGreen"},
		StopMessage:       message,
		StopFailCharacter: "✗",
		StopFailColors:    []string{"fgRed"},
		StopFailMessage:   "failed",
	}

	s, err := yacspin.New(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to make spinner from struct: %w", err)
	}

	stopOnSignal(s)
	return s, nil
}

func stopOnSignal(spinner *yacspin.Spinner) {
	// ensure we stop the spinner before exiting, otherwise cursor will remain
	// hidden and terminal will require a `reset`
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigCh

		spinner.StopFailMessage("interrupted")

		// ignoring error intentionally
		_ = spinner.StopFail()

		os.Exit(0)
	}()
}
