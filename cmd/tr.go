package cmd

import (
	"errors"

	"github.com/SignorMercurio/limner/transformer"
	"github.com/spf13/cobra"
)

var (
	inType  string
	outType string
)

// getTransformer returns a FormatTransformer
var getTransformer = func(inType, outType string) *transformer.FormatTransformer {
	return &transformer.FormatTransformer{
		InType:  inType,
		OutType: outType,
	}
}

// NewTrCmd represents the tr command which performs format tranformation
func NewTrCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tr",
		Short: "tr helps to transform data format",
		Long: `tr helps to transform data format
Example:
	curl -s https://api.github.com/users/SignorMercurio | lm tr -i json -o yaml ("-i json" can be omitted)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			if inType == outType {
				return errors.New("no need for transformation")
			}
			output, err = getTransformer(inType, outType).Transform(input)
			return err
		},
	}

	cmd.Flags().StringVarP(&inType, "in-type", "i", "", "The original format of data, currently supported format: json")
	cmd.Flags().StringVarP(&outType, "out-type", "o", "", "The desired format of data, currently supported format: yaml")

	return cmd
}

func init() {
	RootCmd.AddCommand(NewTrCmd())
}
