package common

import (
  "testing"
)

func TestCommon_Cleansing(t *testing.T) {
  line := "shadow-utils.x86_64    2:4.5-9.fc28    @anaconda"
  actual := Cleansing(line)
  expected := "shadow-utils    2:4.5-9    @anaconda"
  if actual != expected {
    t.Errorf("got: %v\nwant: %v", actual, expected)
  }
}