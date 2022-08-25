package mrss

import (
	"encoding/xml"
	"strings"
)

// Guid is the guid element of the MRSS Item
type Guid struct {
	Text        string `xml:",chardata"`
	IsPermaLink string `xml:"isPermaLink,attr"`
}

// Content is the content element of the MRSS Item
type Content struct {
	URL      string `xml:"url,attr"`
	FileSize string `xml:"fileSize,attr"`
	Type     string `xml:"type,attr"`
	Medium   string `xml:"medium,attr"`
	Duration *int   `xml:"media:duration,omitempty"`
}

// mrssItem is the item element of the MRSS feed
type Item struct {
	XMLName     xml.Name `xml:"item"`
	Text        string   `xml:",chardata"`
	Title       string   `xml:"title"`
	PubDate     string   `xml:"pubDate"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	Guid        *Guid    `xml:"guid,omitempty"`
	Content     *Content `xml:"media:content,omitempty"`
}

// mimeType returns the MIME type of the file
func MimeType(mime string) string {
	s := strings.Split(mime, "/")
	if len(s) > 0 {
		return s[0]
	}
	return ""
}

// mimeSubType returns the MIME subtype of the file
func MimeSubType(mime string) string {
	s := strings.Split(mime, "/")
	if len(s) > 1 {
		return s[1]
	}
	return ""
}

// setDuration sets the duration of an image in the MRSS feed
func (c *Content) SetDuration(duration int) {
	c.Duration = &duration
}
