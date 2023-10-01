package main

import (
	"bufio"
	"fmt"
	"github.com/go-check-hdrezka/notification"
	"github.com/go-check-hdrezka/parser"
	"github.com/go-check-hdrezka/request"
	"io"
	"os"
	"time"
)

var userRequest request.Req

func main() {
	ticker := time.NewTicker(15 * time.Minute)

	getInput()
	process()
	for range ticker.C {
		process()
	}
}

func process() {
	responseBody := request.MakeRequest()

	content, err := io.ReadAll(responseBody)
	if err != nil {
		fmt.Printf("unable to read body: %v", err)
		os.Exit(1)
	}

	if isSeriesReleased := parser.ParseContent(string(content), userRequest); isSeriesReleased == true {
		notification.PlaySound()
		notification.SendPush(userRequest)
	}
}

func getInput() {
	fmt.Printf("Enter title of the series: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userRequest.Title = scanner.Text()

	fmt.Printf("Enter voiceover: ")
	scanner.Scan()
	userRequest.Voiceover = scanner.Text()
}
