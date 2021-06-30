package main

import (
	"get.porter.sh/mixin/helm2/pkg/helm"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *helm.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the upgrade functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Upgrade()
		},
	}
	return cmd
}
