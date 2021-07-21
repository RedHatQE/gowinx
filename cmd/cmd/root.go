package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/RedHatQE/gowinx/pkg/util"
	"github.com/RedHatQE/gowinx/pkg/util/logging"
	"github.com/spf13/cobra"
	"k8s.io/utils/exec"
)

const (
	commandName      = "gowinx"
	descriptionShort = "manage windows ux"
	descriptionLong  = "manage windowx ux from cmdline"

	defaultErrorExitCode = 1
)

var (
	baseDir = filepath.Join(util.GetHomeDir(), ".gowinx")
	logFile = "gowinx.log"
)

var rootCmd = &cobra.Command{
	Use:   commandName,
	Short: descriptionShort,
	Long:  descriptionLong,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return runPrerun(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		runRoot()
		_ = cmd.Help()
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func runPrerun(cmd *cobra.Command) error {
	logging.InitLogrus(logging.LogLevel, baseDir, logFile)
	return nil
}

func runRoot() {
	fmt.Println("No command given")
}

func Execute() {
	attachMiddleware([]string{}, rootCmd)

	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		runPostrun()
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		var e exec.CodeExitError
		if errors.As(err, &e) {
			os.Exit(e.ExitStatus())
		} else {
			os.Exit(defaultErrorExitCode)
		}
	}
	runPostrun()
}

func attachMiddleware(names []string, cmd *cobra.Command) {
	if cmd.HasSubCommands() {
		for _, command := range cmd.Commands() {
			attachMiddleware(append(names, cmd.Name()), command)
		}
	} else if cmd.RunE != nil {
		fullCmd := strings.Join(append(names, cmd.Name()), " ")
		src := cmd.RunE
		cmd.RunE = executeWithLogging(fullCmd, src)
	}
}

func executeWithLogging(fullCmd string, input func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		logging.Infof("Running '%s'", fullCmd)
		return input(cmd, args)
	}
}

func runPostrun() {
	logging.CloseLogging()
}
