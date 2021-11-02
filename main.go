package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type API struct {
	Artists   string
	Locations string
	Dates     string
	Relation  string
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	MeberStr     string
	CreationDate int
	FirstAlbum   string
	Focations    string
	ConcertDates string
	Relations    string
}

func main() {

	var links API
	linkAPI := "https://groupietrackers.herokuapp.com/api"
	jsonErr := json.Unmarshal(openLink(linkAPI), &links)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(links)
}
func openLink(linkAPI string) []byte { // read file by http:
	response, err := http.Get(linkAPI)
	check(err)
	Body, err := io.ReadAll(response.Body)
	check(err)
	defer response.Body.Close()
	return Body
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
