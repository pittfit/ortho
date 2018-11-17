package testing

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/go-jsonnet"
)

var fixtureFiles = []string{
	"one",
}

type fixture struct {
	Input  string          `json:"input",string`
	Tokens []string        `json:"tokens"`
	Ast    string          `json:"ast"`
	Output []fixtureOutput `json:"output"`
}

type fixtureOutput struct {
	Pattern string   `json:"pattern"`
	Strings []string `json:"strings"`
	// Settings map[string]bool `json:"options"`
}

func loadFixtures() ([]fixture, error) {
	vm := jsonnet.MakeVM()
	bytes, err := vm.EvaluateSnippet("_", snippet())

	var result []fixture
	err = json.Unmarshal([]byte(bytes), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func snippet() string {
	snippets := []string{}

	for _, file := range fixtureFiles {
		snippets = append(snippets, fmt.Sprintf("(import 'fixtures/%s.jsonnet')", file))
	}

	return strings.Join(snippets, "+")
}
