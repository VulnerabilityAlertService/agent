package yum

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strings"
)
import "github.com/VulnerabilityAlertService/vas/common"

type Yum struct { list []byte }

func (y Yum) ToArray() []map[string]string {
  lines := []map[string]string{}
  for index, line := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(y.list), -1) {
    if index <= 1 {
      continue
    }
    item := SliceLine(line)
    if item != nil {
      lines = append(lines, item)
    }
  }

  return lines
}

func SliceLine(line string) map[string]string {
  words  := strings.Fields(common.Cleansing(line))
  if len(words) == 3 && words[0] != "" {
    return map[string]string{"package": words[0], "version": words[1]}
  }
  return nil
}

func Get() []map[string]string {
  out, err := exec.Command("yum", "list", "installed").Output()
  if err != nil {
    fmt.Print("Command error!! `yum list installed` \n")
    os.Exit(1)
  }
  list := Yum{ out }
  return list.ToArray()
}

func Installed() bool {
  _, err := exec.Command("yum", "--version").Output()
  if err != nil {
    return false
  }
  return true
}
