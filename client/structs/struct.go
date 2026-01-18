package structs

import (
	
)

type TagsResponse struct {
	Models []struct {
		Name string `json:"name"`
	} `json:"models"`
}