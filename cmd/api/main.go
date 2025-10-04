package main

import (
	"fmt"
	"log"
	"myapp/api/router"
	"myapp/config"
	"myapp/util/validator"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

const fmtDBString = "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"

//	@title			MYAPP API
//	@version		1.0
//	@description	This is a sample RESTful API with a CRUD

//	@contact.name	Natnael Alemayehu
//	@contact.url	se.natnael.alemayehu@gmail.com

//	@license.name	MIT License

// @host		localhost:8080
// @baseurl	/v1
func main() {

	c := config.New()
	v := validator.New()

	var logLevel gormlogger.LogLevel
	if c.DB.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}

	dbString := fmt.Sprintf(fmtDBString, c.DB.Host, c.DB.Username, c.DB.Password, c.DB.DBName, c.DB.Port)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{Logger: gormlogger.Default.LogMode(logLevel)})
	if err != nil {
		log.Fatal("DB connection start failure")
		return
	}

	r := router.New(db, v)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server ", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
