package db

import (
	"context"
	"os"
	"testing"

	config "github.com/Apale7/LeetCode-Rank/config_loader"
)

func TestMain(m *testing.M) {
	os.Chdir("../")
	config.Init()
	Init(context.TODO())
	os.Exit(m.Run())
}

func TestInitMongo(t *testing.T) {
	InitMongo(context.TODO())
}
