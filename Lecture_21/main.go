package main

import (
	"fmt"
	"scraper/scraper"
)

func main() {
	scraper := scraper.NewStories()
	stories := scraper.GetTopHackerNewsStories()
	urls := scraper.GenerateDetailsUrls(stories)
	storyList := scraper.GetDetailsFromTopStories(urls)
	fmt.Println(storyList)

	
}