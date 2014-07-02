package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Query struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel ChannelType
}

type Forecast struct {
	Date string `xml:"date,attr"`
	High int    `xml:"high,attr"`
	Low  int    `xml:"low,attr"`
	Text string `xml:"text,attr"`
}

type ItemType struct {
	XMLName   xml.Name   `xml:"item"`
	Title     string     `xml:"title"`
	Forecasts []Forecast `xml:"http://xml.weather.yahoo.com/ns/rss/1.0 forecast"`
}

type ChannelType struct {
	XMLName xml.Name `xml:"channel"`
	Title   string   `xml:"title"`
	Item    ItemType
}

func (s Forecast) String() string {
	return fmt.Sprintf("%s, \"%s\", Min: %dC Max: %dC", s.Date, s.Text, s.Low, s.High)
}

func main() {
	weatherUrl := "http://xml.weather.yahoo.com/forecastrss?w=834463&u=c"
	resp, err := http.Get(weatherUrl)
	if err != nil {
		fmt.Println("Can't perform HTTP GET:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Can't perform HTTP GET:", err)
	}

	var q Query
	err = xml.Unmarshal(body, &q)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	for _, forecast := range q.Channel.Item.Forecasts {
		fmt.Println(forecast)
	}
}
