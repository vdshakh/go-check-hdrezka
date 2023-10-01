package request

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const RequestURL = "https://hdrezka.co"

type Req struct {
	Title     string `json:"title" validate:"required"`
	Voiceover string `json:"voiceover" validate:"required"`
}

func MakeRequest() io.ReadCloser {
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
