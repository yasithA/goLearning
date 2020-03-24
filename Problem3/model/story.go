package model

import (
	"encoding/json"
	"io"
)

// Story Represents a story
type Story map[string]Chapter

// Chapter represents a chapter in the story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option represents an option in a chapter
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// JSONStory Decode io.Reader containing the story into a JSON document
func JSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}
