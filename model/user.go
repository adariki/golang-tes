package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/adariki/go-service/components"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	//gorm.Model
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	TipeUser      int    `json:"tipe_user" db:"tipe_user"`
	InstitutionID string `json:"institution_id" db:"institution_id"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	type msgerr struct {
		Status   int    `json:"status"`
		Messages string `json:"messages"`
	}

	type Results struct {
		Status int  `json:"status"`
		Data   User `json:"data"`
	}
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("mysql", "root@/tests?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var users User
	err = json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		panic("salah")

	}
	hasher := md5.New()
	hasher.Write([]byte(users.Password))
	HashRes := hex.EncodeToString(hasher.Sum(nil))
	if users.Username != "" {
		db.Where("username = ? and password = ?", users.Username, HashRes).Find(&users)
		if users.TipeUser != 0 {
			users.Password = "-"
			slice := Results{
				Status: 200,
				Data:   users,
			}
			json.NewEncoder(w).Encode(slice)
		} else {

			slice := msgerr{
				Status:   101,
				Messages: components.Eksepsi(101),
			}

			json.NewEncoder(w).Encode(slice)
		}

	} else {
		json.NewEncoder(w).Encode("harus diisi")
	}

}
