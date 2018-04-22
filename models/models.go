package models

import (
	"time"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"fmt"
	"path"
	"github.com/astaxie/beego/orm"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	uid             int64
	Title           string
	Content         string    `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	exist, err := PathExists(_DB_NAME)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}

	if exist {
		fmt.Printf("has dir![%v]\n", _DB_NAME)
	} else {
		fmt.Printf("no dir![%v]\n", _DB_NAME)
		// 创建文件夹
		err := os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
		} else {
			os.Create(_DB_NAME)
		}
	}

	orm.RegisterModel(new(Category), new(Topic))

	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)

	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
