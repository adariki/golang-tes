package model

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/adariki/go-service/components"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	BookName  string  `json:"book_name" db:"book_name"`
	BookIsbn  string  `json:"book_isbn" db:"book_isbn"`
	BookPrice float64 `json:"book_price" db:"book_price"`
}

type Log struct {
	DateAccess time.Time `json:"date_access" db:"date_access"`
	IP         string    `json:"ip" db:"ip"`
	Message    string    `json:"message" db:"message"`
}

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("mysql", "root@/tests?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	type msgerr struct {
		Status   int    `json:"status"`
		Messages string `json:"messages"`
	}

	type Results struct {
		Status int    `json:"status"`
		Data   []Book `json:"data"`
	}
	var books []Book
	db.Find(&books)
	slice := Results{
		Status: 200,
		Data:   books,
	}
	jeson, _ := json.Marshal(slice)
	jesonreq, _ := json.Marshal(r.Body)
	json.NewEncoder(w).Encode(slice)
	components.SaveLog(string(jeson), string(jesonreq))
	fmt.Println(r.Body)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db, err := gorm.Open("mysql", "root@/tests?parseTime=true")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	type msgerr struct {
		Status   int    `json:"status"`
		Messages string `json:"messages"`
	}

	type Results struct {
		Status   int    `json:"status"`
		Messages string `json:"messages"`
	}

	var books Book
	slice := Results{
		Status:   200,
		Messages: "success",
	}
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(r.Body)
	}
	// Restore the io.ReadCloser to its original state
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Use the content
	bodyString := string(bodyBytes)
	jeson, _ := json.Marshal(slice)
	//jesonreq, _ := json.Marshal(r.Body)
	//jsonreq := r.Body
	er := json.NewDecoder(r.Body).Decode(&books)
	if er != nil {
		http.Error(w, "", http.StatusBadRequest)
		slice = Results{
			Status:   401,
			Messages: "failed",
		}
		jeson, _ := json.Marshal(slice)
		json.NewEncoder(w).Encode(slice)
		components.SaveLog(string(jeson), bodyString)
	} else {
		//db.Create(&books)
		db.Create(&books)
		//http.Error(w, "", http.StatusBadRequest)

		json.NewEncoder(w).Encode(slice)
		components.SaveLog(string(jeson), bodyString)
	}

}
