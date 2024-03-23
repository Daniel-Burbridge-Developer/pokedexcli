package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"

	pokecache "github.com/Daniel-Burbridge-Developer/pokedexcli/internal/cache"
	"github.com/Daniel-Burbridge-Developer/pokedexcli/internal/pokedex"
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
	PokeDex   pokedex.Pokedex
}

func NewClient() *PokeClient {
	// fmt.Println(Red, "Making a New PokeClient", Reset)
	pc := PokeClient{}
	pc.pokeCache = *pokecache.NewCache()
	pc.PokeDex = *pokedex.New()
	go pc.pokeCache.ReapLoop(30)
	// fmt.Println(Red, "Returning a new PokeClient", Reset)
	return &pc
}

/// SO ALL of the below functions are actually identical thinking about it. Why do I have multiple functions for that???
/// Like it's working for now and I don't want to dig around my files fixing things, but uh yeah.

func (pc *PokeClient) RequestLocationData(url any) ([]byte, error) {
	// fmt.Println(Red, "Checking if in Cache", Reset)
	val, ok := pc.pokeCache.Get(fmt.Sprint(url))
	if !ok {
		fmt.Println(Yellow, "MISSING FROM CACHE", Reset)
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
		// fmt.Println(body)
		return body, nil
	}

	fmt.Println(Magenta, "FOUND IN CACHE", Reset)
	return val, nil
}

func (pc *PokeClient) ExploreLocation(url any) ([]byte, error) {

	// fmt.Println(Red, "EXPLORING THE LOCATION", Reset)

	// Check if in cache
	val, ok := pc.pokeCache.Get(fmt.Sprint(url))

	if !ok {
		fmt.Println(Yellow, "MISSING FROM CACHE", Reset)
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

	fmt.Println(Magenta, "FOUND IN CACHE", Reset)
	return val, nil

}

func (pc *PokeClient) RequestPokemonData(url any) ([]byte, error) {

	// fmt.Println(Red, "Requesting the pokemon", Reset)

	val, ok := pc.pokeCache.Get(fmt.Sprint(url))

	if !ok {
		fmt.Println(Yellow, "MISSING FROM CACHE", Reset)
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

	fmt.Println(Magenta, "FOUND IN CACHE", Reset)
	return val, nil
}
