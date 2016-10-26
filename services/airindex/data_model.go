package airindex


import (
	"encoding/xml"
)

// Item is
type Item struct {
	XMLName xml.Name `xml:"item"`
	// Title           string   `xml:"title"`
	// Link            string   `xml:"link"`
	Description string `xml:"description"`
	// Param           string   `xml:"Param"`
	// Conc            string   `xml:"Conc"`
	// AQI             string   `xml:"AQI"`
	// Desc            string   `xml:"Desc"`
	// ReadingDateTime string   `xml:"ReadingDateTime"`
}

// Channel is
type Channel struct {
	XMLName xml.Name `xml:"channel"`
	// Title       string   `xml:"title"`
	// Link        string   `xml:"link"`
	// Description string   `xml:"description"`
	// Language    string   `xml:"language"`
	// TTL         int      `xml:"ttl"`
	Item []Item `xml:"item"`
}

// Rss is
type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel
}
