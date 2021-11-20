package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

var base string = "%s:%s@tcp(%s:%d)/%s?charset=utf8"
var Db *gorm.DB
var err error

func Init() {
	viper.SetConfigName("dbconf")
	viper.AddConfigPath("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Error(errors.WithStack(err))
		panic("viper readInConfig error")
	}
	var dbconf conf
	if err = viper.Unmarshal(&dbconf); err != nil {
		log.Error(errors.WithStack(err))
		panic("viper Unmarshal error")
	}
	mysql_conf := dbconf.Mysql
	dsn := fmt.Sprintf(base, mysql_conf.Username, mysql_conf.Password, mysql_conf.Host, mysql_conf.Port, mysql_conf.Dbname)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix:   "gormv2_",
			SingularTable: true,
		}})
	if err != nil {
		log.Error(errors.WithStack(err))
		return
	} else {
		fmt.Println("database linked")
	}
	Db.AutoMigrate(&Accepted{}, Problem{})
	InitRedis()
}

func AddProblem(id string, name string, level int) {
	p := Problem{Id: id, Name: name, Difficulty: level}
	Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&p)
	//if Db.Error != nil {
	//	log.Error(errors.WithStack(Db.Error))
	//}
}

func AddAccepted(acceptedTime int64, username string, problemId string) {
	t := time.Unix(acceptedTime, 0)
	a := Accepted{Time: t, Username: username, ProblemId: problemId}
	Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&a)
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
	Username  string `gorm:"size:64;primary_key"`
	ProblemId string `gorm:"size:128;primary_key"`
}

type Problem struct {
	Id         string `gorm:"primary_key;size:128"`
	Name       string `gorm:"size:255"`
	Difficulty int
}
