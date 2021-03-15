package main

import (
	"github.com/mohamedveron/web-crawler/runner"
)

func main() {

	webpages := make(map[string][]string)
	var visitedLinks = make(map[string]bool)

	webpages["https://bigplato.tilda.ws/home"] = []string{
		"https://bigplato.tilda.ws/team",
		"https://bigplato.tilda.ws/solution",
	}

	webpages["https://bigplato.tilda.ws/team"] = []string{
		"https://bigplato.tilda.ws/team/members",
		"https://teambuilding/team/workingspace",
	}

	webpages["https://bigplato.tilda.ws/solution"] = []string{
		"https://bigplato.tilda.ws/solution/Corporateprofilemanagement",
		"https://bigplato.tilda.ws/solution/Regulatorytechnology",
		"https://bigplato.tilda.ws/solution/Corporategovernancesolution",
	}

	webpages["https://bigplato.tilda.ws/solution/Corporateprofilemanagement"] = []string{
		"https://bigplato.tilda.ws/solution/Corporateprofilemanagement/dolphin",
	}
	concurrentWorkers := 3

	crawler := runner.NewCrawler("https://bigplato.tilda.ws/home", "https://bigplato.tilda.ws/services/Corporateprofilemanagement/dolphin",
		"https://bigplato.tilda.ws/", concurrentWorkers, webpages, visitedLinks)

	crawler.Run()

	//fmt.Println(visitedLinks)

}
