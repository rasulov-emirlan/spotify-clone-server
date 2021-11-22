package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewDBConfig() (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	type dbconfig struct {
		DBhost     string
		DBport     string
		DBuser     string
		DBpassword string
		DBname     string
	}

	conf := dbconfig{}
	conf.DBhost = os.Getenv("DBHOST")
	conf.DBport = os.Getenv("DBPORT")
	conf.DBuser = os.Getenv("DBUSER")
	conf.DBpassword = os.Getenv("DBPASSWORD")
	conf.DBname = os.Getenv("DBNAME")

	result := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBhost, conf.DBport, conf.DBuser, conf.DBpassword, conf.DBname)
	log.Println(result)
	return result, nil
}

func NewTestDBConfig() (string, error) {

	err := godotenv.Load(".test.env")
	if err != nil {
		return "", err
	}

	type dbconfig struct {
		DBhost     string
		DBport     string
		DBuser     string
		DBpassword string
		DBname     string
	}

	conf := dbconfig{}
	conf.DBhost = os.Getenv("DBHOST")
	conf.DBport = os.Getenv("DBPORT")
	conf.DBuser = os.Getenv("DBUSER")
	conf.DBpassword = os.Getenv("DBPASSWORD")
	conf.DBname = os.Getenv("TESTDBNAME")

	result := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBhost, conf.DBport, conf.DBuser, conf.DBpassword, conf.DBname)
	log.Println(result)
	return result, nil
}
