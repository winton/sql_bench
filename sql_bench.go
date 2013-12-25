package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
)

//

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
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '" + randomHex() + "'")
}

func QueryAllSingleTagSegmentedDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '" + randomHex() + "'")
}

func QueryAllSingleTagSegmentedGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id = 16 AND segment = '" + randomHex() + "' GROUP BY device_id")
}

//

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
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '" + randomHex() + "'")
}

func QueryAllTwoTagsSegmentedDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '" + randomHex() + "'")
}

func QueryAllTwoTagsSegmentedGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,19) AND segment = '" + randomHex() + "' GROUP BY device_id")
}

//

func QueryAllFourTags() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) GROUP BY device_id")
}

//

func QueryAllFourTagsSegmented() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment = '" + randomHex() + "'")
}

func QueryAllFourTagsSegmentedParallel10() {
	QueryParallel(10, "SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment = '" + randomHex() + "'")
}

func QueryAllFourTagsTwoSegments() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment IN('" + randomHex() + "','" + randomHex() + "')")
}

func QueryAllFourTagsSegmentedDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment = '" + randomHex() + "'")
}

func QueryAllFourTagsSegmentedGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment = '" + randomHex() + "' GROUP BY device_id")
}

func QueryAllFourTagsSegmentedGroupedParallel10() {
	QueryParallel(10, "SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment = '" + randomHex() + "' GROUP BY device_id")
}

func QueryAllFourTagsTwoSegmentsGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,157,161,173) AND segment IN('" + randomHex() + "','" + randomHex() + "') GROUP BY device_id")
}

//

func QueryAllFourTagsSharded() {
	Query("SELECT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsShardedParallel10() {
	QueryParallel(10, "SELECT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsShardedDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsShardedDistinctParallel10() {
	QueryParallel(10, "SELECT DISTINCT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173)")
}

func QueryAllFourTagsShardedGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173) GROUP BY device_id")
}

func QueryAllFourTagsShardedGroupedParallel10() {
	QueryParallel(10, "SELECT device_id FROM mobile_subscriptions_3d WHERE tag_id IN(16,157,161,173) GROUP BY device_id")
}

//

func QueryAllTenTagsSegmented() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment = '" + randomHex() + "'")
}

func QueryAllTenTagsTwoSegments() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment IN('" + randomHex() + "','" + randomHex() + "')")
}

func QueryAllTenTagsSegmentedDistinct() {
	Query("SELECT DISTINCT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment = '" + randomHex() + "'")
}

func QueryAllTenTagsSegmentedGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment = '" + randomHex() + "' GROUP BY device_id")
}

func QueryAllTenTagsSegmentedGroupedParallel10() {
	QueryParallel(10, "SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment = '" + randomHex() + "' GROUP BY device_id")
}

func QueryAllTenTagsTwoSegmentsGrouped() {
	Query("SELECT device_id FROM mobile_subscriptions WHERE tag_id IN(16,18,19,20,24,84,157,160,161,173) AND segment IN('" + randomHex() + "','" + randomHex() + "') GROUP BY device_id")
}

//

func QueryParallel(x int, query string) {
	var chans []chan int

	for i := 0; i < x; i++ {
		chans = append(chans, make(chan int))
		go QueryWithChannel(query, chans[i])
	}

	for i := 0; i < x; i++ {
		<-chans[i]
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

	// fmt.Printf("(%v)\t", count)
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

func conn() *sql.DB {
	db, err := sql.Open("mysql", "root:@/br1")
	checkErr(err)
	return db
}

func randomHex() string {
	u      := make([]byte, 1)
	_, err := rand.Read(u)

	checkErr(err)

	return hex.EncodeToString(u)
}