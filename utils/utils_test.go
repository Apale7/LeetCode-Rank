package utils

import (
	"context"
	"os"
	"testing"

	config "github.com/Apale7/LeetCode-Rank/config_loader"
	"github.com/Apale7/LeetCode-Rank/db"
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
