package main

import (
  "fmt"
  forecast "github.com/mlbright/forecast/v2"
  "os"
)

func main() {
  lat := "37.7833"
  long := "-122.4167"

  currentConditions := fetchConditionsFor(lat, long)

  results := stringifyResults(currentConditions)
  fmt.Printf(results)
}

func fetchConditionsFor(lat, long string) (f *forecast.Forecast) {
  key := os.Getenv("FORECAST_IO_API_KEY")

  f, err := forecast.Get(key, lat, long, "now", forecast.US)
  if err != nil {
    fmt.Printf("Error occurred in API call: %s", err)
  }

  return
}

func stringifyResults(f *forecast.Forecast) (prettyForecast string) {
  prettyForecast +=
    fmt.Sprintf("Summary: %s", f.Currently.Summary) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Temperature: %.2f", f.Currently.Temperature) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Cloud Cover: %.2f", f.Currently.CloudCover) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Humidity: %.2f", f.Currently.Humidity) +
      fmt.Sprintf("\n") +
      fmt.Sprintf("Visibility: %.2f", f.Currently.Visibility)

  return
}
