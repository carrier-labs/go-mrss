package mrss

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

// Channel is the channel element of the MRSS feed
type Channel struct {
	XMLName   xml.Name `xml:"channel"`
	Text      string   `xml:",chardata"`
	Title     string   `xml:"title"`
	Link      string   `xml:"link"`
	Generator string   `xml:"generator"`
	Ttl       string   `xml:"ttl"`
	Item      []Item   `xml:"item"`
}

// addItem adds an item to the MRSS feed
func (c *Channel) AddItem(item Item) {
	c.Item = append(c.Item, item)
}

// Feed is the root element of the MRSS feed
type Feed struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Media   string   `xml:"xmlns:media,attr"`
	Channel Channel  `xml:"channel"`
}

// Bytes returns the MRSS feed as XML Bytes
func (m *Feed) ToBytes() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(`<?xml version="1.0" encoding="utf-8" ?>`)
	b, err := xml.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("xml marshall")
	}
	buf.Write(b)
	return buf.Bytes(), nil
}

// Bytes returns the MRSS feed as XML String
func (m *Feed) ToString() (string, error) {
	b, err := m.ToBytes()
	return string(b), err
}

// New creates a new MRSSFeed ready to be populated
func New(title string, link string) *Feed {
	return &Feed{
		XMLName: xml.Name{},
		Text:    "",
		Version: "2.0",
		Media:   "http://search.yahoo.com/mrss/", //https://www.rssboard.org/media-rss
		Channel: Channel{
			Text:      "",
			Title:     title,
			Link:      link,
			Generator: "https://github.com/carrier-labs/go-mrss",
			Ttl:       "5",
		},
	}
}
