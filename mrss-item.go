package mrss

import (
	"encoding/xml"
	"strings"
)

// Guid is the guid element of the MRSS Item
type Guid struct {
	XMLName     xml.Name `xml:"guid"`
	Text        string   `xml:",chardata"`
	IsPermalink string   `xml:"isPermalink,attr"`
}

// Content is the content element of the MRSS Item
type Content struct {
	XMLName  xml.Name `xml:"content"`
	URL      string   `xml:"url,attr"`
	FileSize string   `xml:"fileSize,attr"`
	Type     string   `xml:"type,attr"`
	Medium   string   `xml:"medium,attr"`
}

// mrssItem is the item element of the MRSS feed
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:",chardata"`
	Title       string   `xml:"title"`
	PubDate     string   `xml:"pubDate"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Guid        Guid     `xml:"guid"`
	Content     struct {
		URL      string `xml:"url,attr"`
		FileSize string `xml:"fileSize,attr"`
		Type     string `xml:"type,attr"`
		Medium   string `xml:"medium,attr"`
	} `xml:"media:content"`
}

// mimeType returns the MIME type of the file
func (i Item) mimeType() string {
	s := strings.Split(i.Content.Type, "/")
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

// mimeSubType returns the MIME subtype of the file
func (i Item) mimeSubType() string {
	s := strings.Split(i.Content.Type, "/")
	if len(s) > 1 {
		return s[1]
	}
	return ""
}
