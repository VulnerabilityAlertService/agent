package yum

import (
  "testing"
)

func TestYum_Slice(t *testing.T) {
  list := "name.noarch 1.2.3-45.el7 @CentOS\n"
  actual := SliceLine(list)
  expected := map[string]string{}
  expected = map[string]string {
               "package": "name",
               "version": "1.2.3-45",
             }
  if actual["package"] != expected["package"] {
    t.Errorf("got: %v\nwant: %v", actual["package"], expected["package"])
  }
  if actual["version"] != expected["version"] {
    t.Errorf("got: %v\nwant: %v", actual["version"], expected["version"])
  }
}
