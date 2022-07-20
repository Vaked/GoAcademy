package scraper

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	"io"
)

const(
	storiesEndpoint string = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty";
	detailsEndpoint string = "https://hacker-news.firebaseio.com/v0/item/"
)

type Stories struct{
	Stories []Story `json:"top_stories"`
}

func (s *Stories) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(s)
}

type Story struct{
	Title string `json:"title"`
	Score int `json:"score"`
}


func NewStories() *Stories{
	return &Stories{}
}

func NewStory() *Story{
	return &Story{}
}

func httpClient() *http.Client{
	httpClient := http.Client{
		Timeout: time.Second * 10,

	}

	return &httpClient
}

func getTopTenStoriesFromJSON(stories []int) []int{
	var result []int
	for i := 0;i < 10; i++{
		result = append(result, stories[i])
	}

	return result
}

func (s *Stories)GetTopHackerNewsStories() []int{
	httpClient := httpClient()

	req, err := http.NewRequest(http.MethodGet, storiesEndpoint, nil)
	if err != nil{
		log.Fatal(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var	scraper []int
	jsonErr := json.Unmarshal(body, &scraper)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return getTopTenStoriesFromJSON(scraper)
}

func CreatDetailsUrlFromStory(storyId int) string{
	storyIdString := strconv.Itoa(storyId)
	url := detailsEndpoint + storyIdString + ".json?print=pretty"
	return url
}

func (s *Stories)GenerateDetailsUrls(storyIdList []int) []string{
	stories := s.GetTopHackerNewsStories()

	var urls []string
	for _, val := range stories{
		urls = append(urls, CreatDetailsUrlFromStory(val))
	}

	return urls
}

func (s Stories)GetDetailsFromTopStories(urls []string) *Stories{
	httpClient := httpClient()
	stories := NewStories()

	for _, url := range urls{
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil{
			log.Fatal(err)
		}

		res, err := httpClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var	story = NewStory()
		jsonErr := json.Unmarshal(body, &story)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		stories.Stories = append(stories.Stories, *story)
	}	
	return stories
}