package commands

import (
	"fmt"

	"github.com/itsmikoj/mod-go/internal/gomod"
	"github.com/itsmikoj/mod-go/internal/registry"
	"github.com/spf13/cobra"
)

func NewInstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install [package]",
		Short: "Install a Go package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pkgName := args[0]

			url, err := registry.Resolve(pkgName)
			if err != nil {
				return err
			}

			if err := gomod.Install(url); err != nil {
				return fmt.Errorf("installation failed: %v", err)
			}

			fmt.Printf("Successfully installed %s (%s)\n", pkgName, url)
			return nil
		},
	}

	return cmd
}
