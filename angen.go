package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	delimiters    = " _~!?@#$%^&*-+=,."
	delimitersMax = len(delimiters) - 1
)

var (
	//go:embed adjectives.txt
	adjectivesTxt string
	//go:embed names.txt
	namesTxt string

	adjectivesMax int
	namesMax      int

	adjectives []string
	names      []string
)

func Generate() string {
	rand.Seed(time.Now().UnixNano())
	adjectivesIdx := rand.Intn(adjectivesMax)
	namesIdx := rand.Intn(namesMax)
	delimitersIdx := rand.Intn(delimitersMax)
	number := rand.Intn(999)
	return fmt.Sprintf(
		"%s%s%s%s%03d",
		adjectives[adjectivesIdx],
		string(delimiters[delimitersIdx]),
		names[namesIdx],
		string(delimiters[delimitersIdx]),
		number,
	)
}

func init() {
	adjectives = strings.Split(adjectivesTxt, "\n")
	names = strings.Split(namesTxt, "\n")

	adjectivesMax = len(adjectives) - 1
	namesMax = len(names) - 1
}
