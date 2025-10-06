package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

// Structs to map the API response
type WeatherResponse struct {
    Name string `json:"name"`
    Main struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func main() {
    apiKey := os.Getenv("OWM_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set the OWM_API_KEY environment variable.")
        return
    }

    city := "Nairobi" // You can change this to any city you like
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to call API:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != 200 {
        fmt.Println("API call failed with status:", resp.Status)
        return
    }

    var weather WeatherResponse
    err = json.NewDecoder(resp.Body).Decode(&weather)
    if err != nil {
        fmt.Println("Failed to parse response:", err)
        return
    }

    fmt.Printf("Weather in %s: %.1f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
}
