package dnf

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strings"
)
import "github.com/VulnerabilityAlertService/vas/common"

type Dnf struct { list []byte }

func (d Dnf) ToArray() []map[string]string {
  lines := []map[string]string{}
  for index, line := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(d.list), -1) {
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
  out, err := exec.Command("dnf", "list", "installed").Output()
  if err != nil {
    fmt.Print("Command error!! `dnf list installed` \n")
    os.Exit(1)
  }
  list := Dnf{ out }
  return list.ToArray()
}

func Installed() bool {
  _, err := exec.Command("dnf", "--version").Output()
  if err != nil {
    return false
  }
  return true
}
