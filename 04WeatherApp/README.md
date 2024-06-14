# Weather App

WeatherApp is a simple command-line application written in Go that allows users to get current weather information. User can also use --city flag for specific City.

## Installation 
1. **Clone the repository:**
```sh
git clone https://github.com/niteshkamble/go-lang-learning
```
2. **Install dependencies:**
Make sure you have Go installed. If not, download and install it from [the official Go website](https://golang.org/dl/). Then Navigate to `04WeatherApp` directory.

```sh
go mod tidy
```

3. **Create a `.env` file:**
Create a file named `.env` in the 04WeatherApp directory of your project and add your API key. [Weather API](https://www.weatherapi.com/)

```plaintext
API_KEY=your_api_key_here
```

4. **Run the application:**

```sh
go run main.go
```
```sh
go run main.go --city tokyo
```
