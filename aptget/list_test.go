package aptget

import (
  "testing"
)

func TestAptget_Cleansing(t *testing.T) {
  line := "ii name:amd64 1.2.3.ubuntu2 amd64 this is - test -\n"
  actual := Cleansing(line)
  expected := "ii name 1.2.3 amd64 this is - test -\n"
  if actual != expected {
    t.Errorf("got: %v\nwant: %v", actual, expected)
  }
}

func TestAptget_Slice(t *testing.T) {
  list := "ii name:amd64 1.2.3.ubuntu2 amd64 this is - test -\n"
  actual := SliceLine(list)
  expected := map[string]string{}
  expected = map[string]string {
               "package": "name",
               "version": "1.2.3",
             }
  if actual["package"] != expected["package"] {
    t.Errorf("got: %v\nwant: %v", actual["package"], expected["package"])
  }
  if actual["version"] != expected["version"] {
    t.Errorf("got: %v\nwant: %v", actual["version"], expected["version"])
  }
}
