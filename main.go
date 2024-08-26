package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"jonppenny.co.uk/DigitalOceanAPI/structs"
)

const VERSION string = "0.0.1"

func main() {
	//fmt.Printf("Version %s\n", VERSION)

	config := flag.String("config", "config.yml", "Configuration file")
	flag.Parse()

	viper.SetConfigFile(*config)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configuration file: %s\n", err)
	}

	req, _ := http.NewRequest("GET", "https://api.digitalocean.com/v2/droplets", nil)

	var bearer = "Bearer " + viper.GetString("api_key")

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var dropletsData structs.DropletsData
	err = json.Unmarshal(body, &dropletsData)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range dropletsData.Droplets {
		fmt.Printf("ID: %d, Name: %s\n", v.ID, v.Name)
	}
	fmt.Printf("You have %d total droplets", dropletsData.Meta.Total)
}
