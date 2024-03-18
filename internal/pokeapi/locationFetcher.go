package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/Daniel-Burbridge-Developer/pokedexcli/internal/cache"
)

func RequestLocationData(url any) ([]byte, error) {
	cache := pokecache.NewCache(20)
	val, ok := cache.Get(fmt.Sprint(url))
	if !ok {
		res, err := http.Get(fmt.Sprint(url))
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(res.Body)
		res.Body.Close()
		if res.StatusCode > 299 {
			log.Fatalf("response failed with status code: %d\nbody: %s\n", res.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}

		cache.Add(fmt.Sprint(url), body)
		return body, nil
	}

	return val, nil
}
