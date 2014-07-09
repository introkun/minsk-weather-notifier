package weather_providers

import (
	"encoding/xml"
	"fmt"
)

type YahooQuery struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel YahooChannelType
}

type YahooForecast struct {
	Date string `xml:"date,attr"`
	High int    `xml:"high,attr"`
	Low  int    `xml:"low,attr"`
	Text string `xml:"text,attr"`
}

type YahooItemType struct {
	XMLName   xml.Name        `xml:"item"`
	Title     string          `xml:"title"`
	Forecasts []YahooForecast `xml:"http://xml.weather.yahoo.com/ns/rss/1.0 forecast"`
}

type YahooChannelType struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	Item    YahooItemType
}

func (s YahooForecast) String() string {
	return fmt.Sprintf("%s, \"%s\", Min: %dC Max: %dC", s.Date, s.Text, s.Low, s.High)
}
