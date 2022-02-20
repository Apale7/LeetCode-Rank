package dal

import (
	config "LeetCode-Rank/config_loader"
	"LeetCode-Rank/db"
	"LeetCode-Rank/db/model"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Chdir("../../")
	config.Init()
	db.Init(context.Background())
	os.Exit(m.Run())
}

func TestGetUser(t *testing.T) {
	gotUser, err := GetUser(context.TODO())
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	b, _ := json.Marshal(gotUser)
	fmt.Printf("gotUser: %v\n", string(b))
}

func TestGetUsers(t *testing.T) {
	gotUser, err := GetUsers(context.TODO())
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	b, _ := json.Marshal(gotUser)
	fmt.Printf("gotUser: %v\n", string(b))
}

func TestCreateUser(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create user Apale",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "apale",
					Nickname: "Apale",
				},
			},
		},
		{
			name: "create user 洲哥",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "guan-shui-6",
					Nickname: "洲哥",
				},
			},
		},
		{
			name: "create user 豪哥",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "haoge_365",
					Nickname: "豪哥",
				},
			},
		},
		{
			name: "create user 英明哥",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "ming-rd",
					Nickname: "英明哥",
				},
			},
		},
		{
			name: "create user 大帅鸽",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "p2snt9d30z",
					Nickname: "大帅鸽",
				},
			},
		},
		{
			name: "create user 霹霹",
			args: args{
				ctx: context.TODO(),
				user: &model.User{
					Meta:     *model.NewMeta(),
					Username: "33qwg5mua3",
					Nickname: "霹霹",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserLatest(t *testing.T) {
	gotUser, err := GetUserLatest(context.TODO())
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	b, _ := json.Marshal(gotUser)
	fmt.Printf("gotUser: %v\n", string(b))
}
