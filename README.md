## OpenWeatherMap Go Demo

## Overview

This is a beginner-friendly Go application that fetches and displays the current weather for a specified city using the OpenWeatherMap API. It demonstrates how to make HTTP requests, handle JSON responses, and use environment variables for API keys in Go.

## Features

- Fetches current weather data for any city
- Displays temperature and weather description
- Uses only Go standard library packages

## System Requirements

- Linux (but works on Mac/Windows too)
- Go 1.17+ installed
- Internet connection
- OpenWeatherMap API key (free from [openweathermap.org](https://openweathermap.org/))

## Setup Instructions

1. **Install Go**

    ```sh
    sudo apt update
    sudo apt install golang-go
    ```

2. **Clone this repository**

    ```sh
    git clone https://github.com/NaomiKwamboka0/openweathermap-go-demo.git
    cd openweathermap-go-demo
    ```

3. **Get your OpenWeatherMap API key**

    - Sign up at [https://home.openweathermap.org/users/sign_up](https://home.openweathermap.org/users/sign_up)
    - Log in and find your API key in the API keys section

4. **Set your API key as an environment variable**

    ```sh
    export OWM_API_KEY=your_actual_api_key
    ```

5. **Run the program**

    ```sh
    go run main.go
    ```

## Example Output

```
Weather in Nairobi: 24.5°C, few clouds
```

## Changing the City

- Open `main.go` and edit the line:
    ```go
    city := "Nairobi"
    ```
  Replace `"Nairobi"` with your desired city.

## Troubleshooting

- **API key error:**  
  If you see `Please set the OWM_API_KEY environment variable.`, make sure you set the environment variable correctly.

- **API call failed:**  
  If you see a status error, check your API key or network connection.

## References

- [OpenWeatherMap API Docs](https://openweathermap.org/api)
- [Go Documentation](https://go.dev/doc/)
- [Go net/http Package](https://pkg.go.dev/net/http)
- [Go encoding/json Package](https://pkg.go.dev/encoding/json)

## License

This project is for educational purposes.
