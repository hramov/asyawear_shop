package utils

import "testing"

func TestCutStrSuffix(t *testing.T) {
	str := "where name ilike 'John' and age >= '1' and age < '20' and "
	newStr := CutStrSuffix(str, 4)

	if newStr != "where name ilike 'John' and age >= '1' and age < '20' " {
		t.Errorf("wrong result: %s", newStr)
	}
}
