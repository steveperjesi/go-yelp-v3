package yelp

import (
    "fmt"

    "github.com/spf13/viper"
)

type AuthOptions struct {
    ApiKey string `json:"api_key"`
}

func GetApiKey() (string, error) {
    viper.SetConfigName("config")
    viper.AddConfigPath("./")
    viper.AddConfigPath("../") // This is needed to run the tests
    viper.AutomaticEnv()
    viper.SetConfigType("yml")

    if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("Error reading config file, %s", err)
        return "", err
    }

    apiKey := viper.GetString("YELP_V3_API_KEY")

    return apiKey, nil
}
