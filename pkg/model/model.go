package model

import (
	"github.com/facebookincubator/ent/dialect/sql"
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
