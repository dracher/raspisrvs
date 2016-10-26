package airindex


import "testing"
import "strings"

var data = `
<rss version="2.0">
<channel>
    <title>Latest Air Quality for Beijing</title>
    <link>http://www.stateair.net/web/post/1/1.html</link>
    <description>Latest Air Quality for Beijing</description>
    <language> en_US</language>
    <ttl>15</ttl>
    <item>
        <title>10/10/2016 2:00:00 PM</title>
        <link>http://www.stateair.net/web/post/1/1.html</link>
        <description>10-10-2016 14:00; PM2.5; 105.0; 177; Unhealthy (at 24-hour exposure at this level)</description>
        <Param>PM2.5</Param>
        <Conc>105.0</Conc>
        <AQI>177</AQI>
        <Desc>Unhealthy (at 24-hour exposure at this level)</Desc>
        <ReadingDateTime>10/10/2016 2:00:00 PM</ReadingDateTime>
    </item>
    <item>
        <title>10/10/2016 1:00:00 PM</title>
        <link>http://www.stateair.net/web/post/1/1.html</link>
        <description>10-10-2016 13:00; PM2.5; 94.0; 171; Unhealthy (at 24-hour exposure at this level)</description>
        <Param>PM2.5</Param>
        <Conc>94.0</Conc>
        <AQI>171</AQI>
        <Desc>Unhealthy (at 24-hour exposure at this level)</Desc>
        <ReadingDateTime>10/10/2016 1:00:00 PM</ReadingDateTime>
    </item>
</channel>
</rss>
`

func TestNewCacheObject(t *testing.T) {
	cache := NewAqiData()
	t.Log(len(cache.Result))
}

func TestGetCachedData(t *testing.T) {
	cache := NewAqiData()
	if len(cache.Result) != 0 {
		t.Fail()
	}
	cache.UpdateData()
	t.Log(cache.CurrentData("beijing"))
	t.Log(cache.CurrentData("shanghai"))
}

func TestFetchData(t *testing.T) {
	res, _ := fetchData(urlMapper["beijing"])
	

}
