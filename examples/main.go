package main

import (
	"fmt"

	"github.com/steveperjesi/go-yelp-v3/yelp"
)

func main() {
	apiKey, err := yelp.GetApiKey()
	if err != nil {
		panic(err)
	}

	client, err := yelp.NewClient(apiKey)
	if err != nil {
		panic(err)
	}

	results, err := client.Search(yelp.SearchOptions{
		Term:      "restaurants",
		Latitude:  36.0813328,
		Longitude: -115.3161651,
		SortBy:    "distance",
		// Location: "las vegas",
		// Radius:    40000, // 40000 meters is the max allowed value
		// Limit:  10,
		// Categories: "localservices",
		// OpenNow:   true,
		// Offset:    10,
		// Price:  "1,2,3,4",  // 1 = $, 2 = $$, 3 = $$$, 4 = $$$$
		// Attributes: "hot_and_new",
		// Locale:     "en_US",
		// OpenAt:    1572494399,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("\nTOTAL: ", results.Total)
	fmt.Println("\nREGION: ", results.Region)

	for _, biz := range results.Businesses {
		fmt.Println("Name\t\t", biz.Name)
		fmt.Println("ID\t\t", biz.Id)
		fmt.Println("Alias\t\t", biz.Alias)
		fmt.Println("Rating\t\t", biz.Rating)
		fmt.Println("Price\t\t", biz.Price)
		fmt.Println("IsClosed\t", biz.IsClosed)
		fmt.Println("Url\t\t", biz.Url)
		fmt.Println("Distance\t", biz.Distance)
		fmt.Println("ReviewCount\t", biz.ReviewCount)
		fmt.Println("Latitude\t", biz.Coordinates.Latitude)
		fmt.Println("Longitude\t", biz.Coordinates.Longitude)

		fmt.Println("Phone\t\t", biz.Phone)

		fmt.Println("Address1\t", biz.Location.Address1)
		fmt.Println("Address2\t", biz.Location.Address2)
		fmt.Println("Address3\t", biz.Location.Address3)
		fmt.Println("City\t\t", biz.Location.City)
		fmt.Println("State\t\t", biz.Location.State)
		fmt.Println("ZipCode\t\t", biz.Location.ZipCode)
		fmt.Println("Country\t\t", biz.Location.Country)

		cats := yelp.CategoriesToString(biz.Categories)
		fmt.Println("Categories\t", cats)
		fmt.Println("-----")
	}
}
