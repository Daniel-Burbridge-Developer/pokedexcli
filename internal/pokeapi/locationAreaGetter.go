package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
	//"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func RequestData() {

	config := models.Config{Next: "https://pokeapi.co/api/v2/location-area/?&limit=20", Previous: ""}
	locationsData := models.LocationsData{}

	for i := 0; i < 5; i++ {
		url := config.Next.(string)

		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(body, &config)
		json.Unmarshal(body, &locationsData)

		// fmt.Println("------------ START OF BODY --------------")
		// fmt.Printf("%s", body)
		// fmt.Println("------------ END OF BODY --------------")

		fmt.Println("------------ START OF CALL --------------")
		fmt.Println("------------ START OF CONFIG --------------")
		fmt.Printf("Config Next: %v\n", config.Next)
		fmt.Printf("Config Previous: %v\n", config.Previous)
		fmt.Println("------------ END OF CONFIG --------------")
		fmt.Println("------------ START OF LOCATIONS --------------")
		for i, result := range locationsData.Results {
			fmt.Printf("%v) Location Data Name: %s\n", i+1, result.Name)
			fmt.Printf("%v) Location Data URL: %s\n", i+1, result.URL)
		}
		fmt.Println("------------ END OF LOCATIONS --------------")
		fmt.Println("------------ END OF CALL --------------")
	}
}
