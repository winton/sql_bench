package main

import (
    _ "github.com/Go-SQL-Driver/MySQL"
    "crypto/rand"
    "database/sql"
    "encoding/hex"
)

func conn() *sql.DB {
    db, err := sql.Open("mysql", "root:@/br1")
    checkErr(err)
    return db
}

func QueryAllSingleTag() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16")
}

func QueryAllSingleTagDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id = 16")
}

func QueryAllSingleTagGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 GROUP BY device_id")
}

func QueryAllSingleTagSegmented() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '4d'")
}

func QueryAllSingleTagSegmentedDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '4d'")
}

func QueryAllSingleTagSegmentedGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '4d' GROUP BY device_id")
}

func QueryAllTwoTags() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19)")
}

func QueryAllTwoTagsDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19)")
}

func QueryAllTwoTagsGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) GROUP BY device_id")
}

func QueryAllTwoTagsSegmented() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '4d'")
}

func QueryAllTwoTagsSegmentedDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '4d'")
}

func QueryAllTwoTagsSegmentedGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '4d' GROUP BY device_id")
}

func QueryAllThreeTags() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24)")
}

func QueryAllThreeTagsDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24)")
}

func QueryAllThreeTagsGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24) GROUP BY device_id")
}

func QueryAllThreeTagsSegmented() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24) AND segment = '4d'")
}

func QueryAllThreeTagsSegmentedDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24) AND segment = '4d'")
}

func QueryAllThreeTagsSegmentedGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19,24) AND segment = '4d' GROUP BY device_id")
}

func QueryAllTenTagsSegmented() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(10,16,18,19,20,22,24,84,160,161) AND segment = '4d'")
}

func QueryAllTenTagsSegmentedDistinct() {
    Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(10,16,18,19,20,22,24,84,160,161) AND segment = '4d'")
}

func QueryAllTenTagsSegmentedGrouped() {
    Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(10,16,18,19,20,22,24,84,160,161) AND segment = '4d' GROUP BY device_id")
}

func QueryAllTenTagsSegmentedGroupedParallelX(x int) {
    var chans []chan int
    
    for i := 0; i < x; i++ {
       chans = append(chans, make(chan int))
    }

    for j := 0; j < x; j++ {
        go QueryWithChannel(
            "SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(10,16,18,19,20,22,24,84,160,161) AND segment = '" + randomHex() + "' GROUP BY device_id",
            chans[j],
        )
    }

    for j := 0; j < x; j++ {
        <-chans[j]
    }
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

    // fmt.Println("row count: ", count)
}

func QueryWithChannel(query string, channel chan int) {
    Query(query)
    channel <- 1

}

func main() {
    QueryAllSingleTag()
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func randomHex() string {
    u      := make([]byte, 1)
    _, err := rand.Read(u)

    checkErr(err)

    return hex.EncodeToString(u)
}