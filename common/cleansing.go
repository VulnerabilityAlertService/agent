package common

import (
  "regexp"
  "strings"
)

func Cleansing(line string) string {
  rep := strings.NewReplacer(".x86_64", "", ".noarch", "")
  line = rep.Replace(line)
  reg := regexp.MustCompile(`\.el7[\.\-a-z0-9]*`)
  line = reg.ReplaceAllString(line, "")
  reg  = regexp.MustCompile(`\.fc[\.\-a-z0-9]*`)
  line = reg.ReplaceAllString(line, "")
  return line
}
