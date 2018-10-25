package main

import (
	"fmt"
	"gorm/model"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./local.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open database
	db, err := gorm.Open(os.Getenv("DIALECT"), os.Getenv("DSN"))
	if err != nil {
		fmt.Printf("failed to connect database : %v", err)
	}
	defer db.Close()

	tx := db.Begin()
	tx.LogMode(true)

	m := model.User{}

	nameBase := "dummy"
	pictureBase := "hogehoge"
	messageBase := "message"
	userArr := []model.User{}

	// create array
	for i := 0; i < 100; i++ {
		nStr := strconv.Itoa(i)
		m.ID = 0
		m.Name = nameBase + nStr
		m.StatusMessage = messageBase + nStr
		m.PictureUrl = pictureBase + nStr
		m.Datetime = time.Now().Local()

		userArr = append(userArr, m)
	}

	// create query
	query := toInsertFriendQuery(userArr)

	// exec bulk insert
	err = tx.Exec(query).Error
	if err != nil {
		fmt.Printf("failed to Bulk Insert : %v", err)
		tx.Rollback()
		return
	}

	tx.Commit()
}

func toInsertFriendQuery(s []model.User) string {
	var query string = "INSERT INTO `user` (`id`, `name`, `picture_url`, `status_message`, `datetime`) VALUES "
	var values = make([]string, len(s))
	for i, data := range s {
		values[i] = toInsertFriendValue(data)
	}

	return query + strings.Join(values, ",") + ";"
}

func toInsertFriendValue(s model.User) string {
	return "(" + strings.Join([]string{
		"'" + strconv.FormatInt(s.ID, 10) + "'",
		"'" + s.Name + "'",
		"'" + s.PictureUrl + "'",
		"'" + s.StatusMessage + "'",
		"'" + s.Datetime.Format("2006-01-02 15:04:05") + "'",
	}, ",") + ")"
}