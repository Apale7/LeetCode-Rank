package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../")
	Init()
	os.Exit(m.Run())
}

func TestGetAll(t *testing.T) {
	all := GetAll()
	if all == nil {
		t.Error("GetAll() failed")
	}
	fmt.Printf("all: %+v\n", all)
}
