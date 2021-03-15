package runner

type Crawler struct {
	Source       string
	Destination  string
	Webpages     map[string][]string
	VisitedPages map[string]bool
}

func NewCrawler(source string, destination  string, webpages map[string][]string, visitedPages map[string]bool) *Crawler{

	return &Crawler{
		Source:       source,
		Destination:  destination,
		Webpages:     webpages,
		VisitedPages: visitedPages,
	}
}