package utils

import (
	"LeetCode-Rank/db"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../")
	db.Init()
	m.Run()
}

func TestUpdate(t *testing.T) {
	Update()
}
