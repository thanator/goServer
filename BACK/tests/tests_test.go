package tests

import (
	"testing"

	"../db"
	"../model"
	_ "github.com/lib/pq"
)

func TestReadingOrder(t *testing.T) {
	_, i, err := model.FindOrderById(-1)
	if i != -1 && err != nil {
		t.Errorf("ID was incorrect, got: %d, want: %d.", i, -1)
	}
}

func TestCustomerGareway1(t *testing.T) {
	i, err := db.ReadCustomerByPhone("NO PHONE")
	if i != -1 && err != nil {
		t.Errorf("ID was incorrect, got: %d, want: %d.", i, -1)
	}
}

func TestCustomerGareway2(t *testing.T) {
	i, err := db.ReadCustomerByPhone("9234234243")
	if i == -1 || err != nil {
		t.Errorf("ID was incorrect, got: %d, want: %d.", i, -1)
	}
}

func TestBossMethod(t *testing.T) {
	iActual := model.FindAllOrderIds()
	if iActual[0] == -1 {
		t.Errorf("ID was incorrect, got: %d, not wanted: %d.", iActual[0], -1)
	}
}

func TestProductMethod(t *testing.T) {
	i, err := db.ReadProductByParams("-1", "101", "Test")
	if i != -1 && err == nil {
		t.Errorf("ID was incorrect, got: %d, not wanted: %d.", i, -1)
	}
}

func TestReadProduct(t *testing.T) {
	str := db.ReadProductById(2)
	if len(str) < 1 {
		t.Errorf("ID was incorrect, got: %d, not wanted: %d.", len(str), -1)
	}
}
