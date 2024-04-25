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
)

func getCommitCountOfEachDay(commits []string) map[string]int {
	commitsCountPerDay := make(map[string]int)
	for _, day := range commits {
		commitsCountPerDay[day] += 1

	}
	return commitsCountPerDay
}

func getCurrentRepo() int {
	cmd := exec.Command("git", "log", "--date=short", "--pretty=format:%ad")
	cmd.Dir = "."
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("error %s", err)
		return 0
	}
	logOutput := string(output)
	logs := strings.Split(logOutput, "\n")
	for _, log := range logs {
		fmt.Println(log)
	}
	commitsPerDay := getCommitCountOfEachDay(logs)
	fmt.Print(commitsPerDay)
	commitCount := len(logs)
	return commitCount

}
func main() {
	fmt.Print(DarkGreen, "This is dark green.\n")
	fmt.Print(ModerateGreen, "This is moderate green.\n")
	fmt.Print(LightGreen, "This is light green.\n")
	fmt.Println(getCurrentRepo())
}
