package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Relation struct {
	Id             int
	DatesLocations map[string][]string
}
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
	Concert      Relation
}

func main() {

	var links API
	var Artists []Artist
	// var Relations []Relation
	//var Relation Relation
	linkAPI := "https://groupietrackers.herokuapp.com/api"
	jsonErr := json.Unmarshal(openLink(linkAPI), &links)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//fmt.Println(links.Artists)

	jsonErr = json.Unmarshal(openLink(links.Artists), &Artists)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	//fmt.Println(Artists)
	for i, value := range Artists {

		var rel Relation
		jsonErr = json.Unmarshal(openLink(value.Relations), &rel)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
		Artists[i].Concert = rel
	}
	fmt.Println(Artists)

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
