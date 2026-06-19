package git

import (
	"encoding/json"
	"os/exec"
)

// PRInfo matches the JSON structure returned by the gh CLI
type PRInfo struct {
	HeadRefName string `json:"headRefName"`
	State       string `json:"state"`
}

// GetPRStatuses fetches up to 50 recent PRs and maps them by branch name
func GetPRStatuses() (map[string]string, error) {
	// Execute: gh pr list --state all --json headRefName,state --limit 50
	cmd := exec.Command("gh", "pr", "list", "--state", "all", "--json", "headRefName,state", "--limit", "50")

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var prs []PRInfo
	if err := json.Unmarshal(out, &prs); err != nil {
		return nil, err
	}

	statusMap := make(map[string]string)
	for _, pr := range prs {
		statusMap[pr.HeadRefName] = pr.State
	}

	return statusMap, nil
}
