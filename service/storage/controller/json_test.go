package controller

import (
	"testing"

	//. "github.com/xiaoping378/blockchain-on-sql/parser"
	"fmt"
)
func TestMapToObject(t *testing.T) {
	fmt.Print("dddd")
}
/*
func TestParseQuantity(t *testing.T) {
	i, err := ParseQuantity("0x41")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 65 {
		t.Errorf("Returned wrong int value")
	}

	i, err = ParseQuantity("0x400")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 1024 {
		t.Errorf("Returned wrong int value")
	}

	i, err = ParseQuantity("0x0")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 0 {
		t.Errorf("Returned wrong int value")
	}
}

func TestMapToObject(t *testing.T) {
	list := []string{"Hello", "bye"}
	l2 := []string{}

	err := MapToObject(list, &l2)
	if err != nil {
		t.Errorf("%v", err)
	}

	if list[0] != l2[0] || list[1] != l2[1] {
		t.Errorf("Invalid list returned")
	}
}
*/