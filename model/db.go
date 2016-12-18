package model

import (
    "database/sql"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)

func NewDB(driver string, host string, port int, username string, password string, database string) (*sql.DB, error) {
    db, err := sql.Open(driver, username +":"+ password +"@tcp("+ host +":"+ strconv.Itoa(port) +")/"+ database)
    
    if err != nil {
        return nil, err
    }
    
    if err = db.Ping(); err != nil {
        return nil, err
    }
    
    return db, nil
}