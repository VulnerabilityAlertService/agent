package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

import (
  "github.com/VulnerabilityAlertService/vas/config"
  "github.com/VulnerabilityAlertService/vas/yum"
  "github.com/VulnerabilityAlertService/vas/dnf"
  "github.com/VulnerabilityAlertService/vas/aptget"
  "github.com/VulnerabilityAlertService/vas/apk"
  "github.com/VulnerabilityAlertService/vas/transport"
)

var cfgFile string

var token string

var RootCmd = &cobra.Command {
  Use:   "vas",
  Short: fmt.Sprint("Transport installed packages info(name & version) to V.A.S.(", transport.HOST ,")"),
  Run: func(cmd *cobra.Command, args []string) {
    if token == "" {
      token = config.Token()
    }
    if token == "" {
      fmt.Println("Please set a token key!")
      fmt.Println("You can set a token key! -> `vas config or vas -t yourtoken`")
      os.Exit(1)
    }

    values := getValues()

    if values == nil {
      fmt.Println("Couldn't get packages lists!")
      os.Exit(1)
    }
    transport.Post(transport.Dataset(values, token))
    fmt.Println("====== VulnerabilityAlertService ======")
    fmt.Println("------        Success!!          ------")
    fmt.Println("Check it out!", transport.HOST)
    fmt.Println("=======================================")
  },
}

func init() {
  RootCmd.Flags().StringVarP(&token, "token", "t", "", "set a token key")
}

func getValues() []map[string]string {
  var values []map[string]string
  if dnf.Installed() {
    values = append(values, dnf.Get()...)
  } else if yum.Installed() {
    values = append(values, yum.Get()...)
  }

  if aptget.Installed() {
    values = append(values, aptget.Get()...)
  }

  if apk.Installed() {
    values = append(values, apk.Get()...)
  }

  return values
}
