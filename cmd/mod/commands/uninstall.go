package commands

import (
	"fmt"

	"github.com/itsmikoj/mod-go/internal/gomod"
	"github.com/itsmikoj/mod-go/internal/registry"
	"github.com/spf13/cobra"
)

func NewUninstallCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall [package]",
		Short: "Uninstall a Go package",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			pkgName := args[0]

			url, err := registry.Resolve(pkgName)
			if err != nil {
				return err
			}

			if err := gomod.Uninstall(url); err != nil {
				return fmt.Errorf("uninstallation failed: %v", err)
			}

			gomod.Tidy()

			return nil
		},
	}
	return cmd
}
