package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var base string = "%s:%s@tcp(%s:%d)/%s?charset=utf8"
var Db *gorm.DB
var err error

func init() {
	viper.SetConfigName("dbconf")
	viper.AddConfigPath("config")
	viper.ReadInConfig()
	var dbconf conf
	viper.Unmarshal(&dbconf)
	mysql_conf := dbconf.Mysql
	Db, err = gorm.Open("mysql", fmt.Sprintf(base, mysql_conf.Username, mysql_conf.Password, mysql_conf.Host, mysql_conf.Port, mysql_conf.Dbname))
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	} else {
		fmt.Println("database linked")
	}
	if !Db.HasTable(&Accepted{}) {
		Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Accepted{})
	}
	if !Db.HasTable(&Problem{}) {
		Db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Problem{})
	}
}

func AddProblem(id string, name string) {
	p := Problem{Id: id, Name: name}
	Db.Create(&p)
	//if Db.Error != nil {
	//	log.Error(errors.WithStack(Db.Error))
	//}
}

func AddAccepted(acceptedTime int64, username string, problemId string) {
	t := time.Unix(acceptedTime, 0)
	a := Accepted{Time: t, Username: username, ProblemId: problemId}
	Db.Create(&a)
	//if Db.Error != nil {
	//	log.Error(errors.WithStack(Db.Error))
	//}
}

type conf struct {
	Mysql Mysql `json:"mysql"`
}
type Mysql struct {
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
}

type Accepted struct {
	Time      time.Time
	Username  string    `gorm:"size:64;primary_key"`
	ProblemId string    `gorm:"size:128;primary_key"`
}

type Problem struct {
	Id   string `gorm:"primary_key;size:128"`
	Name string `gorm:"size:255"`
}
