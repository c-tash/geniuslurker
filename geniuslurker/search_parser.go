package geniuslurker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const geniusBaseURL = "https://api.genius.com/search"
const geniusToken = "Bearer us4hrg63-ZYFCFmecW9iS3nXoLs5rkTkFIGhECwNHtMda0GyCINDkleGdmiKjAmx"

type ResultJSON struct {
	FullTitle string `json:"full_title"`
	Url       string `json:"url"`
}

type HitJSON struct {
	Result ResultJSON `json:"result"`
}

type ResponseJSON struct {
	Hits []HitJSON `json:"hits"`
}

type BaseJSON struct {
	Response ResponseJSON `json:"response"`
}

// Returns search results
func GetSearchResults(searchString string) []ResultJSON {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", geniusBaseURL, nil)
	req.Header.Add("Authorization", geniusToken)
	q := req.URL.Query()
	q.Add("q", searchString)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		panic(err)
	}

	var parsedJSON BaseJSON
	err = json.NewDecoder(resp.Body).Decode(&parsedJSON)
	if err != nil {
		fmt.Println("JSON parsing error:", err)
		panic(err)
	}

	results := make([]ResultJSON, len(parsedJSON.Response.Hits), len(parsedJSON.Response.Hits))
	for index, element := range parsedJSON.Response.Hits {
		results[index] = element.Result
	}
	return results
}
