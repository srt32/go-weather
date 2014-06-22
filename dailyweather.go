package main

import (
	"fmt"
	forecast "github.com/mlbright/forecast/v2"
  "net/http"
  "log"
  //_ "github.com/lib/pq"
  "github.com/lib/pq"
  "database/sql"
	"os"
  "time"
)

// createdb weather_development
// create table conditions(cloud_cover float, humidity float, summary string, temperature float, visibility float, location_id int, created_at timestamp);

func main() {
  http.HandleFunc("/", retrieveLatestConditions)
  http.HandleFunc("/update", updateData)

  err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
  if err != nil {
    log.Fatal("Error: %v:", err)
  }
}

func retrieveLatestConditions(res http.ResponseWriter, req *http.Request) {
  latestConditions := fetchFromDatabase()

  data := []byte(fmt.Sprintln(latestConditions))
  res.Write(data)
}

func updateData(res http.ResponseWriter, req *http.Request) {
	lat := "37.7833"
	long := "-122.4167"

	current := fetchConditionsFor(lat, long)
	fmt.Println(current)

  error := insertCondition(current)

  res.Write([]byte(fmt.Sprintln(error)))
}

func fetchFromDatabase() conditions {
  db := openDatabase()
  latestRecordedConditions := db.QueryRow(`SELECT * FROM conditions WHERE location_id IS NOT NULL ORDER BY created_at ASC LIMIT 1;`)

  var cloudCover float64
  var humidity float64
  var summary string
  var temperature float64
  var visibility float64
  var location_id int
  var created_at time.Time

  err := latestRecordedConditions.Scan(&cloudCover, &humidity, &summary, &temperature, &visibility, &location_id, &created_at)
  if err != nil {
    log.Fatal(err)
  }

  return conditions{summary, temperature, cloudCover, humidity, visibility}
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

func insertCondition(current conditions) (status bool) {
  status = true
  db := openDatabase()
  _, err := db.Query(`INSERT INTO conditions(cloud_cover, humidity, summary, temperature, visibility, location_id, created_at) VALUES($1, $2, $3, $4, $5, $6, $7)`, current.CloudCover, current.Humidity, current.Summary, current.Temperature, current.Visibility, 1, time.Now())

  if err != nil {
    status = false
  }
  return status
}

func openDatabase() (db *sql.DB) {

  var connection string
  url := os.Getenv("DATABASE_URL")
  if url == "" {
    connection = "host=localhost dbname=weather_development sslmode=disable"
  } else {
    connection, _ = pq.ParseURL(url)
    connection += " sslmode=require"
  }

  db, err := sql.Open("postgres", connection)
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
