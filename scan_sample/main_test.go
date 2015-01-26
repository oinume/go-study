package main

import (
	_ "fmt"
	"encoding/json"
	"strings"
	"testing"
	"time"
	scan "github.com/mattn/go-scan"
)

type Tweet struct {
	IdStr           string                     `json:"idStr"`
	Text            string                     `json:"text"`
	Url             string                     `json:"url"`
	Retweeted       bool                       `json:"retweeted"`
	RetweetCount    uint64                     `json:"retweetCount"`
	FavouritesCount uint64                     `json:"favouritesCount"`
	CreatedAt       time.Time                  `json:"createdAt"`
}


func BenchmarkScanJSON(b *testing.B) {
	input := `{
		"id_str": "12345",
		"text": "あいう",
		"url": "http://hogehoge.com",
		"retweeted": true,
		"retweet_count": 10000,
		"favourites_count": 100,
		"created_at": "2014-12-25 00:00:00"
	}`

	for i := 0; i < b.N; i++ {
		tweet := make(map[string]interface{})
		err := scan.ScanJSON(strings.NewReader(input), "/", &tweet)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("hoge = %s\n", hoge["id_str"])
	}

}

func BenchmarkScanTree(b *testing.B) {
	input := `{
		"id_str": "12345",
		"text": "あいう",
		"url": "http://hogehoge.com",
		"retweeted": true,
		"retweet_count": 10000,
		"favourites_count": 100,
		"created_at": "2014-12-25 00:00:00"
	}`

	tweet := Tweet{}
	json.Unmarshal([]byte(input), &tweet)

	for i := 0; i < b.N; i++ {
		var id_str, text, url string
		var retweeted bool
		var retweet_count, favourites_count uint64
		var created_at string
		scan.ScanTree(&tweet, "/id_str", &id_str)
		scan.ScanTree(&tweet, "/text", &text)
		scan.ScanTree(&tweet, "/url", &url)
		scan.ScanTree(&tweet, "/retweeted", &retweeted)
		scan.ScanTree(&tweet, "/retweet_count", &retweet_count)
		scan.ScanTree(&tweet, "/favourites_count", &favourites_count)
		scan.ScanTree(&tweet, "/created_at", &created_at)

		//fmt.Printf("hoge = %s\n", id_str)
	}
}
