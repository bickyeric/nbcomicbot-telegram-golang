// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EpisodeEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Episode `json:"node"`
}

type PageInfo struct {
	StartCursor string `json:"startCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}
