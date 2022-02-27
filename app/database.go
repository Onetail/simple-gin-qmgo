package app

import (
	"fmt"

	"context"

	"github.com/qiniu/qmgo"
	"github.com/spf13/viper"
)

type Database struct {
	app *App
}

func (database *Database) Init() (*qmgo.QmgoClient, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbtype := viper.GetString("database.dbtype")
	dbname := viper.GetString("database.dbname")

	fmt.Println(dbtype + "://" + user + ":" + password + "@" + host + ":" + port)
	ctx := context.Background()
	cli, err := qmgo.Open(ctx, &qmgo.Config{Uri: dbtype + "://" + user + ":" + password + "@" + host + ":" + port, Database: dbname, Coll: "user"})
	if err != nil {
		return nil, err
	}

	return cli, nil
}
