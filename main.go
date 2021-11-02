package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
type API struct{
	Artists string
	Locations string
	Dates string
	Relation string
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
	//var artist []Artist

	

	response, err := http.Get("https://groupietrackers.herokuapp.com/api")
	check(err)

	fApi, err := io.ReadAll(response.Body)
	check(err)
	defer response.Body.Close()
	jsonErr := json.Unmarshal(fApi, &links)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(links)

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
