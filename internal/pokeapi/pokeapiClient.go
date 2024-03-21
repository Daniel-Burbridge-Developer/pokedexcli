package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/Daniel-Burbridge-Developer/pokedexcli/internal/cache"
)

// Got this from GPT to help change colors for console -- not entirely sure what the const () syntax is creating.
// It's the same as the import syntax, seems just like a way to do bulk variables?
// loading consts and build time I think
const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Magenta = "\033[35m"
)

type PokeClient struct {
	pokeCache pokecache.Cache
}

func NewClient() *PokeClient {
	// fmt.Println(Red, "Making a New PokeClient", Reset)
	pc := PokeClient{}
	pc.pokeCache = *pokecache.NewCache()
	go pc.pokeCache.ReapLoop(30)
	// fmt.Println(Red, "Returning a new PokeClient", Reset)
	return &pc
}

func (pc *PokeClient) RequestLocationData(url any) ([]byte, error) {
	// fmt.Println(Red, "Checking if in Cache", Reset)
	val, ok := pc.pokeCache.Get(fmt.Sprint(url))
	if !ok {
		fmt.Println(Yellow, "MISSING FROM LOWCATION DATA CACHE", Reset)
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
		// fmt.Println(Red, "Adding it to the cache!", Reset)
		pc.pokeCache.Add(fmt.Sprint(url), body)
		// fmt.Println(Red, "Added to the cache, time to send it back!", Reset)
		return body, nil
	}

	fmt.Println(Magenta, "FOUND IN LOCATION DATA CACHE", Reset)
	return val, nil
}

func (pc *PokeClient) ExploreLocation(url any) ([]byte, error) {
	fmt.Println("EXPLORING THE LOCATION")

	// Just here so my program doesn't error on compile prior to feat implemention
	bytes := make([]byte, 5)
	return bytes, nil
}
