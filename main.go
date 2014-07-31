package main

import (
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "minsk_weather_notifier/communication"
    "minsk_weather_notifier/dal"
    "minsk_weather_notifier/weather_providers"
    "net/http"
    "os"
)

const (
    ConfigFile = "conf.json"
    DbName     = "ql.db"
)

var (
    Db  *dal.Database
    Ctx *dal.DatabaseContext
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

func saveForecast(q *weather_providers.YahooQuery) {
    if Db == nil {
        Db, Ctx = dal.InitDb(DbName)
    }
    forecast := dal.Forecast{}
    forecast.WeatherProvider = "Yahoo"
    item := q.Channel.Item.Forecasts[0]
    forecast.MinTemp = int32(item.Low)
    forecast.MaxTemp = int32(item.High)
    dal.InsertRecord(Db, Ctx, &forecast)
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

    //sendEmail(totalForecast)
    saveForecast(&q)
    defer dal.Flush(Db)
}
