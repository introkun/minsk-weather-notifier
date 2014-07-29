package main

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "minsk_weather_notifier/communication"
    "minsk_weather_notifier/weather_providers"
    "net/http"
    "os"
)

const (
    ConfigFile = "conf.json"
)

func sendEmail(body string) {
    configFile, _ := os.Open(ConfigFile)
    configDecoder := json.NewDecoder(configFile)
    smtpInfo := communication.SmtpInfo{}
    err := configDecoder.Decode(&smtpInfo)
    if err != nil {
        fmt.Println("error:", err)
    }

    fmt.Println(smtpInfo)
    communication.Email(smtpInfo, "Weather forecast", body)
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

    var q weather_providers.YahooQuery
    err = xml.Unmarshal(body, &q)
    if err != nil {
        fmt.Println("Error Unmarshal:", err)
        return
    }

    totalForecast := ""
    for _, forecast := range q.Channel.Item.Forecasts {
        totalForecast += fmt.Sprintln(forecast)
    }
    totalForecast += "\r\n--\r\nBest,\r\nSergey"

    sendEmail(totalForecast)
}
