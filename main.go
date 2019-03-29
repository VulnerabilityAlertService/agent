package main

import (
  "fmt"
  "os"
)
import "github.com/VulnerabilityAlertService/vas/cmd"

func main() {
  if err := cmd.RootCmd.Execute(); err != nil {
    fmt.Printf("%+v\n", err)
    os.Exit(1)
  }
}