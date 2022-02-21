package handler

import (
	"LeetCode-Rank/biz/dal"
	db_model "LeetCode-Rank/db/model"
	"LeetCode-Rank/model"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetList(c *gin.Context) {
	users, err := dal.GetUsers(c)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	data := make([]model.Rank, 0, len(users))
	for _, user := range users {
		tmp := model.Rank{
			Name: user.Nickname,
		}

		ac24H := acceptedNDay(c, user.ID, time.Now(), 1)
		tmp.Easy = ac24H.Easy
		tmp.Medium = ac24H.Medium
		tmp.Hard = ac24H.Hard
		ac7Day := acceptedNDay(c, user.ID, time.Now(), 7)
		tmp.TotalAC7Day = ac7Day.Easy + ac7Day.Medium + ac7Day.Hard
		actotal, err := dal.GetAcceptedLatest(c, dal.UserID(user.ID))
		if err != nil {
			logrus.Error(err)
			continue
		}
		tmp.TotalAC = actotal.Easy + actotal.Medium + actotal.Hard
		data = append(data, tmp)
	}
	c.JSON(http.StatusOK, data)
}

func acceptedNDay(ctx context.Context, userID primitive.ObjectID, end time.Time, n int) *db_model.Accepted {
	begin := end.AddDate(0, 0, -n)
	beginAc, err := dal.GetAcceptedEarlist(ctx, dal.CreatedAtGTE(begin), dal.UserID(userID))
	if err != nil {
		logrus.Error(err)
		return nil
	}
	fmt.Printf("%+v\n", beginAc)
	endAc, err := dal.GetAcceptedLatest(ctx, dal.CreatedAtLT(end), dal.UserID(userID))
	if err != nil {
		logrus.Error(err)
		return nil
	}
	fmt.Printf("%+v\n", endAc)
	if beginAc == nil {
		return endAc
	}
	if endAc == nil {
		return nil
	}
	return endAc.Sub(beginAc)
}
