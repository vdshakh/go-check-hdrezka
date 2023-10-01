package parser

import (
	"fmt"
	"github.com/go-check-hdrezka/request"
	"strings"
)

const RequestURL = "https://hdrezka.co"

func ParseContent(data string, userRequest request.Req) bool {
	updates := strings.Split(data, "<a class=\"b-seriesupdate__block_list_link\"")

	for _, series := range updates {
		if strings.Contains(series, userRequest.Title) && strings.Contains(series, userRequest.Voiceover) {
			fmt.Printf("Updates are here!\n")
			fmt.Println(composeLink(series))

			return true
		}
	}

	fmt.Printf("There is no updates. I'll check %v in %v voiceover later\n", userRequest.Title, userRequest.Voiceover)

	return false
}

func composeLink(series string) string {
	path := strings.Split(series, "href=\"")
	path = strings.Split(path[1], "\">")

	return RequestURL + path[0]
}
