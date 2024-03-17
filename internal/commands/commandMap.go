package commands

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Daniel-Burbridge-Developer/pokedexcli/models"
)

func CommandMap(config models.Config) error {
	//SCRAP THIS - JUST FOR QUICK TEST - GOING TO BED
	// IS WORKING until hitting hard-coded stuff. But want to redo anyway means other stuff is working
	res, err := http.Get(fmt.Sprint(config.Next))
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

	fmt.Printf("%s", body)
	return nil
}
