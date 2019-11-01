package yelp

import (
    "fmt"

    "github.com/spf13/viper"
)

type AuthOptions struct {
    ApiKey string `json:"api_key"`
}

func GetAuthOptions() (AuthOptions, error) {
    authOptions := AuthOptions{}

    viper.SetConfigName("config")
    viper.AddConfigPath("./")
    viper.AddConfigPath("../") // This is needed to run the tests
    viper.AutomaticEnv()
    viper.SetConfigType("yml")

    fmt.Println(viper.ConfigFileUsed())

    if err := viper.ReadInConfig(); err != nil {
        fmt.Printf("Error reading config file, %s", err)
        return authOptions, err
    }

    authOptions.ApiKey = viper.GetString("YELP_API_KEY")

    return authOptions, nil
}
