package advent

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const CookieVar = "ADVENT_COOKIE"

type Solution func(input *Input) (int, error)

type Runner interface {
	Run(day int, s Solution)
}

type runner struct {
	cookie string
}

func (r *runner) Run(day int, s Solution) {
	client, err := NewClient(r.cookie)
	if err != nil {
		log.Fatalf("Day %d Error creating Advent Client: %s", day, err.Error())
		return
	}
	input, err := client.GetInput(day)
	if err != nil {
		fmt.Printf("Day %d Error reading Advent input: %s", day, err.Error())
		return
	}
	start := time.Now()
	answer, err := s(input)
	if err != nil {
		fmt.Printf("Day %d Error within Solution: %s", day, err.Error())
		return
	}
	fmt.Printf("Day %d Answer: %v - took %s\n", day, answer, time.Since(start))
}

func getEnvVar(name string) string {
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if pair[0] == name {
			return pair[1]
		}
	}
	return ""
}

func NewRunner() Runner {
	cookie := getEnvVar(CookieVar)
	if cookie == "" {
		log.Fatalf("env variable: %s not set", CookieVar)
	}
	return &runner{cookie: cookie}
}
