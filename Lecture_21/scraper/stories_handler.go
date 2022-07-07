package scraper

// import "net/http"

type StoriesHandlerResponse struct {
	TopStories []Story `json:"top_stories,omitempty"`
}

