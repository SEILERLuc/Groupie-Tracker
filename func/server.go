package _func

import (
	"fmt"
	"groupie-tracker/dataAPI"
	"groupie-tracker/utils"
	"html/template"
	"net/http"
	"strconv"
)

var data []dataAPI.ArtistsHTML // A variable to call data from the other file

func Serve() { // the function that create the server and the URL to display the data
	data = dataAPI.ArtistHTML()
	server := http.NewServeMux() // a website with two pages
	server.HandleFunc("/", DisplayAll)
	server.HandleFunc("/artist", DisplayArtist)
	server.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./CSS")))) // to call the css from gohtml files
	server.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./images")))) // to call the css from gohtml files
	fmt.Println("Server listening on http://localhost:8000")
	http.ListenAndServe(":8000", server) // load the server
}

func DisplayAll(w http.ResponseWriter, r *http.Request) { // that display all the data to the website
	if r.Method == "POST" {
		userData := []dataAPI.ArtistsHTML{}
		userSearch := r.FormValue("search")
		filterDate := r.FormValue("date")
		filterMember := r.FormValue("member")
		filterDateInt, _ := strconv.Atoi(filterDate)
		filterMemberInt, _ := strconv.Atoi(filterMember)
		if userSearch != "" && len(userData) == 0 {
			userData = Search(userSearch)
		}
		if filterDate != "" {
			userData = DateFilter(userData, filterDateInt)
		}
		if filterMember != "" {
			userData = MemberFilter(userData, filterMemberInt)
		}
		tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
		_ = tmpl.Execute(w, struct {
			Artists []dataAPI.ArtistsHTML
		}{
			Artists: userData,
		})
	} else {
		tmpl := template.Must(template.ParseFiles("templates/index.gohtml"))
		_ = tmpl.Execute(w, struct {
			Artists []dataAPI.ArtistsHTML
		}{
			Artists: data,
		})
	}
}

func DisplayArtist(w http.ResponseWriter, r *http.Request) { // tha display the data of just 1 artist choose by the user
	artist := r.FormValue("artist")
	artistInt, _ := strconv.Atoi(artist)
	userArtist := []dataAPI.ArtistsHTML{}
	for i := range data {
		if data[i].ID == artistInt {
			userArtist = append(userArtist, data[i])
		}
	}
	tmpl := template.Must(template.ParseFiles("templates/artist.gohtml"))
	_ = tmpl.Execute(w, struct {
		Artists []dataAPI.ArtistsHTML
	}{
		Artists: userArtist,
	})
}

func Search(userSearch string) []dataAPI.ArtistsHTML {
	var c = []dataAPI.ArtistsHTML{}
	for i := range data {
		if utils.CaseInsensitiveContains(data[i].Name, userSearch) {
			c = append(c, data[i])
			break
		} else {
			for j := range data[i].Members {
				if utils.CaseInsensitiveContains(data[i].Members[j], userSearch) {
					c = append(c, data[i])
				}
			}
			for j := range data[i].Locations {
				if utils.CaseInsensitiveContains(data[i].Locations[j], userSearch) {
					c = append(c, data[i])
				}
			}
			if utils.CaseInsensitiveContains(data[i].FirstAlbum, userSearch) {
				c = append(c, data[i])
			}
		}
	}
	return c
}

func DateFilter(c []dataAPI.ArtistsHTML, filterDate int) []dataAPI.ArtistsHTML {
	for i := range data {
		if data[i].CreationDate >= filterDate {
			c = append(c, data[i])
		}
	}
	return c
}

func MemberFilter(c []dataAPI.ArtistsHTML, filterMember int) []dataAPI.ArtistsHTML {
	for i := range data {
		if filterMember == len(data[i].Members) {
			c = append(c, data[i])
		}
	}
	return c
}
