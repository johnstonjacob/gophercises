package parsejson

import (
	"encoding/json"
	"io/ioutil"
)

// ParseJSON returns a parsedJSON struct
func ParseJSON(path string) (Story, error) {
	var c Story

	f, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(f, &c)

	if err != nil {
		return nil, err
	}

	return c, nil

}

// StoryArc is the struct in our map
type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Story is a map of storyarc
type Story map[string]StoryArc
