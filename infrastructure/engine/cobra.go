package engine

import "github.com/spf13/cobra"

type CobraCommand interface {
	Cobra() *cobra.Command
}
