package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"context"
)

type searchResponse struct {
	Query struct {
		Search []struct {
			Title string `json:"title"`
		} `json:"search"`
	} `json:"query"`
}

func searchTitle(question string) (string, error) {
	base := "https://en.wikipedia.org/w/api.php"

	params := url.Values{}
	params.Set("action", "query")
	params.Set("list", "search")
	params.Set("srsearch", question)
	params.Set("format", "json")

	req, err := http.NewRequest(
		"GET",
		base+"?"+params.Encode(),
		nil,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set(
		"User-Agent",
		"GoCP/1.0 (contact@example.com)",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Query.Search) == 0 {
		return "", fmt.Errorf("no wikipedia results")
	}

	return result.Query.Search[0].Title, nil
}



type summaryResponse struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func fetchSummary(title string) (*summaryResponse, error) {
	escaped := url.PathEscape(title)

	apiURL :=
		"https://en.wikipedia.org/api/rest_v1/page/summary/" + escaped

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"User-Agent",
		"GoCP/1.0 (contact@example.com)",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var summary summaryResponse
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return nil, err
	}

	return &summary, nil
}


func FetchWikipediaData(ctx context.Context, args map[string]any) (any, error){
	query, ok := args["query"].(string)
	if !ok || query == "" {
		return nil, fmt.Errorf("query must be a non-empty string")
	}
	title, err := searchTitle(query)
	if err != nil {
		return nil, err
	}

	return fetchSummary(title)
}