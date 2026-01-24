package tools

import (
	"context"
	"fmt"
	"log"

	tools_internals "github.com/muskiteer/GoCP/server/tool_internals"
)

func FetchonlineData(ctx context.Context, args map[string]any) (any, error) {

	query, ok := args["query"].(string)
	if !ok || query == "" {
		return nil, fmt.Errorf("query must be a non-empty string")
	}

	title1, err := tools_internals.SearchTitle(query)
	if err != nil {
		return nil, err
	}

	wiki1, err := tools_internals.FetchSummary(title1)
	if err != nil {
		return nil, err
	}

	wiki := "Wikipedia Result 1:\n" + wiki1

	duck, err := tools_internals.SearchDuckDuckGo(query)
	if err != nil {
		return nil, err
	}

	result := fmt.Sprintf(
		"Wikipedia Summary:\n%s\n\nDuckDuckGo Search Results:\n%s",
		wiki,
		duck,
	)
	log.Println(result)

	return result, nil
}
