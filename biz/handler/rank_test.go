package handler

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/Apale7/LeetCode-Rank/biz/dal"
	config "github.com/Apale7/LeetCode-Rank/config_loader"
	"github.com/Apale7/LeetCode-Rank/db"
	"github.com/Apale7/LeetCode-Rank/model"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

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

func BenchmarkDB(b *testing.B) {
	for i:=0; i<b.N; i++ {
		users, err := dal.GetUsers(context.TODO())
		if err != nil {
			log.Error(err)

			return
		}
		data := make([]model.Rank, 0, len(users))
		for _, user := range users {
			tmp := model.Rank{
				Name: user.Nickname,
			}

			ac24H := acceptedNDay(context.TODO(), user.ID, time.Now(), 1)
			tmp.Easy = ac24H.Easy
			tmp.Medium = ac24H.Medium
			tmp.Hard = ac24H.Hard
			ac7Day := acceptedNDay(context.TODO(), user.ID, time.Now(), 7)
			tmp.TotalAC7Day = ac7Day.Easy + ac7Day.Medium + ac7Day.Hard
			actotal, err := dal.GetAcceptedLatest(context.TODO(), dal.UserID(user.ID))
			if err != nil {
				logrus.Error(err)
				continue
			}
			tmp.TotalAC = actotal.Easy + actotal.Medium + actotal.Hard
			data = append(data, tmp)
		}
	}
}

func Benchmark_acceptedNDay(b *testing.B) {
	id, _ := primitive.ObjectIDFromHex("621232f9ef4c3ed54aef25eb")
	for i := 0; i < b.N; i++ {
		acceptedNDay(context.TODO(), id, time.Now(), 1)
	}
}

func Benchmark_acceptedTotal(b *testing.B) {
	id, _ := primitive.ObjectIDFromHex("621232f9ef4c3ed54aef25eb")
	for i := 0; i < b.N; i++ {
		dal.GetAcceptedLatest(context.TODO(), dal.UserID(id))
	}
}