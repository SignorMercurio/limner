package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/SignorMercurio/limner/color"
	"github.com/SignorMercurio/limner/printer"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	plainMode bool
	lightBg   bool
	mustType  string
	Stdout    = colorable.NewColorableStdout()
	Stderr    = colorable.NewColorableStderr()
	RootCmd   = NewRootCmd()
	Log       = logrus.New()
)

type Printers struct {
	ColorPrinter printer.Printer
	ErrorPrinter printer.Printer
}

// getPrinters return a pair of printers
var getPrinters = func(mustType string, lightBg bool, args []string) *Printers {
	return &Printers{
		ColorPrinter: &printer.ColorPrinter{
			Type:    mustType,
			LightBg: lightBg,
			Args:    args,
		},
		ErrorPrinter: &printer.CustomPrinter{
			ColorPicker: func(line string) color.Color {
				if strings.HasPrefix(strings.ToLower(line), "error") {
					return color.Red
				}
				return color.Yellow
			},
		},
	}
}

// NewRootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "lm",
		Short:         "Limner colorizes and transforms CLI outputs",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			actualCmd := args[0]
			actualArgs := args[1:]
			// after --, there are args
			command := exec.Command(actualCmd, actualArgs...)
			command.Stdin = os.Stdin

			if plainMode {
				command.Stdout = Stdout
				command.Stderr = Stderr
				if err := command.Start(); err != nil {
					return err
				}

				return waitForExitCode(command, actualCmd)
			}

			cmdOut, err := command.StdoutPipe()
			if err != nil {
				return err
			}
			cmdErr, err := command.StderrPipe()
			if err != nil {
				return err
			}

			buf := new(bytes.Buffer)
			cmdOutReader := io.TeeReader(cmdOut, buf)
			cmdErrReader := io.TeeReader(cmdErr, buf)
			if err := command.Start(); err != nil {
				return err
			}

			printers := getPrinters(mustType, lightBg, actualArgs)

			wg := &sync.WaitGroup{}
			wg.Add(2)
			go func() {
				defer wg.Done()
				printers.ColorPrinter.Print(cmdOutReader, Stdout)
			}()
			go func() {
				defer wg.Done()
				printers.ErrorPrinter.Print(cmdErrReader, Stderr)
			}()
			wg.Wait()

			return waitForExitCode(command, actualCmd)
		},
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.limner.yaml)")
	cmd.PersistentFlags().BoolVar(&plainMode, "plain", false, "Do not colorize the output")
	cmd.PersistentFlags().BoolVar(&lightBg, "light-bg", false, "Adapt a more suitable color theme in a terminal with light background")
	cmd.PersistentFlags().StringVarP(&mustType, "type", "t", "", "Force limner to view the output as a specific type: yaml / json / xml / table, etc.")

	return cmd
}

func waitForExitCode(cmd *exec.Cmd, actualCmd string) error {
	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s error", actualCmd))
	}
	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer func() { recover() }()
	err := RootCmd.Execute()
	if err != nil {
		Log.Panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".limner" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".limner")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
