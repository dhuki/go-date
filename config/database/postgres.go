package database

import (
	"bytes"
	"strconv"
	"time"

	"github.com/dhuki/go-date/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(conf *config.Database) (err error) {
	PostgresDb.Master, err = PostgresDb.OpenPostgresConnection(&conf.Master, &conf.DbConnectionInfo)
	if err != nil {
		return err
	}
	PostgresDb.Slave, err = PostgresDb.OpenPostgresConnection(&conf.Slave, &conf.DbConnectionInfo)
	if err != nil {
		return err
	}
	return
}

func (d *Database) OpenPostgresConnection(dbInfo *config.DBInfo, dbConnInfo *config.DbConnectionInfo) (*sqlx.DB, error) {
	var bufferStr bytes.Buffer
	bufferStr.WriteString(" host=" + dbInfo.Host)
	bufferStr.WriteString(" port=" + strconv.Itoa(dbInfo.Port))
	bufferStr.WriteString(" user=" + dbInfo.User)
	bufferStr.WriteString(" dbname=" + dbInfo.DBName)
	bufferStr.WriteString(" password=" + dbInfo.Password)
	bufferStr.WriteString(" sslmode=disable fallback_application_name=go-date")
	connectionSource := bufferStr.String()

	db, err := sqlx.Connect("postgres", connectionSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbConnInfo.SetMaxIdleCons)
	db.SetMaxOpenConns(dbConnInfo.SetMaxOpenCons)
	db.SetConnMaxIdleTime(time.Duration(dbConnInfo.SetConMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(dbConnInfo.SetConMaxLifetime) * time.Minute)
	return db, nil
}
