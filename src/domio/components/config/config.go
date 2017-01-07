package config

import (
    "encoding/json"
    "os"
    "log"
    "path/filepath"
)

type Configuration struct {
    AWS_ACCESS_KEY_ID     string `json:"AWS_ACCESS_KEY_ID"`
    AWS_SECRET_ACCESS_KEY string `json:"AWS_SECRET_ACCESS_KEY"`
    DOMIO_DB_USER         string `json:"DOMIO_DB_USER"`
    DOMIO_DB_NAME         string `json:"DOMIO_DB_NAME"`
    DOMIO_DB_PASSWORD     string `json:"DOMIO_DB_PASSWORD"`
}

func LoadConfig() Configuration {
    dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

    file, _ := os.Open(dir + "/config.json")

    decoder := json.NewDecoder(file)
    config := Configuration{}
    err := decoder.Decode(&config)
    if err != nil {
        log.Fatalln("error:", err)
    }
    return config
}