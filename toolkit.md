# OpenWeatherMap API Beginner’s Toolkit in Go

## 1. Title & Objective

**Title:**  
Getting Started with OpenWeatherMap API – A Beginner’s Toolkit in Go

**Objective:**  
- Learn how to fetch and display current weather data using the OpenWeatherMap API and Go.
- Provide a step-by-step guide so beginners can replicate the workflow and understand each step.
- Share the AI prompts used, common errors, and troubleshooting, making this a practical resource for others.

---

## 2. Quick Summary of the Technology

- **What is it?**  
  OpenWeatherMap is a popular online service that provides current weather data, forecasts, and historical weather information via easy-to-use APIs.

- **Where is it used?**  
  It’s used by developers to build weather apps, dashboards, smart home integrations, and more. You can fetch live weather data for any city, country, or even by geo-coordinates.

- **One real-world example:**  
  Many mobile weather apps and smart home systems (like digital displays or voice assistants) use OpenWeatherMap to display local weather updates for users.

---

## 3. System Requirements

- **Operating System:** Linux (also works on Windows/Mac)
- **Editor:** VS Code, Vim, Nano, or any text editor
- **Software:**  
  - Go (Golang) [Download and install from here](https://go.dev/dl/) (version 1.17+ recommended)
- **Internet connection:** Required to install Go and access the API
- **OpenWeatherMap account:** Sign up at [openweathermap.org](https://home.openweathermap.org/users/sign_up) for a free API key
- **No extra Go packages required:** Uses Go’s standard library (`net/http`, `encoding/json`)

---

## 4. Installation & Setup Instructions

### 1. Install Go

On Ubuntu/Debian:
```sh
sudo apt update
sudo apt install golang-go
```
Check installation:
```sh
go version
```

### 2. Create Your Project Folder

```sh
mkdir openweathermap-demo
cd openweathermap-demo
```

### 3. Initialize a Go Module

```sh
go mod init openweathermap-demo
```

### 4. Get Your OpenWeatherMap API Key

- Register at [OpenWeatherMap Sign Up](https://home.openweathermap.org/users/sign_up)
- Log in, go to “API keys,” and copy your default key

### 5. Set Your API Key as an Environment Variable

```sh
export OWM_API_KEY=your_actual_api_key
```
*(Replace `your_actual_api_key` with your actual API key)*

To make this automatic, add the line above to your `~/.bashrc` or `~/.zshrc`.

---

## 5. Minimal Working Example

### What does this example do?
This Go program fetches and displays the current weather for a specified city (default: Nairobi) using the OpenWeatherMap API. It demonstrates:
- Making HTTP requests in Go
- Parsing JSON responses from an API
- Using environment variables securely for API keys

#### Code (`main.go`)

```go
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
    // Get the API key from the environment variable
    apiKey := os.Getenv("OWM_API_KEY")
    if apiKey == "" {
        fmt.Println("Please set the OWM_API_KEY environment variable.")
        return
    }

    // Set the city name. Change "Nairobi" to any city you like!
    city := "Nairobi"

    // Build the API URL
    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", city, apiKey)

    // Make the HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Failed to call API:", err)
        return
    }
    defer resp.Body.Close()

    // Check for successful response
    if resp.StatusCode != 200 {
        fmt.Println("API call failed with status:", resp.Status)
        return
    }

    // Decode the JSON response
    var weather WeatherResponse
    err = json.NewDecoder(resp.Body).Decode(&weather)
    if err != nil {
        fmt.Println("Failed to parse response:", err)
        return
    }

    // Print the weather information
    fmt.Printf("Weather in %s: %.1f°C, %s\n", weather.Name, weather.Main.Temp, weather.Weather[0].Description)
}
```

#### How to Run

1. **Install Go** and verify with `go version`.
2. **Set your API key** as an environment variable:
    ```sh
    export OWM_API_KEY=your_actual_api_key
    ```
3. **Run the program**:
    ```sh
    go run main.go
    ```
4. **Expected output:**
    ```
    Weather in Nairobi: 24.5°C, few clouds
    ```

##### To change the city:
Edit the line `city := "Nairobi"` in the code and replace `"Nairobi"` with your preferred city name.

---

## 6. AI Prompt Journal

For each step, I used the following AI prompts.  
I refined and repeated prompts based on the output and any errors encountered.

| Step | Prompt Used | AI’s Helpfulness/Notes |
|------|-------------|-----------------------|
| Setup Go | "How do I install Go on Linux?" | Clear, step-by-step commands provided. |
| Project Init | "How do I initialize a Go module?" | Helpful, explained `go mod init`. |
| Fetching from API | "How do I fetch data from an API in Go?" | Provided sample code using `net/http`. |
| Parsing JSON | "What is the simplest way to parse JSON in Go?" | Useful, showed how to use `encoding/json`. |
| Env variable usage | "How do I use environment variables in Go?" | Showed how to use `os.Getenv`. |
| OpenWeatherMap usage | "How do I get an API key from OpenWeatherMap?" | Explained signup and API key retrieval steps. |
| Error handling | "How do I handle errors from HTTP requests in Go?" | Provided best practices and code samples. |
| City change | "How do I make my code fetch weather for a different city?" | Explained how to change the city variable. |

**Reflection:**  
Using AI greatly improved my productivity. I iterated on prompts whenever I hit errors or needed clarification — for example, refining “fetch data from API in Go” to ask specifically about error handling or parsing nested JSON. The AI also recommended best practices, such as storing API keys in environment variables.

---

## 7. Common Issues & Fixes

**Issue:**  
`Please set the OWM_API_KEY environment variable.`  
**Fix:**  
Make sure you exported the API key in your terminal session before running the code.

---

**Issue:**  
`API call failed with status: 401 Unauthorized`  
**Fix:**  
Double-check your API key; make sure there are no spaces and that it is active.

---

**Issue:**  
`Failed to parse response: ...`  
**Fix:**  
This usually means the API returned an error JSON instead of the expected format (often due to a wrong city name or invalid API key). Print the response body for debugging.

---

**Helpful links:**  
- [Stack Overflow: OpenWeatherMap API 401 Unauthorized](https://stackoverflow.com/questions/32626064/openweathermap-api-always-returns-401)
- [Go: Handling environment variables](https://stackoverflow.com/questions/31873396/how-to-get-environment-variable-value)

---

## 8. References

- [OpenWeatherMap API Docs](https://openweathermap.org/api)
- [Go Documentation](https://go.dev/doc/)
- [Go net/http Package](https://pkg.go.dev/net/http)
- [Go encoding/json Package](https://pkg.go.dev/encoding/json)
- [How to use environment variables in Go](https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go)
- [Go by Example: Environment Variables](https://gobyexample.com/environment-variables)

---

## 9. Appendix / Further Ideas

- Try fetching weather by geographic coordinates (lat/lon) instead of city name.
- Enhance the program to accept user input for the city.
- Add error logging or write a more sophisticated CLI.
- Use a configuration file for more options.

---

**For any questions or to see my step-by-step AI prompt journey, see this toolkit or contact me on GitHub!**
