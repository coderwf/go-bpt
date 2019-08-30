package bound

import (
	"encoding/xml"
	"math/rand"
	"time"
)

func genNumbers(cnt int) []int{
	rand.Seed(time.Now().Unix())
	nums := make([]int, 0, cnt)
	for i:= 0;i< cnt; i++{
        nums = append(nums, rand.Intn(1000))
	}
	return nums
}


var file = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
        <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
        <title>Object Oriented Programming Mechanics</title>
        <description>Go is an amazing language.</description>
        <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
</channel>
</rss>`

type item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []item   `xml:"item"`
}

type document struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
	URI     string
}


func generateList(totalDocs int) []string {
	docs := make([]string, totalDocs)
	for i := 0; i < totalDocs; i++ {
		docs[i] = "test.xml"
	}
	return docs
}
