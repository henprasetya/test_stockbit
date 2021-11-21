package model

type Response struct {
	TotalResults string        `json:"totalResults"`
	Response     string        `json:"response"`
	Search       []SearchMovie `json:data`
}
