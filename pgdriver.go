package pgdriver

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

var Used = false
var pgd = pq.Driver{}

func Use() driver.Driver {
	Used = true
	return &pgd
}

func GetDriver() driver.Driver {
	return &pgd
}
