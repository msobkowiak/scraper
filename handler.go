package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/mux"
)

const baseUrl = "http://www.amazon.de/gp/product/"

// Movie defines the fields included in api response
type Movie struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

// GetMovie handles the request. Crawls amazon prime page for the movie specified
// by the amazon_id, finds information about the movie and return it in the response.
func GetMovie(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	aid := params["amazon_id"]

	respCode := http.StatusOK
	resp, err := http.Get(baseUrl + aid)
	if err != nil {
		respCode = http.StatusBadRequest
		log.Println(err)
	}
	defer resp.Body.Close()

	m, err := parseDom(resp.Body)
	if err != nil {
		respCode = http.StatusInternalServerError
	}

	writeResponse(m, respCode, w)
}

func parseDom(httpBody io.Reader) (Movie, error) {
	m := Movie{}

	doc, err := goquery.NewDocumentFromReader(httpBody)
	if err != nil {
		log.Println(err)
		return Movie{}, err
	}

	// release year
	n := doc.Find("span.release-year")
	releaseYear, err := strconv.Atoi(n.Text())
	if err != nil {
		log.Println(err)
	}
	m.ReleaseYear = releaseYear

	// title
	n = doc.Find("h1#aiv-content-title")
	n.Children().Remove()
	m.Title = strings.Trim(n.Text(), "\n ")

	// poster
	n = doc.Find(".dp-meta-icon-container").ChildrenFiltered("img")
	if link, ok := n.Attr("src"); ok {
		m.Poster = link
	}

	// actors
	s := doc.Find("dd").First().Text()
	s = strings.Trim(s, "\n ")
	actors := strings.Split(s, ", ")
	m.Actors = actors

	// similar links
	n = doc.Find("li").Each(func(i int, s *goquery.Selection) {
		if id, ok := s.Attr("data-asin"); ok {
			m.SimilarIds = append(m.SimilarIds, id)
		}
	})

	return m, nil
}

func writeResponse(data interface{}, respCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(respCode)
	if respCode == http.StatusOK {
		json.NewEncoder(w).Encode(&data)
	}
}
