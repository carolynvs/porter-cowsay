package main

import (
	"github.com/deislabs/porter-cowsay/pkg/cowsay"
	"github.com/spf13/cobra"
)

func buildSchemaCommand(m *cowsay.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintSchema()
		},
	}
	return cmd
}
