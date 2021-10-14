package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/SignorMercurio/limner/printer"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	plainMode bool
	mustType  string
	input     []byte
	output    []byte
	Stdout    = colorable.NewColorableStdout()
	RootCmd   = NewRootCmd()
	Log       = logrus.New()
)

// getPrinter returns a ColorPrinter
var getPrinter = func(mustType string) *printer.ColorPrinter {
	return &printer.ColorPrinter{
		Type: mustType,
	}
}

// NewRootCmd represents the base command when called without any subcommands
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "lm",
		Short:         "Limner colorizes and transforms CLI outputs",
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			input, err = io.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			output = input
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			var p printer.Printer
			if plainMode {
				p = &printer.PlainPrinter{}
			} else {
				p = getPrinter(mustType)
			}
			p.Print(string(output), Stdout)
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.limner.yaml)")
	cmd.PersistentFlags().BoolVarP(&plainMode, "plain", "p", false, "Do not colorize the output")
	cmd.PersistentFlags().StringVarP(&mustType, "type", "t", "", "Force limner to view the output as a specific type: yaml / json / xml / table, etc.")

	return cmd
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
	cobra.OnInitialize(initConfig, printer.InitColorTheme)
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
