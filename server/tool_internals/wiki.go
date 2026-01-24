package tools_internals

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)



type searchResponse struct {
	Query struct {
		Search []struct {
			Title string `json:"title"`
		} `json:"search"`
	} `json:"query"`
}


func SearchTitle(question string) (string,  error) {

	time.Sleep(200 * time.Millisecond)

	base := "https://en.wikipedia.org/w/api.php"

	params := url.Values{}
	params.Set("action", "query")
	params.Set("list", "search")
	params.Set("srsearch", question)
	params.Set("format", "json")
	params.Set("srlimit", "2")

	req, err := http.NewRequest(
		http.MethodGet,
		base+"?"+params.Encode(),
		nil,
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "GoCP/1.0 (contact@example.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wikipedia API returned status %d", resp.StatusCode)
	}

	var result searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Query.Search) == 0 {
		return "", fmt.Errorf("no wikipedia results")
	}

	if len(result.Query.Search) == 1 {
		return result.Query.Search[0].Title, nil
	}

	return result.Query.Search[0].Title,  nil
}



type summaryResponse struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func FetchSummary(title string) (string, error) {

	time.Sleep(200 * time.Millisecond)

	escaped := url.PathEscape(title)
	apiURL := "https://en.wikipedia.org/api/rest_v1/page/summary/" + escaped

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "GoCP/1.0 (contact@example.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wikipedia summary API returned status %d", resp.StatusCode)
	}

	var summary summaryResponse
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Title: %s\n\n", summary.Title))
	sb.WriteString(fmt.Sprintf("Summary: %s\n", summary.Extract))
	sb.WriteString("\n")

	return sb.String(), nil
}
