package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewPortForServer() (string, error) {
	err := godotenv.Load(".dev.env")
	if err != nil {
		return "", err
	}
	port := os.Getenv("PORT")
	return port, nil
}

func NewJWTToken() (string, error) {
	err := godotenv.Load(".dev.env")
	if err != nil {
		return "", err
	}
	key := os.Getenv("JWTKEY")
	return key, nil
}

// NewSQLDBconfig: this function reads config data from ".env" file
// and it reads it from the repository that you are executing the application
func NewSQLDBconfig() (string, error) {
	err := godotenv.Load(".dev.env")
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

// NewTESTSQLDBconfig: this function requires you to have ".test.env" file
// if you want to use it for testing than the ".test.env" file should be in the root
// of the package that you are testing
// !!!dont specify reall data in ".test.env" use some mockups
func NewTESTSQLDBconfig() (string, error) {
	err := godotenv.Load(".dev.test.env")
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

	result := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DBuser, conf.DBpassword, conf.DBhost, conf.DBport, conf.DBname)
	log.Println(result)
	return result, nil
}
