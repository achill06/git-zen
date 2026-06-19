package git

import (
	"os/exec"
	"strings"
)

func GetRecentBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--sort=-committerdate", "--format=%(refname:short)")

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Clean and parse the output
	rawOutput := string(out)
	lines := strings.Split(rawOutput, "\n")

	var branches []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			branches = append(branches, trimmed)
		}
	}

	return branches, nil
}
