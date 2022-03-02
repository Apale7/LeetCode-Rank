package handler

import (
	"context"
	"os"
	"testing"
	"time"

	config "github.com/Apale7/LeetCode-Rank/config_loader"
	"github.com/Apale7/LeetCode-Rank/db"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")
	config.Init()
	db.Init(context.TODO())
	os.Exit(m.Run())
}

func Test_acceptedNDay(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("621232f9ef4c3ed54aef25eb")
	acceptedNDay(context.TODO(), id, time.Now(), 1)
}
