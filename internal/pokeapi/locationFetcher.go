package pokeapi

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func RequestLocationData(url any) ([]byte, error) {

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

	return body, nil
}
