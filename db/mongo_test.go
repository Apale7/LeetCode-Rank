package db

import (
	config "LeetCode-Rank/config_loader"
	"context"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../")
	config.Init()
	os.Exit(m.Run())
}
func TestInitMongo(t *testing.T) {
	InitMongo(context.TODO())
}
