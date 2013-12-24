package main

import (
    _ "github.com/Go-SQL-Driver/MySQL"
    "database/sql"
    "fmt"
)

func conn() *sql.DB {
    db, err := sql.Open("mysql", "root:@/br1")
    checkErr(err)
    return db
}

func QueryAllSingleTag() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16")
}

func QueryAllSingleTagGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 GROUP BY device_id")
}

func QueryAllTwoTags() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19)")
}

func Query(query string) {
    var count int

    db := conn()
    defer db.Close()

    rows, err := db.Query(query)
    checkErr(err)

    for rows.Next() {
        var device_id string
        err = rows.Scan(&device_id)
        checkErr(err)
        count++
    }

    fmt.Println("row count: ", count)
}

func main() {
    QueryAllSingleTag()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}