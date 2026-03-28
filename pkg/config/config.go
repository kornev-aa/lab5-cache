package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    StorageType string  `json:"storage_type"`
    FilePath    string  `json:"file_path"`
    DefaultLat  float64 `json:"default_lat"`
    DefaultLon  float64 `json:"default_lon"`
}

func Load(path string) (*Config, error) {
    file, err := os.Open(path)
    if err != nil {
        return &Config{
            StorageType: "file",
            FilePath:    "./location.json",
            DefaultLat:  53.6688,
            DefaultLon:  23.8223,
        }, nil
    }
    defer file.Close()
    var cfg Config
    if err := json.NewDecoder(file).Decode(&cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
