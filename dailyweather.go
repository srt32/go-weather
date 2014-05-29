package main

import (
  "fmt"
  forecast "github.com/mlbright/forecast/v2"
  "log"
)

func main() {
  key := "xxx"

  lat := "35"
  long := "-122"

  f, err := forecast.Get(key, lat, long, "now", forecast.US)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%s\n", f.Currently.Summary)
  fmt.Printf("%.2f", f.Currently.Temperature)
}
