package main

//dependency import
import (
	"fmt"
	"log"
	"net/http"

	"github.com/adariki/go-service/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type User struct {
	//gorm.Model
	Username      string `json:"username" db:"username"`
	Password      string `json:"password" db:"password"`
	TipeUser      int    `json:"tipe_user" db:"tipe_user"`
	InstitutionID string `json:"institution_id" db:"institution_id"`
}

func initialMigration() {
	db, err := gorm.Open("mysql", "root@/tests?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")

	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
	db.Debug()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/login", model.Login).Methods("POST")
	myRouter.HandleFunc("/get-books", model.GetAllBook).Methods("GET")
	myRouter.HandleFunc("/get-curel/{ip}", model.GetDataCurl).Methods("GET")
	myRouter.HandleFunc("/entry-books", model.CreateBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Service Run")
	initialMigration()
	handleRequests()
}
