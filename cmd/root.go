package cmd

import (
	"github.com/Joffref/genius/pkg/uuid"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "genius",
		Short: "genius is a tool to generate a UUID",
		Long:  "genius is a tool that generates a UUID following UUIDv4 specifications.",
		RunE: func(cmd *cobra.Command, args []string) error {
			u, err := uuid.New()
			if err != nil {
				return err
			}
			cmd.Println(u)
			return nil
		},
	}
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		panic(err)
	}
}
