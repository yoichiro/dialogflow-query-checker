package check

import (
	"testing"
)

func TestHasSameChildMap(t *testing.T) {
	m := map[string]interface{}{}
	m["key1"] = "value1"
	m["key2"] = map[string]interface{}{
		"key3": "value3",
	}
	if hasSameChildMap("key1", m) {
		t.Fatal("The hasSameChildMap should return false against the string value entry")
	}
	if !hasSameChildMap("key2", m) {
		t.Fatal("The hasSameChildMap should return true against the map entry")
	}
}