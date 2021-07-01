package main

import (
	"get.porter.sh/mixin/helm2/pkg/helm2"
	"github.com/spf13/cobra"
)

func buildUninstallCommand(m *helm2.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Execute the uninstall functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Uninstall()
		},
	}
	return cmd
}
