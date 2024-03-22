package model

import (
	"arithmetic/utils"
	"path"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
)

var _DB *sqlx.DB

func init() {
	home, e := homedir.Dir()
	if e != nil {
		utils.Logger.Fatal(e.Error())
	}
	dbPath := path.Join(home, ".arithmetic-focus")
	utils.Logger.Debug("db info", zap.String("path", dbPath))
	_DB, e = sqlx.Open("sqlite3", "file:"+dbPath)
	if e != nil {
		utils.Logger.Fatal(e.Error())
	}
	if _, e := _DB.Exec(`create table if not exists record (
		expression varchar(512) not null,
		occur int not null default 0,
		correct int not null default 0
	);
	create unique index if not exists exp on record (expression);
	`); e != nil {
		utils.Logger.Fatal(e.Error())
	}

}
