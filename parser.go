package main

import (
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Movie defines the fields included in api response
type Movie struct {
	Title       string   `json:"title"`
	ReleaseYear int      `json:"release_year"`
	Actors      []string `json:"actors"`
	Poster      string   `json:"poster"`
	SimilarIds  []string `json:"similar_ids"`
}

func ParseDom(httpBody io.Reader) (*Movie, error) {
	m := &Movie{}

	doc, err := goquery.NewDocumentFromReader(httpBody)
	if err != nil {
		log.Println(err)
		return nil, err
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
