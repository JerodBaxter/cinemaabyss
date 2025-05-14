package main

import (
    "os"
    "net/http"
    "fmt"
    "strconv"
)

type Config struct {
    Port string
    MonolithURL string
    MoviesServiceURL string
    EventsServiceURL string
    GradualMigration bool
    MoviesMigrationPercent int
}

func GetConfig() Config {
    return Config{
        Port: getEnv("PORT", "8000"),
        MonolithURL: getEnv("MONOLITH_URL", "http://monolith:8080"),
        MoviesServiceURL: getEnv("MOVIES_SERVICE_URL", "http://movies-service:8081"),
        EventsServiceURL: getEnv("EVENTS_SERVICE_URL", "http://events-service:8082"),
        GradualMigration: getEnv("GRADUAL_MIGRATION", "true") == "true",
        MoviesMigrationPercent: parseInt(getEnv("MOVIES_MIGRATION_PERCENT", "50")),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

func parseInt(s string) int {
    if s == "" {
        return 50
    }
    val, err := strconv.Atoi(s)
    if err != nil {
        return 50
    }
    return val
}
