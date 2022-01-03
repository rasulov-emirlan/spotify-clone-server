package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const isFromFile bool = true
const filepath string = ".dev.env"

const (
	port       = "PORT"
	dblink     = "DBLINK"
	jwtKey     = "JWTKEY"
	dbHost     = "DBHOST"
	dbPort     = "DBPORT"
	dbUser     = "DBUSER"
	dbPassword = "DBPASSWORD"
	dbName     = "DBNAME"
)

func NewSQLDBlink() (string, error) {
	if isFromFile {
		err := godotenv.Load(filepath)
		if err != nil {
			return "", err
		}
	}
	link := os.Getenv(dblink)
	return link, nil

}
func NewPortForServer() (string, error) {
	if isFromFile {
		err := godotenv.Load(filepath)
		if err != nil {
			return "", err
		}
		log.Println("is from file")
	}
	p := os.Getenv(port)
	return p, nil
}

func NewJWTToken() (string, error) {
	if isFromFile {
		err := godotenv.Load(filepath)
		if err != nil {
			return "", err
		}
	}
	k := os.Getenv(jwtKey)
	return k, nil
}

// NewSQLDBconfig: this function reads config data from ".env" file
// and it reads it from the repository that you are executing the application
func NewSQLDBconfig() (string, error) {
	if isFromFile {
		err := godotenv.Load(filepath)
		if err != nil {
			return "", err
		}
	}

	type dbconfig struct {
		DBhost     string
		DBport     string
		DBuser     string
		DBpassword string
		DBname     string
	}

	conf := dbconfig{}
	conf.DBhost = os.Getenv(dbHost)
	conf.DBport = os.Getenv(dbPort)
	conf.DBuser = os.Getenv(dbUser)
	conf.DBpassword = os.Getenv(dbPassword)
	conf.DBname = os.Getenv(dbName)

	result := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.DBhost, conf.DBport, conf.DBuser, conf.DBpassword, conf.DBname)
	log.Println(result)
	return result, nil
}

// NewTESTSQLDBconfig: this function requires you to have ".test.env" file
// if you want to use it for testing than the ".test.env" file should be in the root
// of the package that you are testing
// !!!dont specify reall data in ".test.env" use some mockups
func NewTESTSQLDBconfig() (string, error) {
	if isFromFile {
		err := godotenv.Load(filepath)
		if err != nil {
			return "", err
		}
	}

	type dbconfig struct {
		DBhost     string
		DBport     string
		DBuser     string
		DBpassword string
		DBname     string
	}

	conf := dbconfig{}
	conf.DBhost = os.Getenv(dbHost)
	conf.DBport = os.Getenv(dbPort)
	conf.DBuser = os.Getenv(dbUser)
	conf.DBpassword = os.Getenv(dbPassword)
	conf.DBname = os.Getenv(dbName)

	result := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.DBuser, conf.DBpassword, conf.DBhost, conf.DBport, conf.DBname)
	log.Println(result)
	return result, nil
}
