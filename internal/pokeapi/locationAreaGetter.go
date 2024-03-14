package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	//"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func RequestData() {

	for i := 0; i < 5; i++ {
		limit := 20
		offset := 20 * i // this should increment to move to the next page
		url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/?offset=%v&limit=%v", offset, limit)

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
		fmt.Printf("Call %v\n\n\n", i+1)
		fmt.Printf("%s", body)
		fmt.Println()

		// Meant to do some Json Unmarshalling thing here.
		// Probably rework the entire function, but I just want it working as a start
	}
}
