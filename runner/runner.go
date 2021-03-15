package runner

func (cw *Crawler) Run(){

	ch := make(chan string)

	cw.VisitedPages[cw.Source] =  true

	go cw.pageProcessor(ch, cw.Webpages)

	for link, _ := range cw.VisitedPages{
		ch <- link
	}
}

func (cw *Crawler) pageProcessor(pages <-chan string, webpages map[string][]string){

	for page := range pages {

		if links, ok := webpages[page]; ok{

			for _, link := range links{
				cw.VisitedPages[link] = true
			}
		}
	}
}
