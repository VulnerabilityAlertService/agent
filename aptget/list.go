package aptget

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strings"
)

type Aptget struct { list []byte }

func (apt Aptget) ToArray() []map[string]string {
  lines := []map[string]string{}
  for index, line := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(apt.list), -1) {
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
  words := strings.Fields(Cleansing(line))
  if len(words) >= 5 && words[0] == "ii" && words[1] != "" {
    return map[string]string{"package": words[1], "version": words[2]}
  }
  return nil
}

func Cleansing(line string) string {
  rep := strings.NewReplacer(":amd64", "")
  line = rep.Replace(line)
  reg := regexp.MustCompile(`[\.]*ubuntu[\.0-9a-z]*`)
  line = reg.ReplaceAllString(line, "")
  return line
}

func Get() []map[string]string {
  out, err := exec.Command("dpkg", "-l").Output()
  if err != nil {
    fmt.Print("Command error!! `dpkg -l` \n")
    os.Exit(1)
  }
  list := Aptget{ out }
  return list.ToArray()
}

func Installed() bool {
  _, err := exec.Command("dpkg", "--version").Output()
  if err != nil {
    return false
  }
  return true
}
