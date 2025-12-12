Pokédex CLI

A fully functional Command-Line Interface Pokédex, written in Go.
It connects to the PokéAPI, caches responses, and lets you explore Pokémon, locations, and your captured Pokédex — all from the terminal.

This project is built step-by-step following the Boot.dev course and expanded with custom improvements (caching, REPL, commands, modular Go packages).

✨ Features
✔ Interactive REPL

Runs in the terminal and accepts commands like:

help — Show all available commands

map — Show next page of location areas

mapb — Go back one page

explore <area> — Show Pokémon found in an area

catch <pokemon> — Attempt to catch a Pokémon

inspect <pokemon> — View stats, types, etc.

pokedex — View your captured Pokémon

✔ Pokémon Fetching

Uses PokeAPI /pokemon/{name} endpoint to fetch:

ID

Height

Weight

Stats

Types

✔ Location Fetching

Uses /location-area and /location-area/{name} to explore areas and Pokémon inside those areas.

✔ Local Caching

Uses a custom in-memory cache with expiration.
This prevents repeated API calls and massively speeds up the CLI.

✔ Clean Project Structure

All API logic is in internal/pokeapi/
All CLI logic is in repl.go
Configuration stored in config.go

📁 Project Structure
pokedexcli/
│
├── main.go               # Entry point
├── config.go             # Global config + PokeAPI client
├── repl.go               # Interactive REPL + command dispatcher
│
└── internal/
    └── pokeapi/
        ├── pokeapi.go    # API client (requests + caching)
        ├── types.go      # Data models for JSON unmarshalling
        └── cache.go      # In-memory timed cache

🚀 Installation
Clone the project
git clone https://github.com/Liamwolf56/pokedexcli
cd pokedexcli

Run it
go run .

(Optional) Build binary
go build -o pokedex .
./pokedex

🕹 REPL Commands

Below is every command your CLI supports.

help

Displays a list of all supported commands.

map

Shows the next page of location areas from the API.

map

mapb

Moves backward one page of locations.

explore <area>

Shows Pokémon found in a specific location-area.

Example:

explore kanto-route-1

catch <pokemon>

Attempts to catch a Pokémon.
Each Pokémon has a catch difficulty — stronger Pokémon are harder to catch.

Example:

catch pikachu


If successful, Pikachu is added to your Pokédex.

pokedex

Shows all Pokémon you have caught.

inspect <pokemon>

Displays detailed info about a Pokémon you have already caught:

ID

Height / Weight

Stats

Types

Example:

inspect pikachu

🧠 How Caching Works

Your API client uses an internal in-memory cache that:

Stores API responses keyed by URL or Pokémon name

Automatically expires entries after a configurable time

Prevents re-fetching the same Pokémon or location repeatedly

Makes commands like inspect instantaneous after the first fetch

The caching logic lives in internal/pokeapi/cache.go.

🔧 How the API Client Works

The client is created in config.go:

client := pokeapi.NewClient(5 * time.Minute)


Features:

Shared HTTP client with timeout

Thread-safe caching

/pokemon/{name} endpoint

/location-area pagination

/location-area/{name} exploration

Automatic JSON unmarshalling into Go structs

🏗 Code Files Explained
main.go

Starts the program and launches the REPL.

repl.go

Handles:

Reading input

Splitting commands

Mapping commands to functions

Looping REPL

User feedback

config.go

Contains:

Global config struct

PokeAPI client

Pagination state

internal/pokeapi/pokeapi.go

Handles:

HTTP GET

Caching wrapper

Pokémon fetching

Location fetching

JSON parsing

internal/pokeapi/types.go

Defines all API models:

Pokémon

Stats

Types

Location areas

And more
