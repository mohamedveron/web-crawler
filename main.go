package main

import "fmt"


func crawler(source string, dest string, webpages map[string][]string) {
	fmt.Println(source, "->", dest)

	links := pageProcessor(source, webpages)

	for _, link := range links{

		crawler(link, dest, webpages)
	}

}

func pageProcessor(page string, webpages map[string][]string) []string {

	if links, ok := webpages[page]; ok{

		return links
	}

	return nil
}

func main() {

	webpages := make(map[string][]string)

	webpages["https://bigplato.tilda.ws/"] = []string{
		"https://bigplato.tilda.ws/team",
		"https://bigplato.tilda.ws/solution",
	}

	webpages["https://bigplato.tilda.ws/team"] = []string{
		"https://bigplato.tilda.ws/team/members",
		"https://bigplato.tilda.ws/team/workingspace",
	}

	webpages["https://bigplato.tilda.ws/solution"] = []string{
		"https://bigplato.tilda.ws/solution/Corporateprofilemanagement",
		"https://bigplato.tilda.ws/solution/Regulatorytechnology",
		"https://bigplato.tilda.ws/solution/Corporategovernancesolution",
	}

	webpages["https://bigplato.tilda.ws/solution/Corporateprofilemanagement"] = []string{
		"https://bigplato.tilda.ws/services/Corporateprofilemanagement/dolphin",
	}


	crawler("https://bigplato.tilda.ws/", "https://bigplato.tilda.ws/services/Corporateprofilemanagement/dolphin", webpages)

}
