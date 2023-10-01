package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const RequestURL = "https://hdrezka.co"

type Req struct {
	Title     string `json:"title" validate:"required"`
	Voiceover string `json:"voiceover" validate:"required"`
}

var request Req

func main() {
	ticker := time.NewTicker(15 * time.Minute)

	getInput()
	process()
	for range ticker.C {
		process()
	}
}

func process() {
	responseBody := makeRequest()

	content, err := io.ReadAll(responseBody)
	if err != nil {
		fmt.Printf("unable to read body: %v", err)
		os.Exit(1)
	}

	if isSeriesReleased := parseContent(string(content), request); isSeriesReleased == true {
		playSound()
		sendPush(request)
	}
}

func getInput() {
	fmt.Printf("Enter title of the series: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	request.Title = scanner.Text()

	fmt.Printf("Enter voiceover: ")
	scanner.Scan()
	request.Voiceover = scanner.Text()
}

func makeRequest() io.ReadCloser {
	res, err := http.Get(RequestURL)
	if err != nil {
		fmt.Errorf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Errorf("wrong status code: %v\n", res.Status)
		os.Exit(1)
	}

	return res.Body
}

func parseContent(data string, request Req) bool {
	updates := strings.Split(data, "<a class=\"b-seriesupdate__block_list_link\"")

	for _, series := range updates {
		if strings.Contains(series, request.Title) && strings.Contains(series, request.Voiceover) {
			fmt.Printf("Updates are here!\n")
			fmt.Println(composeLink(series))

			return true
		}
	}

	fmt.Printf("There is no updates. I'll check %v in %v voiceover later\n", request.Title, request.Voiceover)

	return false
}

func composeLink(series string) string {
	path := strings.Split(series, "href=\"")
	path = strings.Split(path[1], "\">")

	return RequestURL + path[0]
}
