package config

import (
	"book/utils"
	"encoding/json"
	"io/ioutil"
)

const Dir = "config/"
const MysqlConf = Dir + "mysql.json"

type Mysql struct {
	Server   string
	Port     string
	Database string
	Username string
	Password string
}

func CreateMysql() *Mysql {
	var mysql Mysql
	isExist, _ := utils.IsFileExist(MysqlConf)
	if !isExist {
		return &mysql
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return &mysql
	}

	err = json.Unmarshal(info, &mysql)
	return &mysql
}
