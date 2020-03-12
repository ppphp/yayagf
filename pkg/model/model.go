package model

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/start/ent"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func Open(driverName, dataSourceName string) (*sql.Driver, error) {
	drv, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return drv, nil
}

func GetClient(driver string, url string, client interface{}, newFunc func(...interface{})interface{}) error{
	drv, err := sql.Open(driver, url)
	if err != nil {
		return err
	}
	client = newFunc(ent.Driver(drv))
	return nil
}
