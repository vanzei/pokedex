# 🐾 Pokedex CLI - Your Terminal Pokemon Adventure

A powerful, interactive command-line Pokedex application built with Go that lets you explore Pokemon locations, catch Pokemon, and build your collection - all from your terminal!

![Pokedex Banner](https://raw.githubusercontent.com/PokeAPI/media/master/logo/pokeapi_256.png)

## ⭐ Features

- **Interactive CLI**: Navigate the Pokemon world through an intuitive command interface
- **Location Exploration**: Browse locations and discover Pokemon in the wild
- **Pokemon Catching**: Try to catch Pokemon with a difficulty based on their base experience
- **Detailed Inspection**: View comprehensive stats for your caught Pokemon
- **Caching System**: Efficient request handling with a custom cache implementation
- **Metrics & Monitoring**: Full observability with Prometheus and OpenTelemetry

## 📋 Commands
Pokedex > help Welcome to the Pokedex! Usage:

help - Displays a help message map - Get the next page of locations mapb - Get the previous page of locations explore <location-area-name> - Explore a location area to find Pokémon catch <pokemon-name> - Attempt to catch a pokemon inspect <pokemon-name> - View details of a caught pokemon pokedex - List all the Pokemon you've caught exit - Exit the Pokedex




## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/pokedex.git
cd pokedex

# Build the application
go build -o pokedex

# Run the application
./pokedex
```

## 🚀 Usage Examples

Exploring Locations
```
Pokedex > map
canalave-city-area
eterna-city-area
pastoria-city-area
sunyshore-city-area
sinnoh-pokemon-league-area
```
Finding Pokemon in a Location

```
Pokedex > explore eterna-city-area
Exploring eterna-city-area...
Found Pokémon:
- bidoof
- kricketot
- shinx 
- abra
- budew
```
Inspecting Your Pokemon

```
Pokedex > inspect pikachu
Name: pikachu
Height: 4
Weight: 60
Stats:
  -hp: 35
  -attack: 55
  -defense: 40
  -special-attack: 50
  -special-defense: 50
  -speed: 90
Types:
  - electric
```

## 📊 Monitoring with Prometheus and Grafana
The application includes built-in monitoring capabilities through Prometheus and can be visualized with Grafana dashboards.

Prometheus Metrics
Metrics are exposed at http://localhost:8080/metrics. The main metric is my_app_requests_total, which tracks API calls with method and path labels.

```
# HELP my_app_requests_total Total number of requests to my app
# TYPE my_app_requests_total counter
my_app_requests_total{method="GET",path="catch"} 12
my_app_requests_total{method="GET",path="cacheHit"} 45
my_app_requests_total{method="GET",path="explore"} 23
my_app_requests_total{method="GET",path="mapHit"} 8
my_app_requests_total{method="GET",path="pokedex"} 5
```
1. Setting up Grafana Dashboard
Make sure Prometheus is configured to scrape the application metrics:

```bash
# prometheus.yml
scrape_configs:
  - job_name: 'pokedex'
    scrape_interval: 5s
    static_configs:
      - targets: ['localhost:8080']
```
2. Create a Grafana dashboard with the following panels:

**API Calls by Type**: Shows the distribution of API calls
```
sum by (path) (my_app_requests_total)
```

**Cache Hit Ratio**: Shows the efficiency of the cache
```
sum(my_app_requests_total{path="cacheHit"}) / sum(my_app_requests_total)
```

**API Calls by Type**: Shows requests per second
```
rate(my_app_requests_total[1m])
```

## 🏗️ Architecture
Command Structure: Clean separation of concerns with individual command files
API Client: Efficient client for the PokeAPI with proper error handling
Custom Cache: Time-based eviction cache to reduce API load
Metrics Collection: Prometheus integration for monitoring API usage
OpenTelemetry: Distributed tracing for advanced troubleshooting
## 📚 Implementation Details
The application uses:

Go's built-in concurrency for background cache eviction
Custom HTTP client with timeouts and context support
Prometheus counters for API metrics
Math/rand for Pokemon catch probability calculations
JSON marshaling/unmarshaling for data serialization
## 🧠 Behind the Catch Mechanics
The probability of catching a Pokemon is inversely proportional to its base experience:

This means:

A Pokemon with base experience 50 has ~12.5% difficulty (87.5% catch rate)
A Pokemon with base experience 300 has 75% difficulty (25% catch rate)
Legendary Pokemon (360+ base exp) are capped at 90% difficulty (10% catch rate)
## 🤝 Contributing
Contributions are welcome! Feel free to:

Fork the repo
Create a feature branch (git checkout -b amazing-feature)
Commit your changes (git commit -m 'Add amazing feature')
Push to the branch (git push origin amazing-feature)
Open a Pull Request
## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments
PokeAPI for the comprehensive Pokemon data
The Go community for the amazing language and tools
Happy Pokemon catching! 🎮