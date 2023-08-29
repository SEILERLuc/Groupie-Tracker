package dataAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ArtistsHTML struct { // Struct to display in the web page
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

type Artists struct { // Structs of the Artists_Locations_Dates_Relations
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct { // Locations_data
	ID               int      `json:"id"`
	ConcertLocations []string `json:"locations"`
	ConcertDates     string   `json:"dates"`
}
type LocationsIndex struct { // Delete index for locations
	Index []Locations `json:"index"`
}

type Dates struct { // Dates_data
	ID    int `json:"id"`
	Dates []string `json:"dates"`
}
type DatesIndex struct { // Struct to delete index for dates
	Index []Dates `json:"index"`
}

type Relations struct {
	ID            int                 `json:"id"`
	DateLocations map[string][]string `json:"datesLocations"` // dictionary
}
type RelationsIndex struct { // Struct to delete index for relations
	Index []Relations `json:"index"`
}

// The variables we need to have data and display them
var (
	url           = "https://groupietrackers.herokuapp.com/api" // the URL to find the API's data
	artists       = []Artists{}                                 // array for the struct Artists
	locationsData = LocationsIndex{}
	datesData     = DatesIndex{}
	relationsData  = RelationsIndex{}
	a             ArtistsHTML
)

func ArtistsDatas() {
	urlA := url + "/artists" // go to the URL and add "/artists" to go to the artists page
	resp, err := http.Get(urlA) // errors
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	htmlA, err := ioutil.ReadAll(resp.Body) // read the data of the URL
	if err != nil {
		panic(err)
	}
	json.Unmarshal(htmlA, &artists)
}

func LocationsData() {
	urlL := url + "/locations" // go to the URL and add "/artists" to go to the artists page
	resp, err := http.Get(urlL) // errors
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	htmlLocations, err := ioutil.ReadAll(resp.Body) // read the data of the URL
	if err != nil {
		panic(err)
	}
	json.Unmarshal(htmlLocations, &locationsData) // unmarshal the data and add them to the array artists
}

func DatesData() {
	urlL := url + "/dates" // go to the URL and add "/artists" to go to the artists page
	resp, err := http.Get(urlL) // errors
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	htmlLocations, err := ioutil.ReadAll(resp.Body) // read the data of the URL
	if err != nil {
		panic(err)
	}
	json.Unmarshal(htmlLocations, &datesData) // unmarshal the data and add them to the array artists
}

func RelationsData() {
	urlL := url + "/relation" // go to the URL and add "/artists" to go to the artists page
	// var relationsData = RelationsIndex{}
	resp, err := http.Get(urlL) // errors
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	htmlRelations, err := ioutil.ReadAll(resp.Body) // read the data of the URL
	if err != nil {
		panic(err)
	}
	json.Unmarshal(htmlRelations, &relationsData) // unmarshal the data and add them to the array artists
}

func ArtistHTML() []ArtistsHTML {
	tmp := []ArtistsHTML{}
	ArtistsDatas()
	LocationsData()
	DatesData()
	RelationsData()
	for i := range artists {
		a.ID = artists[i].ID
		a.Image = artists[i].Image
		a.Name = artists[i].Name
		a.Members = artists[i].Members
		a.CreationDate = artists[i].CreationDate
		a.FirstAlbum = artists[i].FirstAlbum
		a.Locations = locationsData.Index[i].ConcertLocations
		a.ConcertDates = datesData.Index[i].Dates
		a.Relations = relationsData.Index[i].DateLocations
		tmp = append(tmp, a)
	}
	return tmp
}
