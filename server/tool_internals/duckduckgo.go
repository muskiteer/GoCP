package tools_internals

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)


func SearchDuckDuckGo(query string) (string, error) {

		maxResults := 3
	

	time.Sleep(300 * time.Millisecond)

	searchURL := "https://html.duckduckgo.com/html/?q=" +
		url.QueryEscape(query)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (GoCP DuckDuckGo Tool)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("duckduckgo returned status %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(
		"Search results for \"%s\":\n\n",
		query,
	))

	count := 0

	doc.Find("div.result").EachWithBreak(func(i int, s *goquery.Selection) bool {
		// Skip advertisements
		if s.HasClass("result--ad") {
			return true
		}

		title := strings.TrimSpace(s.Find("a.result__a").Text())
		snippet := strings.TrimSpace(s.Find("a.result__snippet").Text())
		link, exists := s.Find("a.result__a").Attr("href")

		if title == "" || snippet == "" || !exists {
			return true
		}

		link = cleanDuckDuckGoURL(link)

		count++
		sb.WriteString(fmt.Sprintf(
			"%d. %s\n   %s\n   Source: %s\n\n",
			count,
			title,
			snippet,
			link,
		))

		return count < maxResults
	})

	if count == 0 {
		return "", fmt.Errorf("no results found for query: %s", query)
	}

	return sb.String(), nil
}

func cleanDuckDuckGoURL(raw string) string {
	if !strings.Contains(raw, "uddg=") {
		return raw
	}

	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	decoded := u.Query().Get("uddg")
	if decoded == "" {
		return raw
	}

	result, err := url.QueryUnescape(decoded)
	if err != nil {
		return raw
	}

	return result
}
