package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
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

var links API
var Artists []Artist

type Errors struct {
	Number  int
	Message string
}

var errResult string
var err int

func main() {
	getArtist()
	//fmt.Println(Artists)
	handleRequest()

}

// index page, if address != index, you are redirect to 404err func
func index(w http.ResponseWriter, r *http.Request) {
	tmpl, tmplErr := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if tmplErr != nil {
		err = 404
		errResult = "This page is not exist"
		w.WriteHeader(http.StatusNotFound)
	}
	if r.URL.Path != "/" {
		err = 404
		errResult = "This page is not exist"
		//err404(w, r)
		return
	} else {
		tmpl.ExecuteTemplate(w, "index", Artists)
	}
}

func err404(w http.ResponseWriter, r *http.Request) {
	tmpl, tmplErr := template.ParseFiles("templates/404.html", "templates/header.html", "templates/footer.html")
	dataErr := Errors{err, errResult}

	if tmplErr != nil {
		err = 404
		errResult = "This page is not exist"
		w.WriteHeader(http.StatusNotFound)
	}
	if err == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else if err == 400 {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	tmpl.ExecuteTemplate(w, "404", dataErr)

}
func handleRequest() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", index)
	//http.HandleFunc("/404", err404)
	log.Println("Server running ")
	http.ListenAndServe(":8080", nil)
	log.Println("Server running on: http://localhost:8080")

}

func getArtist() {

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
