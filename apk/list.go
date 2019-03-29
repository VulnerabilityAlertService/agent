package apk

import (
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strings"
  "strconv"
)

type Apk struct { list []byte }

func (apk Apk) ToArray() []map[string]string {
  lines := []map[string]string{}
  for index, line := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(apk.list), -1) {
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
  var name []byte
  var version []byte
  for _, word := range words {
    char := string(getRuneAt(word, 0))
    _, err := strconv.Atoi(char)
    if len(version) == 0 && err != nil {
      name = append(name, []byte("-" + word)...)
    } else {
      version = append(version, []byte("-" + word)...)
    }
  }

  if len(name) > 0 {
    return map[string]string {
      "package": strings.TrimLeft(string(name), "-"),
      "version": strings.TrimLeft(string(version), "-"),
    }
  }

  return nil
}

func getRuneAt(s string, i int) rune {
  rs := []rune(s)
  return rs[i]
}

func Cleansing(line string) string {
  rep := strings.NewReplacer("-", " ")
  line = rep.Replace(line)
  return line
}

func Get() []map[string]string {
  out, err := exec.Command("apk", "info", "-v").Output()
  if err != nil {
    fmt.Print("Command error!! `apk info -v` \n")
    os.Exit(1)
  }
  list := Apk{ out }
  return list.ToArray()
}

func Installed() bool {
  _, err := exec.Command("apk", "info", "-V").Output()
  if err != nil {
    return false
  }
  return true
}
