package math

import (
	"mytestproject/utils"
	"testing"
)

func TestAdd(t *testing.T) {
	utils.Log("Running Add test")
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("expected 3, got %d", result)
	}
}
