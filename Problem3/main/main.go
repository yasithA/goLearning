package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type story map[string]chapter

type chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	storyFile, _ := os.Open("story.json")
	defer storyFile.Close()

	decoder := json.NewDecoder(storyFile)
	//var storyArcs storyArc
	var story story
	err := decoder.Decode(&story)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(story["intro"].Title)
}
