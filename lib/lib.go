package lib

import "github.com/c-bata/go-prompt"

type Prompt struct {
	Title      string
	Prefix     string
	IsEnable   bool
	Suggestion *[]prompt.Suggest
	IsModule   bool
	Module     int
	SubModule  int
}
