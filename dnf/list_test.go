package dnf

import (
  "testing"
)

func TestDnf_Slice(t *testing.T) {
  list := "name.x86_64 1.2.3.fc28 @anaconda\n"
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
