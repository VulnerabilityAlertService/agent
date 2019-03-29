package apk

import (
  "testing"
)

func TestApk_Cleansing(t *testing.T) {
  line := "name-1.2.3-line\n"
  actual := Cleansing(line)
  expected := "name 1.2.3 line\n"
  if actual != expected {
    t.Errorf("got: %v\nwant: %v", actual, expected)
  }
}

func TestApk_Slice(t *testing.T) {
  list := "name-1.2.3-line\n"
  actual := SliceLine(list)
  expected := map[string]string{}
  expected = map[string]string {
               "package": "name",
               "version": "1.2.3-line",
             }
  if actual["package"] != expected["package"] {
    t.Errorf("got: %v\nwant: %v", actual["package"], expected["package"])
  }
  if actual["version"] != expected["version"] {
    t.Errorf("got: %v\nwant: %v", actual["version"], expected["version"])
  }
}
