package utils

import (
	config "LeetCode-Rank/config_loader"
	"LeetCode-Rank/db"
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../")
	config.Init()
	db.Init(context.Background())
	m.Run()
}

func TestUpdate(t *testing.T) {
	Update(context.TODO())
}
