package main

import (
	"fmt"
	forecast "github.com/mlbright/forecast/v2"
  "log"
  _ "github.com/lib/pq"
  "database/sql"
	"os"
)

// createdb weather_development
// create table conditions(cloud_cover float, humidity float, summary string, temperature float, visibility float, location_id int, created_at timestamp);

func main() {
	lat := "37.7833"
	long := "-122.4167"

	current := fetchConditionsFor(lat, long)
	fmt.Println(current)

  insertCondition(current)
}

func fetchConditionsFor(lat, long string) (c conditions) {
	key := os.Getenv("FORECAST_IO_API_KEY")

	f, err := forecast.Get(key, lat, long, "now", forecast.US)
	if err != nil {
		fmt.Printf("Error occurred in API call: %s", err)
	}

  current := f.Currently

	c = conditions{current.Summary, current.Temperature, current.CloudCover, current.Humidity, current.Visibility}

	return
}

func insertCondition(current conditions) {
  db := openDatabase()
  _, err := db.Query(`INSERT INTO conditions(cloud_cover, humidity, summary, temperature, visibility) VALUES($1, $2, $3, $4, $5)`, current.CloudCover, current.Humidity, current.Summary, current.Temperature, current.Visibility)

  if err != nil {
    log.Fatal(err)
  }
}

func openDatabase() (db *sql.DB) {
  db, err := sql.Open("postgres", "host=localhost dbname=weather_development sslmode=disable")
  if err != nil {
    log.Fatal(err)
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
