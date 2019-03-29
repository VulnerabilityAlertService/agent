package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)
import "github.com/VulnerabilityAlertService/vas/config"

func init() {
  RootCmd.AddCommand(newConfigCmd())
}

func newConfigCmd() *cobra.Command {
  cmd := &cobra.Command {
    Use:   "config",
    Short: "Edit config",
    Run: func(cmd *cobra.Command, args []string) {
      config.Save(config.Read(), config.Input())
      fmt.Print("Update successfully! \n")
    },
  }

  return cmd
}
