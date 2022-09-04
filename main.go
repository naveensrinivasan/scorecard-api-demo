package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
	"sync/atomic"
)

type Scorecard struct {
	Date string `json:"date"`
	Repo struct {
		Name   string `json:"name"`
		Commit string `json:"commit"`
	} `json:"repo"`
	Scorecard struct {
		Version string `json:"version"`
		Commit  string `json:"commit"`
	} `json:"scorecard"`
	Score  float64 `json:"score"`
	Checks []struct {
		Name          string   `json:"name"`
		Score         int      `json:"score,omitempty"`
		Reason        string   `json:"reason"`
		Details       []string `json:"details"`
		Documentation struct {
			Short string `json:"short"`
			Url   string `json:"url"`
		} `json:"documentation"`
	} `json:"checks"`
}

func main() {
	repoLocation := os.Args[1]
	if repoLocation == "" {
		panic("repoLocation is empty")
	}
	dependencies, err := FetchDependencies(repoLocation)
	if err != nil {
		panic(err)
	}
	fmt.Println("Fuzzing dependencies:")
	var ops uint64
	var wg sync.WaitGroup

	for _, dep := range dependencies {
		dependency := dep
		wg.Add(1)
		go func(dep string) {
			defer wg.Done()
			maintained, score, err := fuzzed(dependency)
			if err != nil {
				return
			}
			if maintained && score >= 7 {
				atomic.AddUint64(&ops, 1)
				fmt.Println(dependency, score)
			}
		}(dep)
	}
	wg.Wait()
	fmt.Println("-----------------")
	fmt.Println("The number of dependencies are", len(dependencies))
	fmt.Println("The number of dependencies that are fuzzed are", ops)
}

// fuzzed checks if the dependency is fuzzed by checking the scorecard API
func fuzzed(repo string) (bool, int, error) {
	//repo = "github.com/sigstore/sigstore"
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.securityscorecards.dev/projects/%s", repo), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, 0, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, 0, err
	}
	var scorecard Scorecard
	err = json.Unmarshal(result, &scorecard)
	if err != nil {
		return true, 0, err
	}
	for _, check := range scorecard.Checks {
		if check.Name == "Fuzzing" {
			if check.Score >= 7 || check.Score < 0 {
				return true, check.Score, nil
			}
			return false, 0, nil
		}
	}
	return false, 0, nil
}

// FetchDependencies parses the dependencies in the go.mod using the `go list command`
// This functions expects the directory to contain the go.mod file.
func FetchDependencies(directory string) ([]string, error) {
	modquery := `
	go list -m -f '{{if not (or  .Main)}}{{.Path}}{{end}}' all \
	| grep "^github" \
	| sort -u \
	| cut -d/ -f1-3 \
	| awk '{print $1}' \
	| tr '\n' ',' 
	`
	// Runs the modquery to generate the dependencies
	c := exec.Command("bash", "-c", fmt.Sprintf("cd %s;", directory)+modquery)
	data, err := c.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to run go list: %w %s", err, string(data))
	}
	m := make(map[string]bool)
	parameters := []string{}
	result := append(parameters, strings.Split(string(data), ",")...)
	//filter the result to remove empty strings and duplicates
	for _, dep := range result {
		if dep != "" {
			m[dep] = true
		}
	}
	result = []string{}
	for dep := range m {
		result = append(result, dep)
	}
	return result, nil
}
