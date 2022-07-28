package scraper

import (
	"encoding/json"
	"log"
	"net/http"
	"scraper/scraper"
)

type ScraperAPI struct{
	l *log.Logger
}

func NewScraperAPI(l *log.Logger) *ScraperAPI{
	return &ScraperAPI{l}
}

func (s *ScraperAPI) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.getStories(rw, r)
		return 
	}

	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}


func (s *ScraperAPI) getStories(rw http.ResponseWriter, r *http.Request){
	s.l.Println("Handle GET Stories")

	scraper := scraper.NewStories()
	stories := scraper.GetTopHackerNewsStories()
	urls := scraper.GenerateDetailsUrls(stories)
	storyList := scraper.GetDetailsFromTopStories(urls)
	
	js, err := json.MarshalIndent(storyList, "", "    ")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	}
	  
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(js)
}