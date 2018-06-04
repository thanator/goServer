package tests

import (
	"testing"

	"../model"
)

func TestReadingOrder(t *testing.T) {
	_, i, _ := model.FindOrderById(-1)
	if i != -1 {
		t.Errorf("ID was incorrect, got: %d, want: %d.", i, -1)
	}
}
