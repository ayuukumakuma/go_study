package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

var Config ConfigList

func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080), // 8080はデフォルト値(SecionのKeyが存在しない場合に使用される)
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("Name = %v\n", Config.DbName)
	fmt.Printf("Driver = %v\n", Config.SQLDriver)
}
