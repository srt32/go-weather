package main

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  forecast "github.com/mlbright/forecast/v2"
  "log"
  "os"
)

func main() {
  c := fetchAllLocations()

  fmt.Println(c)

  db, err := sql.Open("postgres", "dbname=weather_development")
  if err != nil {
    log.Fatal(err)
  }

  _, err = db.Query(`INSERT INTO conditions(summary, temperature, cloud_cover, humidity, visibility) VALUES($1, $2, $3, $4, $5)`, c.Summary, c.Temperature, c.CloudCover, c.Humidity, c.Visibility)
  if err != nil {
    log.Fatal(err)
  }
}

func fetchAllLocations() (c conditions) {
  lat := "37.7833"
  long := "-122.4167"

  current := fetchConditionsFor(lat, long).Currently
  c = conditions{current.Summary, current.Temperature, current.CloudCover, current.Humidity, current.Visibility}
  return
}

func fetchConditionsFor(lat, long string) (f *forecast.Forecast) {
  key := os.Getenv("FORECAST_IO_API_KEY")

  f, err := forecast.Get(key, lat, long, "now", forecast.US)
  if err != nil {
    fmt.Printf("Error occurred in API call: %s", err)
  }

  return
}

type conditions struct {
  Summary     string
  Temperature float64
  CloudCover  float64
  Humidity    float64
  Visibility  float64
}

func (c conditions) String() string {
  prettyForecast :=
    fmt.Sprintf("Summary: %s", c.Summary) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Temperature: %.2f", c.Temperature) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Cloud Cover: %.2f", c.CloudCover) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Humidity: %.2f", c.Humidity) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Visibility: %.2f", c.Visibility)
  return prettyForecast
}
