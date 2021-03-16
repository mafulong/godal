package utils

import "testing"

func TestToJSON(t *testing.T) {
	t.Log(ToJSON("{}", true))
	t.Log(ToJSON("{}"))
}
