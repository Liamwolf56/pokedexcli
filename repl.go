package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type cliCommand struct {
    name        string
    description string
    callback    func(*config) error
}

var commandRegistry map[string]cliCommand

func init() {
    commandRegistry = map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "map": {
            name:        "map",
            description: "View next 20 location areas",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "View previous 20 location areas",
            callback:    commandMapBack,
        },
    }
}

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    cfg := &config{}

    for {
        fmt.Print("Pokedex > ")
        if !scanner.Scan() {
            break
        }

        input := strings.ToLower(strings.TrimSpace(scanner.Text()))
        if input == "" {
            continue
        }

        args := strings.Fields(input)
        cmdName := args[0]

        cmd, ok := commandRegistry[cmdName]
        if !ok {
            fmt.Println("Unknown command.")
            continue
        }

        if err := cmd.callback(cfg); err != nil {
            fmt.Printf("Error: %v\n", err)
        }
    }
}

func commandHelp(cfg *config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")
    for _, cmd := range commandRegistry {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

func commandExit(cfg *config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandMap(cfg *config) error {
    url := "https://pokeapi.co/api/v2/location-area"
    if cfg.next != nil {
        url = *cfg.next
    }

    resp, err := fetchLocationAreas(url)
    if err != nil {
        return err
    }

    for _, loc := range resp.Results {
        fmt.Println(loc.Name)
    }

    cfg.next = resp.Next
    cfg.previous = resp.Previous
    return nil
}

func commandMapBack(cfg *config) error {
    if cfg.previous == nil {
        fmt.Println("You're on the first page.")
        return nil
    }

    resp, err := fetchLocationAreas(*cfg.previous)
    if err != nil {
        return err
    }

    for _, loc := range resp.Results {
        fmt.Println(loc.Name)
    }

    cfg.next = resp.Next
    cfg.previous = resp.Previous
    return nil
}

