package model

import (
	"github.com/facebook/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// 有什么意义吗，有的，import少写两行。。。
func Open(driverName, dataSourceName string) (*sql.Driver, error) {
	drv, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return drv, nil
}
