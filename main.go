package main

import (
	"fmt"
	"github.com/harshil-jain-08/day3/repo"
	"github.com/harshil-jain-08/day3/service"

	"github.com/harshil-jain-08/day3/Config"
	"github.com/harshil-jain-08/day3/Models"
	"github.com/harshil-jain-08/day3/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer func() {
		err := Config.DB.Close()
		if err != nil {
			fmt.Println("Closing DB failed with error: ", err.Error())
		}
	}()

	Config.DB.AutoMigrate(&Models.Record{})

	service.NewService(repo.NewRepo())

	r := Routes.SetupRouter()
	// running
	err = r.Run()
	if err != nil {
		panic(any(err))
	}
}
