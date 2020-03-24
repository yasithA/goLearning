package main

import (
	"flag"
	"os"

	"github.com/yasithA/goLearning/Problem3/model"
)

func main() {
	filename := flag.String("file", "story.json", "The JSON file with CYOA story.")

	storyFile, _ := os.Open(*filename)
	defer storyFile.Close()

	story, err := model.JSONStory(storyFile)
}
