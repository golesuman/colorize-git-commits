package main

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	HIGH_THRESHOLD     = 20
	MODERATE_THRESHOLD = 10
	LOW_THRESHOLD      = 1
)

const (
	DarkGreen     = "\033[1;32m"
	ModerateGreen = "\033[0;32m"
	LightGreen    = "\033[2;32m"
	Reset         = "\033[0m"
)

func getCommitCountOfEachDay(commits []string) map[string]int {
	commitsCountPerDay := make(map[string]int)
	for _, day := range commits {
		commitsCountPerDay[day]++
	}
	return commitsCountPerDay
}

func getCurrentRepo() {
	cmd := exec.Command("git", "log", "--date=short", "--pretty=format:%ad")
	cmd.Dir = "."
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	logOutput := string(output)
	logs := strings.Split(logOutput, "\n")
	commitsPerDay := getCommitCountOfEachDay(logs)

	// Print table header
	fmt.Printf("| %-12s | %-6s | %-6s |\n", "Day", "Commits", "Color")

	for day, count := range commitsPerDay {
		var color string
		switch {
		case count >= HIGH_THRESHOLD:
			color = DarkGreen
		case count >= MODERATE_THRESHOLD:
			color = ModerateGreen
		default:
			color = LightGreen
		}
		fmt.Printf("| %-12s | %-6d | %s%s%s |\n", day, count, color, "██", Reset)
	}
}

func main() {
	getCurrentRepo()
}
