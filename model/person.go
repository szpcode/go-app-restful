package model

import (
    "database/sql"
    "errors"
    "golang.org/x/net/context"
    "net/http"
    "fmt"
    "github.com/szpcode/tdb"
)

type Person struct {
    Id        int      `json:"id"`
    Name      *string  `json:"name"`
    Surname   *string  `json:"surname"`
    Birthday  *string  `json:"birthday"`
}

func PersonList(ctx context.Context) ([]*Person, error) {

    req := ctx.Value("req").(*http.Request)
    param := req.URL.Query()
    
    name := param.Get("name")
    surname := param.Get("surname")
    
    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return nil, errors.New("Model person: could not get database connection pool from context")
    }

    tdb.BindStrValue(":surname:", "%"+ surname +"%")

    var sql string
    
    if len(name) > 0 {
        sql += " AND name LIKE :name:" 
        tdb.BindStrValue(":name:", "%"+ name +"%")
    }
    
    if len(surname) > 0 {
        sql += " AND surname LIKE :surname:" 
        tdb.BindStrValue(":surname:", "%"+ surname +"%")
    }
    
    sql = "SELECT id, name, surname, birthday FROM person WHERE 1 = 1"+ sql
    sql = tdb.Prepare(sql)

    
    rows, err := db.Query(sql)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    peoples := make([]*Person, 0)
    for rows.Next() {
        person := new(Person)
        err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Birthday)
        if err != nil {
            fmt.Println(err)
            return nil, err
        }
        peoples = append(peoples, person)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }
    
    return peoples, nil
}


func PersonCreate(ctx context.Context) () {
    db, _ := ctx.Value("db").(*sql.DB)

    stmt, err := db.Prepare(`
        CREATE TABLE IF NOT EXISTS person 
        (id INT NOT NULL AUTO_INCREMENT, 
        name VARCHAR(32) NULL, 
        surname VARCHAR(32) NULL, 
        birthday DATE NULL, 
        PRIMARY KEY (id)) ENGINE = InnoDB;
    `)

    if err != nil {
        fmt.Println(err.Error())
    }

    _, err = stmt.Exec()

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Person table successfully migrated....")
    } 
}

func PersonAdd(ctx context.Context) (string, error) {

    req := ctx.Value("req").(*http.Request)
    param := req.URL.Query()
    
    name := param.Get("name")
    surname := param.Get("surname")
    birthday := param.Get("birthday")
  
    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return "false", errors.New("Model person: could not get database connection pool from context")
    }

    sql := "INSERT INTO `person`( name, surname, birthday ) VALUES (?, ?, ?)";

    stmt, err := db.Prepare(sql)

    if err != nil {
        return "", err
    }
    
    _, err = stmt.Exec(name, surname, birthday)
    
    if err != nil {
        return "", err
    }
    
    defer stmt.Close()

    return "true", nil
}

func PersonDelete(ctx context.Context) (string, error) {

    req := ctx.Value("req").(*http.Request)
    param := req.URL.Query()
    id := param.Get("id")
  
    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return "false", errors.New("Model person: could not get database connection pool from context")
    }

    sql := "DELETE FROM `person` WHERE id = ?";

    stmt, err := db.Prepare(sql)

    if err != nil {
        return "", err
    }
    
    _, err = stmt.Exec(id)
    
    if err != nil {
        return "", err
    }
    
    defer stmt.Close()

    return "true", nil
}

func PersonGet(ctx context.Context) (*Person, error) {

    req := ctx.Value("req").(*http.Request)
    param := req.URL.Query()
    id := param.Get("id")

    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return nil, errors.New("Model person: could not get database connection pool from context")
    }

    sql := "SELECT id, name, surname, birthday FROM person WHERE id = ?";
    row := db.QueryRow(sql, id)
    
    person := new(Person)
    
    err := row.Scan( &person.Id, &person.Name, &person.Surname, &person.Birthday )
    if err != nil {
        return nil, err
    }  
    
    return person, nil
}

func PersonEdit(ctx context.Context) (string, error) {
    req := ctx.Value("req").(*http.Request)
    param := req.URL.Query()
    
    id := param.Get("id")
    name := param.Get("name")
    surname := param.Get("surname")
    birthday := param.Get("birthday")
  
    db, ok := ctx.Value("db").(*sql.DB)
    if !ok {
        return "", errors.New("Model person: could not get database connection pool from context")
    }

    sql := "UPDATE  `person` SET name = ?, surname = ?, birthday = ?  WHERE id = ?";

    stmt, err := db.Prepare(sql)

    if err != nil {
        return "", err
    }
    
    _, err = stmt.Exec(name, surname, birthday, id)
    
    if err != nil {
        return "", err
    }
    
    defer stmt.Close()

    return "true", nil
}
