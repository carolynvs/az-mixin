package main

import (
	"github.com/deislabs/porter-az/pkg/az"
	"github.com/spf13/cobra"
)

var (
	commandFile string
)

func buildInstallCommand(m *az.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Execute the install functionality of this mixin",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			//Do something here if needed
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}
	return cmd
}
