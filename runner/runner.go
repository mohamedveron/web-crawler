package runner

func (cw *Crawler) Run(){

	ch := make(chan string)

	cw.VisitedPages[cw.Source] =  true

	go cw.pageProcessor(ch, cw.Webpages)
	go cw.pageProcessor(ch, cw.Webpages)
	go cw.pageProcessor(ch, cw.Webpages)

	cw.mux.Lock()
	for link, _ := range cw.VisitedPages{
		ch <- link
	}
	cw.mux.Unlock()
}

func (cw *Crawler) pageProcessor(pages <-chan string, webpages map[string][]string){

	for page := range pages {

		if links, ok := webpages[page]; ok{

			for _, link := range links{

				if _, ok := cw.VisitedPages[link]; !ok {
					cw.VisitedPages[link] = true
				}

			}
		}
	}
}
