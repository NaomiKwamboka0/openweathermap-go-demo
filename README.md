# OpenWeatherMap Go Demo

## Overview

A beginner-friendly Go application that fetches and displays current weather for a specified city using the OpenWeatherMap API.

## Features

- Fetches current weather for any city
- Uses Go’s standard library (no extra dependencies)
- Securely handles API keys via environment variables

## System Requirements

- Linux, macOS, or Windows
- Go 1.17 or later
- OpenWeatherMap API key (free from [openweathermap.org](https://openweathermap.org/))

## Setup Instructions

1. **Install Go**  
   On Ubuntu/Debian:
   ```sh
   sudo apt update
   sudo apt install golang-go
   ```

2. **Clone the Repository**
   ```sh
   git clone https://github.com/NaomiKwamboka0/openweathermap-go-demo.git
   cd openweathermap-go-demo
   ```

3. **Get Your API Key**
   - Sign up at [OpenWeatherMap](https://home.openweathermap.org/users/sign_up)
   - Find your key in the API keys section

4. **Set Your API Key**
   ```sh
   export OWM_API_KEY=your_actual_api_key
   ```

## How to Run

```sh
go run main.go
```

**Example output:**
```
Weather in Nairobi: 24.5°C, few clouds
```

## Changing the City

Edit the line in `main.go`:
```go
city := "Nairobi"
```
Replace `"Nairobi"` with any city you want.

## AI Prompts Used

Below are some example AI prompts I used to complete this project:

- “How do I fetch data from an API in Go?”
- “What is the simplest way to parse JSON in Go?”
- “How do I use environment variables in Go?”
- “How do I handle errors from HTTP requests in Go?”
- “How do I get an API key from OpenWeatherMap?”

For a detailed step-by-step guide, my reflections, and the full prompt journal, **see [toolkit.md](toolkit.md)**.

## Troubleshooting

- **API key error:** Ensure you've set the API key as described above.
- **API call failed:** Check your internet connection and API key validity.

## References

- [OpenWeatherMap API Docs](https://openweathermap.org/api)
- [Go Documentation](https://go.dev/doc/)

## License

For educational purposes.
