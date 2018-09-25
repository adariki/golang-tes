package components

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Log struct {
	DateAccess string `json:"date_access" db:"date_access"`
	Request    string `json:"request" db:"request"`
	Response   string `json:"response" db:"response"`
}

func SaveLog(res string, req string) {
	db, err := gorm.Open("mysql", "root@/tests?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	t := time.Now().Format("2006-01-02 15:04:05.000 +0700")
	logs := Log{
		DateAccess: t,
		Request:    req,
		Response:   res,
	}
	db.Create(&logs)
}
