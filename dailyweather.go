package main

import (
  "fmt"
  forecast "github.com/mlbright/forecast/v2"
  "os"
)

func main() {
  key := os.Getenv("FORECAST_IO_API_KEY")

  lat := "37.7833"
  long := "-122.4167"

  f, err := forecast.Get(key, lat, long, "now", forecast.US)
  if err != nil {
    fmt.Printf("Error occurred in API call: %s", err)
  }

  fmt.Printf("Summary: %s", f.Currently.Summary)
  fmt.Printf("\n")
  fmt.Printf("Temperature: %.2f", f.Currently.Temperature)
  fmt.Printf("\n")
  fmt.Printf("Cloud Cover: %.2f", f.Currently.CloudCover)
  fmt.Printf("\n")
  fmt.Printf("Humidity: %.2f", f.Currently.Humidity)
  fmt.Printf("\n")
  fmt.Printf("Visibility: %.2f", f.Currently.Visibility)
}
