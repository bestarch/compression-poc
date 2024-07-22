package loader

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func LoadData(rdb *redis.Client) {

	rdb.FlushDB(ctx)

	var prefix = "_SAM_:"
	err1 := rdb.HSet(ctx, prefix+"sh_0000000000000000000_000000000", map[string]interface{}{
		"cluster": "{\"cluster_id\": 00}",
		"profile": "{\"age\": \"Adult\", \"gender\": \"gender\", \"top_genre\": [\"genre\"], \"top_region\": [\"Region\"]}",
	}).Err()
	if err1 != nil {
		log.Fatalf("Could not write hash record: %v", err1)
		panic(err1)
	}

	err2 := rdb.HSet(ctx, prefix+"sh_cluster_00", map[string]interface{}{
		"shortinfodata": "{\"0000\": {\"stackName\": \"cluster_00_stack_1\", \"description\": \"This is cluster description\", \"firstContent\": 11111111111, \"stackScore1\": 11111, \"stackScore2\": 22222, \"stackScore3\": 33333, \"hashvalue\": \"132323234hjhfj83923832nksf900sfs\", \"hashtimestamp\": \"12222222222222\"}}",
		"reco":          "{\"0000\": [1111111111, 22222222222, 3333333333333, 444444444444, 555555555555]}",
	}).Err()
	if err2 != nil {
		log.Fatalf("Could not write hash record: %v", err2)
		panic(err2)
	}

	err3 := rdb.HSet(ctx, prefix+"0000000000000000000_000000000", map[string]interface{}{
		"cos":     "{\"livo\": \"111111111111,222222222222\", \"livs\": \"111111111111,222222222222\"}",
		"profile": "{\"top_genre\": [\"genre\"], \"top_language\": [\"language\"], \"top_timecategory\": null, \"top_region\": [\"country\"], \"watched\": [{\"user_category\": \"original\", \"content\": [\"11111111111_6\"]}]}",
		"clb":     "{\"livo\": \"111111111111,222222222222\", \"livs\": \"111111111111,222222222222\"}",
		"sri":     "{\"livo\": \"111111111111,222222222222\", \"livs\": \"111111111111,222222222222\"}",
		"cldlng":  "{\"livo\": \"111111111111,222222222222\", \"livs\": \"111111111111,222222222222\"}",
	}).Err()
	if err3 != nil {
		log.Fatalf("Could not write hash record: %v", err3)
		panic(err3)
	}

	err4 := rdb.HSet(ctx, prefix+"country_region_age_gender", map[string]interface{}{
		"livo":  "{\"language1\":[111111111111,222222222222],\"language2\":[111111111111,222222222222]}",
		"sport": "{\"language1\":[111111111111,222222222222]}",
		"s":     "{\"language1\":[111111111111,222222222222],\"language2\":[111111111111,222222222222]}",
		"livs":  "{\"language1\":[111111111111,222222222222],\"language2\":[111111111111,222222222222]}",
		"mov":   "{\"language1\":[111111111111,222222222222],\"language2\":[111111111111,222222222222]}",
	}).Err()
	if err4 != nil {
		log.Fatalf("Could not write hash record: %v", err4)
		panic(err4)
	}

	err5 := rdb.HSet(ctx, prefix+"sh_region_platform", map[string]interface{}{
		"shortinfodata": "{\"0000\": {\"stackName\": \"stack_1\", \"description\": \"This is stack description\", \"firstContent\": 11111111111, \"stackScore1\": 11111, \"stackScore2\": 22222, \"stackScore3\": 33333, \"hashvalue\": \"132323234hjhfj83923832nksf900sfs\", \"hashtimestamp\": \"12222222222222\"}}",
		"reco":          "{\"0000\": [1111111111, 22222222222, 3333333333333, 444444444444, 555555555555]}",
	}).Err()
	if err5 != nil {
		log.Fatalf("Could not write hash record: %v", err5)
		panic(err5)
	}

	err6 := rdb.HSet(ctx, prefix+"country_region_platform", map[string]interface{}{
		"livo": "{\"language1\": [11111111111], \"language2\": [2222222222], \"language3\": [3333333333]}",
		"s":    "{\"language1\": [11111111111], \"language2\": [2222222222], \"language3\": [3333333333]}",
		"livs": "{\"language1\": [11111111111], \"language2\": [2222222222], \"language3\": [3333333333]}",
		"mov":  "{\"language1\": [11111111111], \"language2\": [2222222222], \"language3\": [3333333333]}",
	}).Err()
	if err6 != nil {
		log.Fatalf("Could not write hash record: %v", err6)
		panic(err6)
	}

	err7 := rdb.Set(ctx, prefix+"DUMMY_TEXT", "This is a dummy text just to simulate a 'long running textual description' of anything. This is being used for testing purposes so as to estimate how much memory will be optimised if the same text will be compressed using some algorithm in different programming languages for example: Java, Python, Go, Javascript.This is a dummy text just to simulate a 'long running textual description' of anything. This is being used for testing purposes so as to estimate how much memory will be optimised if the same text will be compressed using some algorithm in different programming languages for example: Java, Python, Go, Javascript.This is a dummy text just to simulate a 'long running textual description' of anything. This is being used for testing purposes so as to estimate how much memory will be optimised if the same text will be compressed using some algorithm in different programming languages for example: Java, Python, Go, Javascript.This is a dummy text just to simulate a 'long running textual description' of anything. This is being used for testing purposes so as to estimate how much memory will be optimised if the same text will be compressed using some algorithm in different programming languages for example: Java, Python, Go, Javascript.This is a dummy text just to simulate a 'long running textual description' of anything. This is being used for testing purposes so as to estimate how much memory will be optimised if the same text will be compressed using some algorithm in different programming languages for example: Java, Python, Go, Javascript.", 0).Err()
	if err7 != nil {
		panic(err7)
	}

	members := []redis.Z{
		{Score: 10, Member: "XYZ#{123}+789"},
		{Score: 20, Member: "XYZ#{345}+123"},
	}
	err8 := rdb.ZAdd(ctx, prefix+"ABC#123", members...).Err()
	if err8 != nil {
		log.Fatalf("Could not add members to sorted set: %v", err8)
	}

	err9 := rdb.Set(ctx, prefix+"XYZ#{123}+789", "{\"PK\":\"123\",\"SK\":\"ABC#123\",\"assetId\":\"123\",\"totalDuration\":1000000,\"position\":100000,\"created_at\":1658910745098,\"updated_at\":1658910745098,\"isWatching\":true,\"isOnAir\":true,\"language\":\"\",\"Data\":\"ABC#123\",\"episodeNumber\":1,\"showId\":123,\"seasonId\":123}", 0).Err()
	if err9 != nil {
		panic(err9)
	}

	err10 := rdb.Set(ctx, prefix+"XYZ#{345}+123", "{\"PK\":\"456\",\"SK\":\"ABC#231\",\"assetId\":\"124\",\"totalDuration\":1000000,\"position\":100000,\"created_at\":1658910745098,\"updated_at\":1658910745098,\"isWatching\":true,\"isOnAir\":true,\"language\":\"\",\"Data\":\"ABC#325\",\"episodeNumber\":2,\"showId\":123,\"seasonId\":123}", 0).Err()
	if err10 != nil {
		panic(err10)
	}
}