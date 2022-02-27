package app

import (
	"fmt"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

type Database struct {
	app *App
}

func (database *Database) Init() (*mgo.Database, error) {

	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbtype := viper.GetString("database.dbtype")
	dbname := viper.GetString("database.dbname")

	fmt.Println(dbtype + "://" + user + ":" + password + "@" + host + ":" + port)
	session, err := mgo.Dial(dbtype + "://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?authSource=admin")
	if err != nil {
		fmt.Print(err)
		panic(err)
		// return nil, err
	}

	db := session.DB(dbname)
	return db, nil
}
