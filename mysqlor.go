package hotelaah

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysqlor struct {
	dsn string
	db  *sql.DB
}

// table province only has one field
type ProvinceTab struct {
	provinceName string
}

// contains the fields in city table
type CityTab struct {
	cityName string
	province string
}

func NewMysqlor(d string) *Mysqlor {
	return &Mysqlor{
		dsn: d,
	}
}

func (m *Mysqlor) Open() {
	var err error
	m.db, err = sql.Open("mysql", m.dsn)
	if err != nil {
		log.Fatalf("[mysql] Open database error: %s", err.Error())
	}
	err = m.db.Ping()
	if err != nil {
		log.Fatalf("[mysql] Database Ping failure: %v", err.Error())
	}
}

// func (m *Mysqlor) CreateTable(interface{}) error {
//
// }
//
// func (m *Mysqlor) SetProvince(p ProvinceTab) error {
//
// }
//
// func (m *Mysqlor) SetCity(c CityTab) error {
//
// }

func (m *Mysqlor) QueryTest() error {
	testOut, err := m.db.Query("show databases")
	if err != nil {
		log.Printf("[mysql] QueryTest failed: %v", err.Error())
		return err
	}
	defer testOut.Close()
	for testOut.Next() {
		var databaseName string
		if err = testOut.Scan(&databaseName); err != nil {
			log.Printf("[mysql] QueryTest failed: %v", err.Error())
			return err
		} else {
			fmt.Printf("mysql has database: %s", databaseName)
		}
	}
	return nil
}
