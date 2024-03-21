package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/Daniel-Burbridge-Developer/pokedexcli/internal/cache"
)

type PokeClient struct {
	pokeCache pokecache.Cache
}

func NewClient() *PokeClient {
	pc := PokeClient{}
	pc.pokeCache = *pokecache.NewCache(5)
	return &pc
}

func (pc *PokeClient) RequestLocationData(url any) ([]byte, error) {
	val, ok := pc.pokeCache.Get(fmt.Sprint(url))
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

		pc.pokeCache.Add(fmt.Sprint(url), body)
		return body, nil
	}

	return val, nil
}
