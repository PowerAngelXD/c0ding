package dao

import "gorm.io/gorm"

var DB *gorm.DB

var DBDevice = "localhost"
var DBName = "c0lib"
var DBUser = "fzsgball"
var DBPswd = "123456"
var DBPort = "5432"

var DSN = "host=" + DBDevice +
	" user=" + DBUser +
	" password=" + DBPswd +
	" dbname=" + DBName +
	" port=" + DBPort +
	" sslmode=disable TimeZone=Asia/Shanghai"
