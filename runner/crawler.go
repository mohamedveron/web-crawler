package runner

import "sync"

type Crawler struct {
	Source       string
	Destination  string
	Webpages     map[string][]string
	VisitedPages map[string]bool
	mux  sync.Mutex
	baseUrl		string
	concurrentWorkers int
}

func NewCrawler(source , destination , baseUrl string, concurrentWorkers int, webpages map[string][]string, visitedPages map[string]bool) *Crawler{

	return &Crawler{
		Source:       source,
		Destination:  destination,
		Webpages:     webpages,
		VisitedPages: visitedPages,
		baseUrl: baseUrl,
		concurrentWorkers: concurrentWorkers,
	}
}